package handlers

import (
	"net/http"
	"veritas/core/usecases"
	"veritas/internal/ports/dtos"

	"github.com/gin-gonic/gin"
)

// RoleHandler handles role-related HTTP requests.
type RoleHandler struct {
	roleUseCase usecases.RoleUsecase
}

// NewRoleHandler creates a new RoleHandler with the given RoleUseCase.
func NewRoleHandler(roleUsecase usecases.RoleUsecase) *RoleHandler {
	return &RoleHandler{
		roleUseCase: roleUsecase,
	}
}

// CreateRole godoc
// @Summary Create a new role
// @Description Create a new role with the input payload
// @Tags roles
// @Accept  json
// @Produce  json
// @Param role body dtos.CreateRoleInputDTO true "Create Role"
// @Success 201 {object} dtos.CreateRoleOutputDTO
// @Security ApiKeyAuth
// @Router /roles [post]
func (h *RoleHandler) CreateRole(c *gin.Context) {
	var roleInput dtos.CreateRoleInputDTO
	if err := c.ShouldBindJSON(&roleInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := usecases.CreateRoleInput{
		Name:        roleInput.Name,
		Description: roleInput.Description,
	}

	id, err := h.roleUseCase.CreateRole(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	output := dtos.CreateRoleOutputDTO{
		ID: id,
	}

	c.JSON(http.StatusCreated, output)
}

// GetRole godoc
// @Summary Get a role by ID
// @Description Get a role by ID
// @Tags roles
// @Produce  json
// @Param id path string true "Role ID"
// @Success 200 {object} domain.Role
// @Security ApiKeyAuth
// @Router /roles/{id} [get]
func (h *RoleHandler) GetRole(c *gin.Context) {
	id := c.Param("id")

	role, err := h.roleUseCase.ReadRole(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, role)
}

// UpdateRole godoc
// @Summary Update a role
// @Description Update a role with the input payload
// @Tags roles
// @Accept  json
// @Produce  json
// @Param id path string true "Role ID"
// @Param role body dtos.UpdateRoleInputDTO true "Update Role"
// @Success 200 {object} dtos.UpdateRoleOutputDTO
// @Security ApiKeyAuth
// @Router /roles/{id} [put]
func (h *RoleHandler) UpdateRole(c *gin.Context) {
	id := c.Param("id")

	var roleInput dtos.UpdateRoleInputDTO
	if err := c.ShouldBindJSON(&roleInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := usecases.UpdateRoleInput{
		Name:        roleInput.Name,
		Description: roleInput.Description,
	}

	role, err := h.roleUseCase.UpdateRole(c.Request.Context(), id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	output := dtos.UpdateRoleOutputDTO{
		ID: role.ID,
	}

	c.JSON(http.StatusOK, output)
}

// DeleteRole godoc
// @Summary Delete a role
// @Description Delete a role by ID
// @Tags roles
// @Param id path string true "Role ID"
// @Success 200 {object} object{message=string}
// @Security ApiKeyAuth
// @Router /roles/{id} [delete]
func (h *RoleHandler) DeleteRole(c *gin.Context) {
	id := c.Param("id")

	err := h.roleUseCase.DeleteRole(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role deleted successfully"})
}

// GetAllRoles godoc
// @Summary Get all roles
// @Description Get all roles
// @Tags roles
// @Produce  json
// @Success 200 {array} domain.Role
// @Security ApiKeyAuth
// @Router /roles [get]
func (h *RoleHandler) GetAllRoles(c *gin.Context) {
	roles, err := h.roleUseCase.GetAllRoles(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, roles)
}
