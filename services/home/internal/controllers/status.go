package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/terenceodonoghue/golang/services/home/internal/clients/fronius"
	"github.com/terenceodonoghue/golang/services/home/internal/clients/sensibo"
)

func GetStatus(c *gin.Context) {
	f := fronius.New()
	s := sensibo.New(os.Getenv("SENSIBO_API_KEY"))

	ac, err := s.GetDevices()
	if err != nil {
		return
	}

	pv, err := f.GetRealtimeData()
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ac": ac,
		"pv": pv,
	})
}
