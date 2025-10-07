package main

import (
	"context"
	"log"
	"os"
	"time"
	"veritas/config"
	"veritas/core/usecases"
	"veritas/internal/adapters/db"
	"veritas/internal/handlers"
	"veritas/internal/routes"

	_ "veritas/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Veritas API
// @version 1.0
// @description This is a sample server for a auth managing API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file, proceeding with environment variables")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := config.GetMongoDBClient(ctx)
	if err != nil {
		log.Fatalf("failed to connect to mongodb: %v", err)
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()

	dbName := config.GetDatabaseName()
	userRepository := db.NewUserRepository(client.Database(dbName))
	userUsecase := usecases.NewUserUsecase(userRepository)
	userHandler := handlers.NewUserHandler(*userUsecase)
	authHandler := handlers.NewAuthHandler(*userUsecase)

	roleRepository := db.NewRoleRepository(client.Database(dbName))
	roleUsecase := usecases.NewRoleUsecase(roleRepository)
	roleHandler := handlers.NewRoleHandler(*roleUsecase)

	claimRepository := db.NewClaimRepository(client.Database(dbName))
	claimUsecase := usecases.NewClaimUsecase(claimRepository)
	claimHandler := handlers.NewClaimHandler(*claimUsecase)

	routes.SetupUserRoutes(router, userHandler)
	routes.SetupAuthRoutes(router, authHandler)
	routes.SetupRoleRoutes(router, roleHandler)
	routes.SetupClaimRoutes(router, claimHandler)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Printf("Server listening on port %s", port)
	err = router.Run(":" + port)
	if err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
