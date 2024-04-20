package models

type User struct {
    ID       uint   `json:"id"`
    Username string `json:"username" binding:"required"`
    Email    string `json:"email" binding:"required,email"`
}
