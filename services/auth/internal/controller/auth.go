package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/terenceodonoghue/golang/libs/jwt"
	"github.com/terenceodonoghue/golang/services/auth/internal/database"
)

const (
	access  = "access_token"
	refresh = "refresh_token"
)

func Login(c *gin.Context, conn *pgx.Conn) {
	var request LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if exists, err := database.Exists(conn, database.UserCredentials, "email_address", request.EmailAddress); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	} else if !exists {
		c.Status(http.StatusUnauthorized)
		return
	}

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

func RefreshToken(c *gin.Context, conn *pgx.Conn) {
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

type LoginRequest struct {
	EmailAddress string `json:"email_address" binding:"required"`
	Password     string `json:"password" binding:"required"`
}
