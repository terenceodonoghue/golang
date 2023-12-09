package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Climate(rg *gin.RouterGroup) {
	rg.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
