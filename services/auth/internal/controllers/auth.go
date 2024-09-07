package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/terenceodonoghue/golang/libs/jwt"
)

const (
	access  = "access_token"
	refresh = "refresh_token"
)

func Login(c *gin.Context) {
	expiry := time.Now().Add(10 * time.Minute)
	maxAge := time.Now().AddDate(0, 0, 14)

	token, err := jwt.CreateToken(expiry)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	if token, err := jwt.CreateToken(maxAge); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	} else {
		c.SetSameSite(http.SameSiteStrictMode)
		c.SetCookie(refresh, token, int(time.Until(maxAge).Seconds()), "/api/refresh_token", "", true, true)
	}

	c.JSON(http.StatusOK, gin.H{
		access: token,
	})
}

func RefreshToken(c *gin.Context) {
	cookie, err := c.Cookie(refresh)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if err = jwt.VerifyToken(cookie); err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	expiry := time.Now().Add(10 * time.Minute)
	token, err := jwt.CreateToken(expiry)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		access: token,
	})
}
