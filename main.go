package main

import (
	"WeatherProfile_Service/controllers"
	"WeatherProfile_Service/middleware"
	"WeatherProfile_Service/models"
	"WeatherProfile_Service/repositories"
	"WeatherProfile_Service/services"
	"WeatherProfile_Service/utils"

	"github.com/gin-gonic/gin"
)

func main() {
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
