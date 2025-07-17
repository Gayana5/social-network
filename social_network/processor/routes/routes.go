package routes

import (
	"github.com/Gayana5/social-network/processor/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/signUp", handlers.SignUp)
	r.POST("/signIn", handlers.SignIn)
	r.POST("/post", handlers.CreatePost)
	r.GET("/posts", handlers.GetAllPosts)
	r.GET("/posts/:userId", handlers.GetPostsByUser)

	r.POST("/post/:id/like", handlers.LikePost)
	r.POST("/post/:id/comment", handlers.CommentPost)
}
