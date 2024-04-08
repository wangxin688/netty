package router

import (
	"netty/core"
	"netty/register/middleware"

	"github.com/gin-gonic/gin"
)


// SetUpRoute sets up the routes for the application
func SetUpRoute() *gin.Engine {
	// Determine the environment to run the app in
	env := core.Settings.Env
	if env == "" {
		env = core.Dev
	}

	// Set the mode of the gin router based on the environment
	var mode string
	if env == core.Dev || env == core.Stage {
		mode = gin.DebugMode
	} else {
		mode = gin.ReleaseMode
	}
	gin.SetMode(mode)

	// If the allowed hosts slice is nil, set it to allow all (*) hosts
	allowHosts := core.Settings.AllowedHosts
	if allowHosts == nil {
		allowHosts = []string{"*"}
	}

	// Create a new gin router
	router := gin.New()

	// Set the trusted proxies for the router
	router.SetTrustedProxies(allowHosts)

	// Add middleware to the router
	router.Use(gin.Logger(), gin.Recovery())
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.RequestID())
	router.Use(middleware.I18nMiddleware())

	// router.Use(middleware.AuthMiddleware())

	// Register all routers here
	RegisterRoutes(router)

	// Return the router
	return router
}
