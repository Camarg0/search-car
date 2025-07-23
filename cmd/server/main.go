package main

import (
	"github.com/Camarg0/search-car-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run(":8080")
}
