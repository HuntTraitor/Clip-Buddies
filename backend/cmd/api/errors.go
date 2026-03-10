package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type envelope map[string]any

func (app *application) logError(c *gin.Context, err error) {
	method := c.Request.Method
	uri := c.Request.URL.RequestURI()

	app.logger.Error(err.Error(), "method", method, "uri", uri)
}

// errorResponse writes a code and a message wrapped in an "error" key.
func (app *application) errorResponse(c *gin.Context, status int, message any) {
	env := envelope{"error": message}
	c.JSON(status, env)
}

// serverErrorResponse returns a 500 error.
func (app *application) serverErrorResponse(c *gin.Context, err error) {
	app.logError(c, err)

	message := "An internal server error has occurred. Please try again later."
	app.errorResponse(c, http.StatusInternalServerError, message)
}

// notFoundResponse returns a 404 error.
func (app *application) notFoundResponse(c *gin.Context) {
	message := "The requested resource could not be found."
	app.errorResponse(c, http.StatusNotFound, message)
}

// methodNotAllowedResponse returns a 405 error.
func (app *application) methodNotAllowedResponse(c *gin.Context) {
	message := fmt.Sprintf("the %s method is not supported for this resource", c.Request.Method)
	app.errorResponse(c, http.StatusMethodNotAllowed, message)
}

// badRequestResponse returns a 400 error.
func (app *application) badRequestResponse(c *gin.Context, err error) {
	app.errorResponse(c, http.StatusBadRequest, err.Error())
}

// failedValidationResponse returns a 422 error.
func (app *application) failedValidationResponse(c *gin.Context, errors map[string]string) {
	app.errorResponse(c, http.StatusUnprocessableEntity, errors)
}

func (app *application) rateLimitExceededResponse(c *gin.Context) {
	message := "rate limit exceeded"
	app.errorResponse(c, http.StatusTooManyRequests, message)
}

func (app *application) editConflictResponse(c *gin.Context) {
	message := "unable to update the record due to an edit conflict, please try again"
	app.errorResponse(c, http.StatusConflict, message)
}

func (app *application) invalidCredentialsResponse(c *gin.Context) {
	message := "invalid authentication credentials"
	app.errorResponse(c, http.StatusUnauthorized, message)
}

func (app *application) invalidAuthenticationTokenResponse(c *gin.Context) {
	c.Header("WWW-Authenticate", "Bearer")
	message := "invalid or missing authentication token"
	app.errorResponse(c, http.StatusUnauthorized, message)
}

func (app *application) authenticationRequiredResponse(c *gin.Context) {
	message := "you must be authenticated to access this resource"
	app.errorResponse(c, http.StatusUnauthorized, message)
}

func (app *application) inactiveAccountResponse(c *gin.Context) {
	message := "your user account must be activated to access this feature"
	app.errorResponse(c, http.StatusUnauthorized, message)
}

func (app *application) unknownCoffeeResponse(c *gin.Context) {
	message := "the requested coffee could not be found"
	app.errorResponse(c, http.StatusNotFound, message)
}

func (app *application) unknownMethodResponse(c *gin.Context) {
	message := "the requested method could not be found"
	app.errorResponse(c, http.StatusNotFound, message)
}
