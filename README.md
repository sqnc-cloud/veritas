# Veritas API Template

Veritas is a robust and scalable GoLang API template designed to kickstart your projects with built-in authentication, MongoDB integration, and a focus on the Ports and Adapters architectural pattern. It provides a solid foundation for developing secure and efficient backend services.

## Features

-   **User Authentication**: Secure user registration and login with JWT (JSON Web Tokens).
-   **Role-Based Access Control**: (Optional, can be extended) Foundation for implementing different user roles.
-   **MongoDB Integration**: Seamless integration with MongoDB for data persistence.
-   **Ports and Adapters (Hexagonal Architecture Inspired)**: Organized codebase with clear separation of concerns using the ports and adapters pattern.
-   **Gin Web Framework**: Fast and lightweight web framework for building APIs.
-   **Dockerized Development**: Easy setup and deployment using Docker and Docker Compose.
-   **Swagger Documentation**: Automatically generated API documentation for easy testing and understanding.

## Technologies Used

-   **GoLang**: Backend programming language.
-   **Gin Gonic**: High-performance HTTP web framework.
-   **MongoDB**: NoSQL database.
-   **JWT (JSON Web Tokens)**: For secure authentication.
-   **Docker & Docker Compose**: For containerization and orchestration.
-   **Swag**: For automatic Swagger API documentation generation.

## Getting Started

### Prerequisites

Before you begin, ensure you have the following installed:

-   [Go](https://golang.org/doc/install) (1.20 or higher)
-   [Docker](https://docs.docker.com/get-docker/)
-   [Docker Compose](https://docs.docker.com/compose/install/)

### Installation

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/your-username/veritas.git
    cd veritas
    ```

2.  **Build and run the application with Docker Compose:**
    ```bash
    docker-compose up --build
    ```
    This command will:
    -   Build the Docker images for the API and MongoDB.
    -   Start the MongoDB container.
    -   Start the Veritas API container.

    The API will be accessible at `http://localhost:8080`.

### Running Locally (without Docker)

1.  **Install Go modules:**
    ```bash
    go mod tidy
    ```

2.  **Generate Swagger documentation:**
    ```bash
    make docs
    ```

3.  **Run the application:**
    ```bash
    make run
    ```
    Ensure you have a MongoDB instance running and accessible at `mongodb://localhost:27017` or configure the `MONGODB_URI` environment variable accordingly.

## API Endpoints

Once the application is running, you can access the API documentation via Swagger UI at `http://localhost:8080/swagger/index.html`.

Key authentication endpoints:

-   `POST /auth/register`: Register a new user.
-   `POST /auth/login`: Authenticate a user and receive a JWT.

User management endpoints (require authentication):

-   `GET /users`: Get all users.
-   `GET /users/{id}`: Get a user by ID.
-   `PUT /users/{id}`: Update a user by ID.
-   `DELETE /users/{id}`: Delete a user by ID.

## Project Structure

```
. 
├── cmd/             # Main application entry points
│   └── api/         # API server entry point
├── config/          # Configuration files (e.g., database connection)
├── core/            # Core business logic
│   ├── domain/      # Domain entities and interfaces
│   └── usecases/    # Application-specific business rules
├── docs/            # Swagger documentation files
├── internal/        # Internal implementation details
│   ├── adapters/    # Implementations of ports (e.g., database adapters)
│   │   └── db/      # Database repository implementations
│   ├── handlers/    # HTTP request handlers
│   ├── middleware/  # Gin middleware (e.g., authentication middleware)
│   ├── ports/       # Interfaces for external dependencies (input/output ports)
│   │   ├── dtos/    # Data Transfer Objects
│   │   ├── input/   # Input port interfaces
│   │   └── output/  # Output port interfaces
│   └── routes/      # API route definitions
├── Dockerfile       # Docker build instructions
├── docker-compose.yml # Docker Compose configuration
├── go.mod           # Go modules file
├── go.sum           # Go modules checksums
├── Makefile         # Build and development scripts
└── README.md        # Project README
```

## Contributing

Contributions are welcome! Please feel free to open issues or submit pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details. (You might want to create a LICENSE file if you don't have one.)