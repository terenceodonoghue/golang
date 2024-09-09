package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/terenceodonoghue/golang/services/auth/internal/controller"
	"github.com/terenceodonoghue/golang/services/auth/internal/database"
	"github.com/terenceodonoghue/golang/services/auth/internal/middleware"
)

func main() {
	db, err := database.New()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close(context.Background())

	r := gin.Default()
	api := r.Group("/api")
	{
		api.Use(middleware.CORS())
		auth := api.Group("/auth")
		{
			auth.POST("/login", func(c *gin.Context) {
				controller.Login(c, db)
			})
			auth.GET("/refresh_token", func(c *gin.Context) {
				controller.RefreshToken(c, db)
			})
		}
	}
	r.Run()
}
