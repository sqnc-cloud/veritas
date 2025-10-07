package routes

import (
	"veritas/internal/handlers"
	"veritas/internal/middleware"

	"github.com/gin-gonic/gin"
)

// SetupClaimRoutes sets up the claim routes.
func SetupClaimRoutes(router *gin.Engine, handler *handlers.ClaimHandler) {
	claimRoutes := router.Group("/claims")
	claimRoutes.Use(middleware.AuthMiddleware())
	{
		claimRoutes.POST("", handler.CreateClaim)
		claimRoutes.GET("", handler.GetAllClaims)
		claimRoutes.GET("/:id", handler.GetClaim)
		claimRoutes.PUT("/:id", handler.UpdateClaim)
		claimRoutes.DELETE("/:id", handler.DeleteClaim)
	}
}
