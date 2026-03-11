package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *Application) Routes() http.Handler {
	router := gin.New()

	router.Use(
		app.recoverPanic,
		app.logRequest,
		app.rateLimit(),
	)

	router.NoRoute(app.notFoundResponse)
	router.NoMethod(app.methodNotAllowedResponse)

	app.SystemHandler.RegisterRoutes(router.Group("/healthcheck"))
	app.AuthHandler.RegisterRoutes(router.Group(""))

	return router
}
