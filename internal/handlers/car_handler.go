package handlers

import (
	"net/http"

	"github.com/Camarg0/search-car-api/internal/services"
	"github.com/gin-gonic/gin"
)

func GetCarInfo(c *gin.Context) {
	carModel := c.Param("mocked-model")
	carInfo, found := services.GetMockedCarInfo(carModel)

	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Modelo não encontrado"})
		return
	}

	c.JSON(http.StatusOK, carInfo)
}

func GetCarInfoHandler(c *gin.Context) {
	carModel := c.Param("model")
	carInfo, err := services.GetCarInfoFromOpenAI(carModel)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, carInfo)
}
