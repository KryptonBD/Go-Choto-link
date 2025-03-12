package main

import (
	"choto-link/routes"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Unable to Load env", err)
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":" + port)

	fmt.Println("APP RUNNING ON ", port)
}
