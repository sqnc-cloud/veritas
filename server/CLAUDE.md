# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Veritas is a GoLang API template built with Ports and Adapters (Hexagonal) architecture. It provides user authentication with JWT, MongoDB integration, and Swagger documentation via the Gin web framework.

## Common Commands

**Development:**
```bash
# Run with hot reload
make dev

# Build and run
make run

# Run tests
make test

# Generate Swagger docs (required after API changes)
make docs
```

**Docker:**
```bash
# Start all services (API + MongoDB)
docker-compose up --build

# The API runs on http://localhost:8080
# Swagger UI available at http://localhost:8080/swagger/index.html
```

**Dependencies:**
```bash
# Install/update Go modules
go mod tidy
```

## Architecture

The codebase follows **Ports and Adapters (Hexagonal Architecture)**:

### Core Layer (Business Logic)
- **`core/domain/`** - Domain entities (e.g., User struct)
- **`core/usecases/`** - Application business rules and orchestration
  - Usecases depend on output ports (interfaces) not concrete implementations
  - Example: `UserUsecase` depends on `UserOutputPort` interface

### Ports Layer (Interfaces)
- **`internal/ports/input/`** - Input port interfaces (incoming to core)
- **`internal/ports/output/`** - Output port interfaces (outgoing from core)
  - Example: `UserOutputPort` defines repository contract
- **`internal/ports/dtos/`** - Data Transfer Objects for API boundaries

### Adapters Layer (Implementations)
- **`internal/adapters/db/`** - Database implementations of output ports
  - Example: `UserRepository` implements `UserOutputPort`
- **`internal/handlers/`** - HTTP handlers (Gin controllers)
- **`internal/middleware/`** - Gin middleware (auth, etc.)
- **`internal/routes/`** - Route definitions

### Infrastructure
- **`cmd/api/main.go`** - Application entry point, dependency injection
- **`config/`** - Configuration (MongoDB connection, env vars)

## Key Architectural Patterns

**Dependency Flow:**
```
Handlers → Usecases → Output Ports ← Repositories
         (input)    (interface)    (implementation)
```

**Dependency Injection:**
Main constructs concrete implementations and injects them:
```go
userRepository := db.NewUserRepository(client.Database(dbName))
userUsecase := usecases.NewUserUsecase(userRepository)
userHandler := handlers.NewUserHandler(*userUsecase)
```

## Important Notes

- **Swagger:** Run `make docs` after modifying API endpoints/handlers. Swagger annotations are in handlers and `cmd/api/main.go`
- **JWT Secret:** Currently hardcoded in `internal/middleware/auth.go` as `"your_secret_key"` - should be moved to env var
- **Password Hashing:** Not implemented (marked with TODO comments) - passwords stored in plain text
- **MongoDB Collection:** Users stored in `users` collection (defined in `internal/adapters/db/user_repository.go`)
- **Environment Variables:**
  - `MONGO_URI` - MongoDB connection string (default: `mongodb://localhost:27017`)
  - `DB_NAME` - Database name (default: `housekeeper`)
  - `PORT` - Server port (default: `8080`)

## Testing

Tests are located in `core/usecases/` (e.g., `user_usecase_test.go`). Run with:
```bash
go test -v ./...
```

Run specific test:
```bash
go test -v ./core/usecases -run TestName
```

## API Authentication

Protected endpoints require JWT token in Authorization header:
```
Authorization: Bearer <token>
```

Get token from `POST /auth/login` after registering via `POST /auth/register`.
