package main

import (
	"github.com/gin-gonic/gin"
	"github.com/terenceodonoghue/golang/services/auth/internal/controllers"
	"github.com/terenceodonoghue/golang/services/auth/internal/middleware"
)

func main() {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.Use(middleware.CORS())
		{
			api.POST("/login", controllers.Login)
			api.GET("/refresh_token", controllers.RefreshToken)
		}
	}
	r.Run()
}
