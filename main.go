package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// User представляет модель пользователя
type User struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func main() {
	r := gin.Default()

	// Список пользователей
	var users []User

	// GET /api/v1/users - Получение списка пользователей
	r.GET("/api/v1/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, users)
	})

	// POST /api/v1/register - Регистрация нового пользователя
	r.POST("/api/v1/register", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}
		// Добавляем пользователя в список
		users = append(users, user)
		c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
	})

	// Запуск сервера на порту 8080
	r.Run(":8080")
}
