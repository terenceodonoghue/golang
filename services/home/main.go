package main

import (
	"github.com/gin-gonic/gin"
	"github.com/terenceodonoghue/golang/services/home/internal/controllers"
)

func main() {
	r := gin.Default()
	api := r.Group("/api")
	{
		status := api.Group("/status")
		{
			status.GET("", controllers.GetStatus)
		}
	}
	r.Run()
}
