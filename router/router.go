package router

import (
	"net/http"

	"github.com/asynccnu/grade_service_v2/handler/grade"
	"github.com/asynccnu/grade_service_v2/handler/sd"
	"github.com/asynccnu/grade_service_v2/router/middleware"

	"github.com/gin-gonic/gin"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// The health check handlers
	api := g.Group("/api/grade/v2")
	api.Use(middleware.AuthMiddleware())
	{
		api.GET("", grade.Get)
		api.GET("/detail", grade.GetDetail)
	}

	// The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
