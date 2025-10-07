package routes

import (
	"veritas/internal/handlers"
	"veritas/internal/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRoleRoutes sets up the role routes.
func SetupRoleRoutes(router *gin.Engine, handler *handlers.RoleHandler) {
	roleRoutes := router.Group("/roles")
	roleRoutes.Use(middleware.AuthMiddleware())
	{
		roleRoutes.POST("", handler.CreateRole)
		roleRoutes.GET("", handler.GetAllRoles)
		roleRoutes.GET("/:id", handler.GetRole)
		roleRoutes.PUT("/:id", handler.UpdateRole)
		roleRoutes.DELETE("/:id", handler.DeleteRole)
	}
}
