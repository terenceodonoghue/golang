package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/terenceodonoghue/golang/services/auth/internal/jwt"
)

func Login(c *gin.Context) {
	tokenString, err := jwt.CreateToken(time.Now)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	maxAge := time.Until(time.Now().AddDate(0, 0, 14)).Seconds()

	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie("refresh_token", tokenString, int(maxAge), "", "", true, true)
	c.Status(http.StatusOK)
}
