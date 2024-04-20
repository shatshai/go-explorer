package database

import (
    "database/sql"
    "fmt"
    "os"
	"github.com/shatshai/go-explorer/gin-user-api/models"
    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
    dbUsername := os.Getenv("DB_USERNAME")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)
	fmt.Println(dataSourceName)
    
    db, err := sql.Open("mysql", dataSourceName)
    if err != nil {
        panic(err.Error())
    }

    if err := db.Ping(); err != nil {
        panic(err.Error())
    }

    fmt.Println("Connected to MySQL database")
    DB = db
}

// GetUserByID fetches a user by ID from the database
func GetUserByID(userID uint64) (*models.User, error) {
    // Prepare the SQL statement
    query := "SELECT id, username, email FROM users WHERE id = ?"
    row := DB.QueryRow(query, userID)

    // Create a new User object to store the fetched user data
    var user models.User
    err := row.Scan(&user.ID, &user.Username, &user.Email)
    if err != nil {
        return nil, err // Return nil user and error if user not found or other error occurs
    }

    return &user, nil // Return fetched user and nil error
}

// CreateUser inserts a new user into the database
func CreateUser(user *models.User) error {
    // Prepare the SQL statement
    query := "INSERT INTO users (username, email) VALUES (?, ?)"
    stmt, err := DB.Prepare(query)
    if err != nil {
        return err // Return error if preparing the statement fails
    }
    defer stmt.Close()

    // Execute the SQL statement to insert the user
    result, err := stmt.Exec(user.Username, user.Email)
    if err != nil {
        return err // Return error if executing the statement fails
    }

    // Check the number of rows affected to ensure the user was inserted successfully
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err // Return error if getting rows affected fails
    }
    if rowsAffected == 0 {
        return fmt.Errorf("no rows affected, user not inserted") // Return error if no rows were affected
    }

    return nil // Return nil error if user insertion is successful
}

// UpdateUser updates an existing user in the database
func UpdateUser(userID uint64, updatedUser *models.User) error {
    // Prepare the SQL statement
    query := "UPDATE users SET username = ?, email = ? WHERE id = ?"
    stmt, err := DB.Prepare(query)
    if err != nil {
        return err // Return error if preparing the statement fails
    }
    defer stmt.Close()

    // Execute the SQL statement to update the user
    _, err = stmt.Exec(updatedUser.Username, updatedUser.Email, userID)
    if err != nil {
        return err // Return error if executing the statement fails
    }

    return nil // Return nil error if user update is successful
}

// DeleteUser deletes a user from the database
func DeleteUser(userID uint64) error {
    // Prepare the SQL statement
    query := "DELETE FROM users WHERE id = ?"
    stmt, err := DB.Prepare(query)
    if err != nil {
        return err // Return error if preparing the statement fails
    }
    defer stmt.Close()

    // Execute the SQL statement to delete the user
    _, err = stmt.Exec(userID)
    if err != nil {
        return err // Return error if executing the statement fails
    }

    return nil // Return nil error if user deletion is successful
}