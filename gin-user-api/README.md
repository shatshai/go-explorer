# Go User CRUD Service with Docker

This project is a CRUD (Create, Read, Update, Delete) user service implemented in Go, with MySQL as the database and Docker for containerization.

## Prerequisites

Before running the application, make sure you have the following installed:
- Docker
- Go (if you want to build outside Docker)
- MySQL (for database)

## Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/go-user-crud.git
   cd go-user-crud
   ```
2. Set up the environment variables:

    Copy the .env.example file to .env and configure the database connection settings.

3. Start the application using Docker Compose:
    ```
    docker-compose up
    ```
4. Access the API at http://localhost:8080.

Project Structure

    main.go: Entry point of the application.
    handlers/: Contains HTTP request handlers for user CRUD operations.
    models/: Defines the User model and database operations.
    database/: Handles database initialization and connection.
    Dockerfile: Docker configuration for building the application image.
    docker-compose.yaml: Docker Compose configuration for running the application and MySQL database.

API Endpoints

    GET /users: Get all users.
    GET /users/{id}: Get a user by ID.
    POST /users: Create a new user.
    PUT /users/{id}: Update an existing user.
    DELETE /users/{id}: Delete a user by ID.

Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.
License

This project is licensed under the MIT License - see the LICENSE file for details.

```
Replace placeholders like `your-username` with your actual GitHub username or project-specific information. This README.md provides an overview of the project, setup instructions, project structure, API endpoints, contribution guidelines, and license information. Adjust it as needed based on your project's specifics.
```