package main

import (
	"github.com/Gayana5/social-network/db-service/database"
	"github.com/Gayana5/social-network/db-service/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()
	routes.RegisterRoutes(r)

	err := r.Run(":8082")
	if err != nil {
		return
	}
}
