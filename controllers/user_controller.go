package controllers

import (
	"WeatherProfile_Service/models"
	"WeatherProfile_Service/services"
	"WeatherProfile_Service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{service: service}
}

func (ctrl *UserController) Register(c *gin.Context) {
	var userRequest models.User
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.service.Register(&userRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Формируем ответ без пароля
	userResponse := models.UserResponse{
		ID:       userRequest.ID,
		Username: userRequest.Username,
		Email:    userRequest.Email,
	}
	c.JSON(http.StatusCreated, gin.H{"user": userResponse})
}

func (ctrl *UserController) Login(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := ctrl.service.Login(credentials.Email, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (ctrl *UserController) GetProfile(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	user, err := ctrl.service.GetUserByID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	// Формируем ответ без пароля
	userResponse := models.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
	c.JSON(http.StatusOK, gin.H{"user": userResponse})
}
