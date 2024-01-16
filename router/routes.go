package router

import (
	"github.com/fymorGod/gopportunities/handlers"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1/")
	{
		v1.GET("/opening", handlers.ListOpeningsHandler)
		v1.POST("/openings", handlers.CreateOpeningHandler)
		v1.DELETE("/openings", handlers.DeleteOpeningHandler)
		v1.PUT("/openings", handlers.UpdateOpeningHandler)
		v1.GET("/openings", handlers.ListOpeningsHandler)
	}
}
