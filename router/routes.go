package router

import (
	"github.com/fymorGod/gopportunities/handlers"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	handlers.InitializeHandler()
	v1 := router.Group("/api/v1/")
	{
		v1.GET("/opening", handlers.ShowOpeningHandler)
		v1.POST("/opening", handlers.CreateOpeningHandler)
		v1.DELETE("/opening", handlers.DeleteOpeningHandler)
		v1.PUT("/opening", handlers.UpdateOpeningHandler)
		v1.GET("/openings", handlers.ListOpeningsHandler)
	}
}
