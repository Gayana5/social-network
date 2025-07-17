package handlers

import (
	"github.com/Gayana5/social-network/db-service/database"
	"github.com/Gayana5/social-network/db-service/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	query := `INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id`
	err := database.DB.QueryRow(query, user.Username, user.Password).Scan(&user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func GetUserByUsername(c *gin.Context) {
	username := c.Param("username")

	var user models.User
	query := `SELECT id, username, password FROM users WHERE username = $1`
	err := database.DB.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
