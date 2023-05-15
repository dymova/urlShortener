package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
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
	api.POST("/shorten", controllers.Shorten)
	api.GET("/redirect/:shortCode", controllers.Redirect)
	_ = router.Run(os.Getenv("BASE_URL"))
}
