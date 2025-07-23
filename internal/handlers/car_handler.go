package handlers

import (
	"fmt"
	"net/http"

	"github.com/Camarg0/search-car-api/internal/services"
	"github.com/gin-gonic/gin"
)

func GetCarInfo(c *gin.Context) {
	fmt.Println("Cheguei aqui!")

	carModel := c.Param("mocked-model")
	carInfo, found := services.GetMockedCarInfo(carModel)

	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Modelo não encontrado"})
		return
	}

	c.JSON(http.StatusOK, carInfo)
}

func GetCarInfoFromOpenAI(c *gin.Context) {

	carModel := c.Param("model")

	carInfo, err := services.GetCarInfoFromOpenAI(carModel)

	if err != nil {

		fmt.Println("Cheguei aqui no erro!")

		fmt.Println(err)
		c.JSON(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, carInfo)
}
