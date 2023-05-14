package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"urlShortener/controllers"
	"urlShortener/middlewares"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	//models.ConnectDataBase() todo

	router := gin.Default()

	auth := router.Group("/auth")
	auth.POST("/register", controllers.Register)
	auth.POST("/login", controllers.Login)

	api := router.Group("/api")
	api.Use(middlewares.JwtAuth)

	_ = router.Run(":8080")
}
