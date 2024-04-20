package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/shatshai/go-explorer/gin-user-api/database"
    "github.com/shatshai/go-explorer/gin-user-api/handlers"
    "github.com/joho/godotenv"
    "os"
)

func init() {
	// Check if the .env file exists
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		fmt.Println("Warning: .env file not found, loading default environment variables")
		return
	}

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

func main() {
    // Initialize database connection
    database.InitDB()
    defer database.DB.Close()

    // Initialize Gin router
    r := gin.Default()

    // API endpoints for user CRUD operations
    r.GET("/user/:id", handlers.GetUser)
    r.POST("/user", handlers.CreateUser)
    r.PUT("/user/:id", handlers.UpdateUser)
    r.DELETE("/user/:id", handlers.DeleteUser)

    // Run the server
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Default port if not specified in .env
    }
    r.Run(fmt.Sprintf(":%s", port))
}