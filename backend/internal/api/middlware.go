package api

import (
	"bytes"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tomasen/realip"
	"golang.org/x/time/rate"
)

func (app *Application) recoverPanic(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.Header("Connection", "close")
			app.serverErrorResponse(c, fmt.Errorf("%v", err))
			c.Abort()
		}
	}()

	c.Next()
}

func (app *Application) logRequest(c *gin.Context) {
	start := time.Now()

	var requestBody string
	if c.Request.Body != nil {
		rawBody, err := io.ReadAll(c.Request.Body)
		if err == nil {
			c.Request.Body = io.NopCloser(bytes.NewReader(rawBody))

			if len(rawBody) > 0 {
				requestBody = string(rawBody)
				if len(requestBody) > 2000 {
					requestBody = requestBody[:2000] + "...[truncated]"
				}
			}
		}
	}

	c.Next()

	app.Logger.Info("http request",
		"ip", c.ClientIP(),
		"method", c.Request.Method,
		"path", c.Request.URL.Path,
		"query", c.Request.URL.RawQuery,
		"status", c.Writer.Status(),
		"duration", time.Since(start).String(),
		"body", requestBody,
	)
}

func (app *Application) rateLimit() gin.HandlerFunc {
	type client struct {
		limiter  *rate.Limiter
		lastSeen time.Time
	}

	var (
		mu      sync.Mutex
		clients = make(map[string]*client)
	)

	go func() {
		for {
			time.Sleep(time.Minute)

			mu.Lock()
			for ip, client := range clients {
				if time.Since(client.lastSeen) > app.Config.Limiter.Expiration {
					delete(clients, ip)
				}
			}
			mu.Unlock()
		}
	}()

	return func(c *gin.Context) {
		if app.Config.Limiter.Enabled {
			ip := realip.FromRequest(c.Request)

			mu.Lock()
			if _, found := clients[ip]; !found {
				clients[ip] = &client{
					limiter: rate.NewLimiter(
						rate.Limit(app.Config.Limiter.RPS),
						app.Config.Limiter.Burst,
					),
				}
			}

			clients[ip].lastSeen = time.Now()
			allowed := clients[ip].limiter.Allow()
			mu.Unlock()

			if !allowed {
				app.rateLimitExceededResponse(c)
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
