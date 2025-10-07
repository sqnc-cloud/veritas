package handlers

import (
	"net/http"
	"veritas/core/usecases"
	"veritas/internal/ports/dtos"

	"github.com/gin-gonic/gin"
)

// UserHandler handles user-related HTTP requests.
type UserHandler struct {
	userUseCase usecases.UserUsecase
}

// NewUserHandler creates a new UserHandler with the given UserUseCase.
func NewUserHandler(userUsecase usecases.UserUsecase) *UserHandler {
	return &UserHandler{
		userUseCase: userUsecase,
	}
}

// GetUser godoc
// @Summary Get a user by ID
// @Description Get a user by ID
// @Tags users
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} dtos.CreateUserOutputDTO
// @Security ApiKeyAuth
// @Router /users/{id} [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")

	user, err := h.userUseCase.ReadUser(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	output := dtos.CreateUserOutputDTO{
		ID:    user.ID.Hex(),
		Name:  user.Username,
		Email: user.Email,
	}

	c.JSON(http.StatusOK, output)
}

// UpdateUser godoc
// @Summary Update a user
// @Description Update a user with the input payload
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Param user body dtos.UpdateUserInputDTO true "Update User"
// @Success 200 {object} dtos.UpdateUserOutputDTO
// @Security ApiKeyAuth
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var userInput dtos.UpdateUserInputDTO
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := usecases.UpdateUserInput{
		Name:     userInput.Name,
		Email:    userInput.Email,
		Password: userInput.Password,
	}

	user, err := h.userUseCase.UpdateUser(c.Request.Context(), id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	output := dtos.UpdateUserOutputDTO{
		ID:    user.ID.Hex(),
		Name:  user.Username,
		Email: user.Email,
	}

	c.JSON(http.StatusOK, output)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user by ID
// @Tags users
// @Param id path string true "User ID"
// @Success 200 {object} object{message=string}
// @Security ApiKeyAuth
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	err := h.userUseCase.DeleteUser(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Produce  json
// @Success 200 {array} dtos.CreateUserOutputDTO
// @Security ApiKeyAuth
// @Router /users [get]
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.userUseCase.GetAllUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var outputUsers []dtos.CreateUserOutputDTO
	for _, user := range users {
		outputUsers = append(outputUsers, dtos.CreateUserOutputDTO{
			ID:    user.ID.Hex(),
			Name:  user.Username,
			Email: user.Email,
		})
	}

	c.JSON(http.StatusOK, outputUsers)
}
