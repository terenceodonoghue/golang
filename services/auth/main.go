package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/terenceodonoghue/golang/services/auth/internal/controllers"
	"github.com/terenceodonoghue/golang/services/auth/internal/database"
	"github.com/terenceodonoghue/golang/services/auth/internal/middleware"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close(context.Background())

	r := gin.Default()
	api := r.Group("/api")
	{
		api.Use(middleware.CORS())
		{
			api.POST("/login", func(c *gin.Context) {
				controllers.Login(c, db)
			})
			api.GET("/refresh_token", func(c *gin.Context) {
				controllers.RefreshToken(c, db)
			})
		}
	}
	r.Run()
}
