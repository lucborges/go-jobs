package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lucborges/go-jobs/handler"
)

func initializeRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.GET("/opening", handler.ShowOpeningHandler)
		api.POST("/opening", handler.CreateOpeningHandler)
		api.PUT("/opening", handler.UpdateOpeningHandler)
		api.DELETE("/opening", handler.DeleteOpeningHandler)
		api.GET("/openings", handler.ShowOpeningsHandler)
	}
}
