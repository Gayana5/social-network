package handlers

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func proxyJSON(c *gin.Context, method, path string) {
	body, _ := io.ReadAll(c.Request.Body)
	req, err := http.NewRequest(method, "http://processor:8081"+path, bytes.NewReader(body))
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to create request"})
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(500, gin.H{"error": "processor unavailable"})
		return
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	c.Data(resp.StatusCode, "application/json", respBody)
}

func proxyGET(c *gin.Context, path string) {
	resp, err := http.Get("http://processor:8081" + path)
	if err != nil {
		c.JSON(500, gin.H{"error": "processor unavailable"})
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	c.Data(resp.StatusCode, "application/json", body)
}

func CreatePost(c *gin.Context) {
	proxyJSON(c, "POST", "/post")
}

func GetAllPosts(c *gin.Context) {
	proxyGET(c, "/posts")
}

func GetUserPosts(c *gin.Context) {
	userId := c.Param("userId")
	proxyGET(c, "/posts/"+userId)
}

func LikePost(c *gin.Context) {
	postId := c.Param("id")
	proxyJSON(c, "POST", "/post/"+postId+"/like")
}

func CommentPost(c *gin.Context) {
	postId := c.Param("id")
	proxyJSON(c, "POST", "/post/"+postId+"/comment")
}
