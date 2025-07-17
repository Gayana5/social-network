package main

import (
	"github.com/Gayana5/social-network/processor/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.RegisterRoutes(r)
	err := r.Run(":8081")
	if err != nil {
		return
	}
}
