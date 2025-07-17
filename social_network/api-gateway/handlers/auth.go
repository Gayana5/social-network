package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func forwardRequestToProcessor(path string, body interface{}) (*http.Response, error) {
	jsonData, _ := json.Marshal(body)
	return http.Post("http://processor:8081"+path, "application/json", bytes.NewBuffer(jsonData))
}

func SignUp(c *gin.Context) {
	var req AuthRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	resp, err := forwardRequestToProcessor("/register", req)
	if err != nil {
		c.JSON(500, gin.H{"error": "processor unavailable"})
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	c.Data(resp.StatusCode, "application/json", body)
}

func SignIn(c *gin.Context) {
	var req AuthRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	resp, err := forwardRequestToProcessor("/login", req)
	if err != nil {
		c.JSON(500, gin.H{"error": "processor unavailable"})
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	c.Data(resp.StatusCode, "application/json", body)
}
