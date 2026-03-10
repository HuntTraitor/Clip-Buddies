package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) routes() http.Handler {
	router := gin.New()

	// Middleware
	router.Use(app.recoverPanic, app.logRequest, app.rateLimit())

	// Optional metrics middleware

	// Custom 404 / 405
	router.NoRoute(app.notFoundResponse)
	router.NoMethod(app.methodNotAllowedResponse)

	// Health Check Routes
	app.handlers.System.RegisterRoutes(router.Group("/healthcheck"))

	// User Routes
	//app.loadUserRoutes(router.Group("/users"))
	return router
}

//
//func (app *application) loadUserRoutes(group *gin.RouterGroup) {
//	group.POST("/", app.registerUserHandler)
//	group.PUT("/activated", app.activateUserHandler)
//
//	authenticated := group.Group("/")
//	authenticated.Use(app.requireAuthenticatedUser)
//	authenticated.GET("/verify", app.verifyUserHandler)
//
//	group.PUT("/password", app.updateUserPasswordHandler)
//}
//
//func (app *application) loadTokenRoutes(group *gin.RouterGroup) {
//	group.POST("/authentication", app.createAuthenticationTokenHandler)
//	group.POST("/password-reset", app.createPasswordResetTokenHandler)
//}
//
//func (app *application) loadDebugRoutes(group *gin.RouterGroup) {
//	group.GET("/vars", gin.WrapH(expvar.Handler()))
//}
//
//func (app *application) loadMethodRoutes(group *gin.RouterGroup) {
//	group.GET("/", app.listMethodsHandler)
//	group.GET("/:id", app.getMethodHandler)
//}
//
//func (app *application) loadCoffeeRoutes(group *gin.RouterGroup) {
//	authenticated := group.Group("/")
//	authenticated.Use(app.requireAuthenticatedUser)
//
//	authenticated.GET("/", app.listCoffeesHandler)
//	authenticated.POST("/", app.createCoffeeHandler)
//	authenticated.GET("/:id", app.getCoffeeHandler)
//	authenticated.PATCH("/:id", app.updateCoffeeHandler)
//	authenticated.DELETE("/:id", app.deleteCoffeeHandler)
//}
//
//func (app *application) loadRecipeRoutes(group *gin.RouterGroup) {
//	authenticated := group.Group("/")
//	authenticated.Use(app.requireAuthenticatedUser)
//
//	authenticated.POST("/", app.createRecipeHandler)
//	authenticated.GET("/", app.listRecipesHandler)
//	authenticated.PATCH("/:id", app.updateRecipeHandler)
//	authenticated.DELETE("/:id", app.deleteRecipeHandler)
//}
