package handlers

import (
	"net/http"
	"time"
	"veritas/core/usecases"
	"veritas/internal/ports/dtos"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AuthHandler handles authentication-related HTTP requests.
type AuthHandler struct {
	userUseCase usecases.UserUsecase
}

// NewAuthHandler creates a new AuthHandler with the given UserUseCase.
func NewAuthHandler(userUsecase usecases.UserUsecase) *AuthHandler {
	return &AuthHandler{
		userUseCase: userUsecase,
	}
}

// Login godoc
// @Summary Authenticate a user
// @Description Authenticate a user with the input payload
// @Tags auth
// @Accept  json
// @Produce  json
// @Param user body dtos.LoginInputDTO true "Login"
// @Success 200 {object} object{token=string}
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var loginInput dtos.LoginInputDTO
	if err := c.ShouldBindJSON(&loginInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.userUseCase.VerifyUser(c.Request.Context(), loginInput.Email, loginInput.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": loginInput.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// SignUp godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags auth
// @Accept  json
// @Produce  json
// @Param user body dtos.CreateUserInputDTO true "Create User"
// @Success 201 {object} dtos.CreateUserOutputDTO
// @Router /auth/signup [post]
func (h *AuthHandler) SignUp(c *gin.Context) {
	var userInput dtos.CreateUserInputDTO
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := usecases.CreateUserInput{
		Name:     userInput.Name,
		Email:    userInput.Email,
		Password: userInput.Password,
	}

	userID, err := h.userUseCase.CreateUser(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := h.userUseCase.ReadUser(c.Request.Context(), userID.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	output := dtos.CreateUserOutputDTO{
		ID:    createdUser.ID.Hex(),
		Name:  createdUser.Username,
		Email: createdUser.Email,
	}

	c.JSON(http.StatusCreated, output)
}