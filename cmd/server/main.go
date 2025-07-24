package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Camarg0/search-car-api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Carrega variáveis do .env (interessante fazer isso no main.go)
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("Arquivo .env não encontrado. Continuando com variáveis do sistema.")
	}

	fmt.Println("Chave da OpenAI:", os.Getenv("OPENAI_API_KEY"))

	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run(":8080")
}
