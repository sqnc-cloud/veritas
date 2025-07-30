package routes

import (
	"veritas/internal/handlers"

	"github.com/gin-gonic/gin"
)

// SetupAuthRoutes sets up the auth routes.
func SetupAuthRoutes(router *gin.Engine, handler *handlers.AuthHandler) {
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/login", handler.Login)
		authRoutes.POST("/signup", handler.SignUp)
	}
}
