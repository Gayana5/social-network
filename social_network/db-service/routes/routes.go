package routes

import (
	"github.com/Gayana5/social-network/db-service/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/users", handlers.CreateUser)
	r.GET("/users/:username", handlers.GetUserByUsername)

	r.POST("/post", handlers.CreatePost)
	r.GET("/posts", handlers.GetAllPosts)
	r.GET("/posts/:userId", handlers.GetPostsByUser)

	r.POST("/post/:id/like", handlers.LikePost)
	r.POST("/post/:id/comment", handlers.CommentPost)
}
