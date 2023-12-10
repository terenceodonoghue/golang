package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/terenceodonoghue/golang/services/home/internal/clients/sensibo"
)

func GetStatus(c *gin.Context) {
	s := sensibo.New(os.Getenv("SENSIBO_API_KEY"))
	ac, err := s.GetDevices()
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"air_conditioning": ac,
	})
}
