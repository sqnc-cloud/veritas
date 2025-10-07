package handlers

import (
	"net/http"
	"veritas/core/usecases"
	"veritas/internal/ports/dtos"

	"github.com/gin-gonic/gin"
)

// ClaimHandler handles claim-related HTTP requests.
type ClaimHandler struct {
	claimUseCase usecases.ClaimUsecase
}

// NewClaimHandler creates a new ClaimHandler with the given ClaimUseCase.
func NewClaimHandler(claimUsecase usecases.ClaimUsecase) *ClaimHandler {
	return &ClaimHandler{
		claimUseCase: claimUsecase,
	}
}

// CreateClaim godoc
// @Summary Create a new claim
// @Description Create a new claim with the input payload
// @Tags claims
// @Accept  json
// @Produce  json
// @Param claim body dtos.CreateClaimInputDTO true "Create Claim"
// @Success 201 {object} dtos.CreateClaimOutputDTO
// @Security ApiKeyAuth
// @Router /claims [post]
func (h *ClaimHandler) CreateClaim(c *gin.Context) {
	var claimInput dtos.CreateClaimInputDTO
	if err := c.ShouldBindJSON(&claimInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := usecases.CreateClaimInput{
		Name:        claimInput.Name,
		Description: claimInput.Description,
	}

	id, err := h.claimUseCase.CreateClaim(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	output := dtos.CreateClaimOutputDTO{
		ID: id,
	}

	c.JSON(http.StatusCreated, output)
}

// GetClaim godoc
// @Summary Get a claim by ID
// @Description Get a claim by ID
// @Tags claims
// @Produce  json
// @Param id path string true "Claim ID"
// @Success 200 {object} domain.Claim
// @Security ApiKeyAuth
// @Router /claims/{id} [get]
func (h *ClaimHandler) GetClaim(c *gin.Context) {
	id := c.Param("id")

	claim, err := h.claimUseCase.ReadClaim(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, claim)
}

// UpdateClaim godoc
// @Summary Update a claim
// @Description Update a claim with the input payload
// @Tags claims
// @Accept  json
// @Produce  json
// @Param id path string true "Claim ID"
// @Param claim body dtos.UpdateClaimInputDTO true "Update Claim"
// @Success 200 {object} dtos.UpdateClaimOutputDTO
// @Security ApiKeyAuth
// @Router /claims/{id} [put]
func (h *ClaimHandler) UpdateClaim(c *gin.Context) {
	id := c.Param("id")

	var claimInput dtos.UpdateClaimInputDTO
	if err := c.ShouldBindJSON(&claimInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := usecases.UpdateClaimInput{
		Name:        claimInput.Name,
		Description: claimInput.Description,
	}

	claim, err := h.claimUseCase.UpdateClaim(c.Request.Context(), id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	output := dtos.UpdateClaimOutputDTO{
		ID: claim.ID,
	}

	c.JSON(http.StatusOK, output)
}

// DeleteClaim godoc
// @Summary Delete a claim
// @Description Delete a claim by ID
// @Tags claims
// @Param id path string true "Claim ID"
// @Success 200 {object} object{message=string}
// @Security ApiKeyAuth
// @Router /claims/{id} [delete]
func (h *ClaimHandler) DeleteClaim(c *gin.Context) {
	id := c.Param("id")

	err := h.claimUseCase.DeleteClaim(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Claim deleted successfully"})
}

// GetAllClaims godoc
// @Summary Get all claims
// @Description Get all claims
// @Tags claims
// @Produce  json
// @Success 200 {array} domain.Claim
// @Security ApiKeyAuth
// @Router /claims [get]
func (h *ClaimHandler) GetAllClaims(c *gin.Context) {
	claims, err := h.claimUseCase.GetAllClaims(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, claims)
}
