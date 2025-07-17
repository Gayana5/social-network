package handlers

import (
	"github.com/Gayana5/social-network/db-service/database"
	"github.com/Gayana5/social-network/db-service/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	query := `INSERT INTO posts (user_id, content, likes) VALUES ($1, $2, 0) RETURNING id`
	err := database.DB.QueryRow(query, post.UserID, post.Content).Scan(&post.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(http.StatusCreated, post)
}

func GetAllPosts(c *gin.Context) {
	rows, err := database.DB.Query(`SELECT id, user_id, content, likes FROM posts ORDER BY id DESC`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var p models.Post
		rows.Scan(&p.ID, &p.UserID, &p.Content, &p.Likes)
		posts = append(posts, p)
	}

	c.JSON(http.StatusOK, posts)
}

func GetPostsByUser(c *gin.Context) {
	userId := c.Param("userId")

	rows, err := database.DB.Query(`SELECT id, user_id, content, likes FROM posts WHERE user_id = $1`, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var p models.Post
		rows.Scan(&p.ID, &p.UserID, &p.Content, &p.Likes)
		posts = append(posts, p)
	}

	c.JSON(http.StatusOK, posts)
}

func LikePost(c *gin.Context) {
	postId := c.Param("id")

	_, err := database.DB.Exec(`UPDATE posts SET likes = likes + 1 WHERE id = $1`, postId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "liked"})
}

func CommentPost(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	query := `INSERT INTO comments (post_id, user_id, content) VALUES ($1, $2, $3) RETURNING id`
	err := database.DB.QueryRow(query, comment.PostID, comment.UserID, comment.Content).Scan(&comment.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to comment"})
		return
	}

	c.JSON(http.StatusCreated, comment)
}
