package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type envelope map[string]any

func (app *Application) logError(c *gin.Context, err error) {
	method := c.Request.Method
	uri := c.Request.URL.RequestURI()

	app.Logger.Error(err.Error(), "method", method, "uri", uri)
}

func (app *Application) errorResponse(c *gin.Context, status int, message any) {
	env := envelope{"error": message}
	c.JSON(status, env)
}

func (app *Application) serverErrorResponse(c *gin.Context, err error) {
	app.logError(c, err)

	message := "An internal server error has occurred. Please try again later."
	app.errorResponse(c, http.StatusInternalServerError, message)
}

func (app *Application) notFoundResponse(c *gin.Context) {
	message := "The requested resource could not be found."
	app.errorResponse(c, http.StatusNotFound, message)
}

func (app *Application) methodNotAllowedResponse(c *gin.Context) {
	message := fmt.Sprintf("the %s method is not supported for this resource", c.Request.Method)
	app.errorResponse(c, http.StatusMethodNotAllowed, message)
}

func (app *Application) badRequestResponse(c *gin.Context, err error) {
	app.errorResponse(c, http.StatusBadRequest, err.Error())
}

func (app *Application) failedValidationResponse(c *gin.Context, errors map[string]string) {
	app.errorResponse(c, http.StatusUnprocessableEntity, errors)
}

func (app *Application) rateLimitExceededResponse(c *gin.Context) {
	message := "rate limit exceeded"
	app.errorResponse(c, http.StatusTooManyRequests, message)
}

func (app *Application) editConflictResponse(c *gin.Context) {
	message := "unable to update the record due to an edit conflict, please try again"
	app.errorResponse(c, http.StatusConflict, message)
}

func (app *Application) invalidCredentialsResponse(c *gin.Context) {
	message := "invalid authentication credentials"
	app.errorResponse(c, http.StatusUnauthorized, message)
}

func (app *Application) invalidAuthenticationTokenResponse(c *gin.Context) {
	c.Header("WWW-Authenticate", "Bearer")
	message := "invalid or missing authentication token"
	app.errorResponse(c, http.StatusUnauthorized, message)
}

func (app *Application) authenticationRequiredResponse(c *gin.Context) {
	message := "you must be authenticated to access this resource"
	app.errorResponse(c, http.StatusUnauthorized, message)
}

func (app *Application) inactiveAccountResponse(c *gin.Context) {
	message := "your user account must be activated to access this feature"
	app.errorResponse(c, http.StatusUnauthorized, message)
}

func (app *Application) unknownCoffeeResponse(c *gin.Context) {
	message := "the requested coffee could not be found"
	app.errorResponse(c, http.StatusNotFound, message)
}

func (app *Application) unknownMethodResponse(c *gin.Context) {
	message := "the requested method could not be found"
	app.errorResponse(c, http.StatusNotFound, message)
}
