package main

import (
	"WeatherProfile_Service/controllers"
	"WeatherProfile_Service/middleware"
	"WeatherProfile_Service/models"
	"WeatherProfile_Service/repositories"
	"WeatherProfile_Service/services"
	"WeatherProfile_Service/utils"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Загрузите переменные окружения из .env файла
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	db := utils.GetDBInstance()
	db.AutoMigrate(&models.User{}) // Автоматическая миграция

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	router := gin.Default()

	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)

	protected := router.Group("/", middleware.AuthMiddleware())
	{
		protected.GET("/profile", userController.GetProfile)
	}

	router.Run(":8080")
}
