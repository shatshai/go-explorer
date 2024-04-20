package handlers

import (
    "github.com/gin-gonic/gin"
    "github.com/shatshai/go-explorer/gin-user-api/database"
    "github.com/shatshai/go-explorer/gin-user-api/models"
    "net/http"
    "strconv"
)

// GetUser fetches a user by ID from the database
func GetUser(c *gin.Context) {
    // Get user ID from URL parameter
    userIDStr := c.Param("id")
    userID, err := strconv.ParseUint(userIDStr, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    // Fetch user from the database
    user, err := database.GetUserByID(userID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    // Return user data in JSON response
    c.JSON(http.StatusOK, gin.H{"data": user})
}

// CreateUser creates a new user in the database
func CreateUser(c *gin.Context) {
    // Bind request body to User struct
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }

    // Insert new user into the database
    if err := database.CreateUser(&user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }

    // Return success message
    c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "data": user})
}

// UpdateUser updates an existing user in the database
func UpdateUser(c *gin.Context) {
    // Get user ID from URL parameter
    userIDStr := c.Param("id")
    userID, err := strconv.ParseUint(userIDStr, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    // Bind request body to User struct
    var updatedUser models.User
    if err := c.ShouldBindJSON(&updatedUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }

    // Update user in the database
    if err := database.UpdateUser(userID, &updatedUser); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
        return
    }

    // Return success message
    c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "data": updatedUser})
}

// DeleteUser deletes a user from the database
func DeleteUser(c *gin.Context) {
    // Get user ID from URL parameter
    userIDStr := c.Param("id")
    userID, err := strconv.ParseUint(userIDStr, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    // Delete user from the database
    if err := database.DeleteUser(userID); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
        return
    }

    // Return success message
    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}