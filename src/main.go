package main

import (
	"backend-template/controllers"
	"backend-template/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}

	r := gin.Default()

	models.ConnectDB()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hola k tal ombre"})
	})

	r.GET("/users", controllers.FindUsers)

	err = r.Run(os.Getenv("API_ADDR") + ":" + os.Getenv("API_PORT"))

	if err != nil {
		panic("Error creating router")
	}
}
