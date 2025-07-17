package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/Gayana5/social-network/processor/utils"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SignUp(c *gin.Context) {
	var req AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	hashed, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to hash password"})
		return
	}

	req.Password = hashed
	jsonData, _ := json.Marshal(req)

	resp, err := http.Post("http://db-service:8082/users", "application/json", bytes.NewReader(jsonData))
	if err != nil {
		c.JSON(500, gin.H{"error": "db-service unavailable"})
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	body, _ := io.ReadAll(resp.Body)

	c.Data(resp.StatusCode, "application/json", body)
}

func SignIn(c *gin.Context) {
	var req AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	// запрашиваем пользователя из БД
	resp, err := http.Get("http://db-service:8082/users/" + req.Username)
	if err != nil {
		c.JSON(500, gin.H{"error": "db-service unavailable"})
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)

	if resp.StatusCode != 200 {
		c.JSON(401, gin.H{"error": "user not found"})
		return
	}

	var dbUser AuthRequest
	err = json.NewDecoder(resp.Body).Decode(&dbUser)
	if err != nil {
		return
	}

	if !utils.CheckPassword(dbUser.Password, req.Password) {
		c.JSON(401, gin.H{"error": "invalid credentials"})
		return
	}

	c.JSON(200, gin.H{"status": "login successful"})
}
