package router

import (
	"github.com/brunobotter/ecommerce-usuario/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/health", handler.HealthCheck)
		v1.GET("/usuario/:id", handler.ShowUsuarioHandler)
		v1.GET("/usuarios", handler.ListUsuairoHandler)
		v1.POST("/usuario", handler.CreateUsuarioHandler)
		v1.PUT("/usuario/:id", handler.UpdateUsuarioHandler)
		v1.DELETE("/usuario/:id", handler.DeleteUsuarioHandler)
	}
}
