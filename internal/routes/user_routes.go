package routes

import (
	"veritas/internal/handlers"
	"veritas/internal/middleware"

	"github.com/gin-gonic/gin"
)

// SetupUserRoutes sets up the user routes.
func SetupUserRoutes(router *gin.Engine, handler *handlers.UserHandler) {
	userRoutes := router.Group("/users")
	userRoutes.Use(middleware.AuthMiddleware())
	{
		userRoutes.GET("", handler.GetAllUsers)
		userRoutes.GET("/:id", handler.GetUser)
		userRoutes.PUT("/:id", handler.UpdateUser)
		userRoutes.DELETE("/:id", handler.DeleteUser)
	}
}
