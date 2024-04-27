package http

import (
	"net/http"

	"github.com/Ixorlive/tw_vk/backend/services/auth/internal/entity"
	"github.com/Ixorlive/tw_vk/backend/services/auth/internal/usecase"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Service usecase.AuthService
}

func NewAuthController(auth usecase.AuthService) AuthController {
	return AuthController{auth}
}

type UserRequestBody struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// AuthByLogin authenticates a user by their login credentials.
// @Summary Authenticate by login
// @Description Authenticates users by login and password, returns a JWT token if successful.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param requestBody body UserRequestBody true "Login and Password"
// @Success 200 {object} map[string]string "JWT Token if authentication is successful"
// @Failure 400 {object} map[string]string "Bad request if the request body cannot be parsed"
// @Failure 500 {object} map[string]string "Internal Server Error for any server issues"
// @Router /login [post]
func (c *AuthController) AuthByLogin(ctx *gin.Context) {
	var requestBody UserRequestBody
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := entity.User{Login: requestBody.Login, Password: requestBody.Password}

	token, err := c.Service.Auth(ctx.Request.Context(), user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if token == nil {
		ctx.JSON(http.StatusOK, gin.H{"error": "Invalid token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token.Token})
}

type TokenRequestBody struct {
	Token string `json:"token"`
}

// AuthByToken authenticates a user by a JWT token.
// @Summary Authenticate by token
// @Description Verifies the JWT token and returns user details if the token is valid.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param requestBody body TokenRequestBody true "Login and Password"
// @Success 200 {object} map[string]interface{} "User details if token is valid"
// @Failure 400 {object} map[string]string "Bad request if the request body cannot be parsed"
// @Failure 500 {object} map[string]string "Internal Server Error for any server issues"
// @Router /token [post]
func (c *AuthController) AuthByToken(ctx *gin.Context) {
	var requestBody TokenRequestBody
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := c.Service.AuthByToken(ctx, requestBody.Token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if user.Login == "" {
		ctx.JSON(http.StatusOK, gin.H{"error": "Wrong Login or Password"})
	}
	ctx.JSON(http.StatusOK, gin.H{"id": user.Id, "login": user.Login})
}

// Register registers a new user.
// @Summary Register a new user
// @Description Registers a new user with a login and password, returns a status message.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param requestBody body UserRequestBody true "Login and Password"
// @Success 200 {object} map[string]string "OK status if registration is successful, or an error message"
// @Failure 400 {object} map[string]string "Bad request if the request body cannot be parsed"
// @Failure 500 {object} map[string]string "Internal Server Error for any server issues"
// @Router /register [post]
func (c *AuthController) Register(ctx *gin.Context) {
	var requestBody UserRequestBody
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	status, err := c.Service.Register(ctx, entity.User{Login: requestBody.Login, Password: requestBody.Password})

	if err != nil || status == usecase.Error {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if status == usecase.IncorrectLoginOrPassword {
		ctx.JSON(http.StatusOK, gin.H{"error": "Wrong Login or Password"})
		return
	}
	if status == usecase.UserExists {
		ctx.JSON(http.StatusOK, gin.H{"error": "Login is taken"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
}
