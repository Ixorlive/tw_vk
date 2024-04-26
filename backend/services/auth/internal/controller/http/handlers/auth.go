package handlers

import (
	"net/http"

	"github.com/Ixorlive/tw_vk/backend/services/auth/internal/entity"
	"github.com/Ixorlive/tw_vk/backend/services/auth/internal/usecase"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService usecase.AuthService
}

func NewAuthHandler(auth usecase.AuthService) AuthHandler {
	return AuthHandler{auth}
}

type UserRequestBody struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (ah *AuthHandler) AuthByLogin(c *gin.Context) {
	var requestBody UserRequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := entity.User{Login: requestBody.Login, Password: requestBody.Password}

	token, err := ah.authService.Auth(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error. Please try again later."})
		return
	}

	if token == nil {
		c.JSON(http.StatusOK, gin.H{"error": "Incorrect login or password"})
	}

	c.JSON(http.StatusOK, token)
}

func (ah *AuthHandler) AuthByToken(c *gin.Context) {
	type RequestBody struct {
		Token string `json:"token"`
	}
	var requestBody RequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := ah.authService.AuthByToken(c, requestBody.Token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error. Please try again later."})
	}
	if user.Login == "" {
		c.JSON(http.StatusOK, gin.H{"error": "Token is not valid"})
	}
	c.JSON(http.StatusOK, gin.H{"login": user.Login})
}

func (ah *AuthHandler) Register(c *gin.Context) {
	var requestBody UserRequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	status, err := ah.authService.Register(c, entity.User{Login: requestBody.Login, Password: requestBody.Password})

	if err != nil || status == usecase.Error {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error. Please try again later."})
		return
	}
	if status == usecase.IncorrectLoginOrPassword {
		c.JSON(http.StatusOK, gin.H{"error": "Incorrect login or password"})
		return
	}
	if status == usecase.UserExists {
		c.JSON(http.StatusOK, gin.H{"error": "Login is taken"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
