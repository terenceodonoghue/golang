package main

import (
	"github.com/gin-gonic/gin"
	"github.com/terenceodonoghue/golang/services/auth/internal/controllers"
)

func main() {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.POST("/login", controllers.Login)
	}
	r.Run()
}
