package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/terenceodonoghue/golang/services/home/internal/clients/fronius"
	"github.com/terenceodonoghue/golang/services/home/internal/clients/sensibo"
	"golang.org/x/sync/errgroup"
)

func GetStatus(c *gin.Context) {
	errs, ctx := errgroup.WithContext(c.Request.Context())

	ac := make(chan []sensibo.Device, 1)
	pv := make(chan fronius.Inverter, 1)

	errs.Go(func() error {
		defer close(ac)
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			s := sensibo.New(os.Getenv("SENSIBO_API_KEY"))
			return s.GetDevices(ac)
		}
	})

	errs.Go(func() error {
		defer close(pv)
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			f := fronius.New()
			return f.GetRealtimeData(pv)
		}
	})

	if err := errs.Wait(); err != nil {
		c.AbortWithError(http.StatusServiceUnavailable, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ac": <-ac,
		"pv": <-pv,
	})
}
