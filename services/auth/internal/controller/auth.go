package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/terenceodonoghue/golang/libs/jwt"
	"github.com/terenceodonoghue/golang/services/auth/internal/database"
)

const (
	access  = "access_token"
	refresh = "refresh_token"
)

func Login(conn *pgx.Conn) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var request LoginRequest

		if err := decoder.Decode(&request); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if exists, err := database.Exists(conn, database.UserCredentials, "email_address", request.EmailAddress); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else if !exists {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		expiry := time.Now().Add(10 * time.Minute)
		maxAge := time.Now().AddDate(0, 0, 14)

		token, err := jwt.CreateToken(expiry)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if token, err := jwt.CreateToken(maxAge); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else {
			http.SetCookie(w, &http.Cookie{
				Name:     refresh,
				Value:    token,
				MaxAge:   int(time.Until(maxAge).Seconds()),
				Path:     "/api/auth/refresh_token",
				Domain:   "",
				SameSite: http.SameSiteStrictMode,
				Secure:   true,
				HttpOnly: true,
			})
		}

		json.NewEncoder(w).Encode(map[string]string{access: token})
	})

}

func RefreshToken(conn *pgx.Conn) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(refresh)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err = jwt.VerifyToken(cookie.Value); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		expiry := time.Now().Add(10 * time.Minute)
		token, err := jwt.CreateToken(expiry)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{access: token})
	})
}

type LoginRequest struct {
	EmailAddress string `json:"email_address" binding:"required"`
	Password     string `json:"password" binding:"required"`
}
