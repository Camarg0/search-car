package routes

import (
	"github.com/Camarg0/search-car-api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		//api.GET("/cars/:mocked-model", handlers.GetCarInfo)
		api.POST("/cars/:model", handlers.GetCarInfoFromOpenAI)
	}
}
