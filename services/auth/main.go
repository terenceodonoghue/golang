package main

import (
	"context"
	"log"
	"net/http"

	"github.com/terenceodonoghue/golang/services/auth/internal/controller"
	"github.com/terenceodonoghue/golang/services/auth/internal/database"
	"github.com/terenceodonoghue/golang/services/auth/internal/middleware"
)

func main() {
	db, err := database.New()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close(context.Background())

	r := http.NewServeMux()
	r.HandleFunc("POST /api/auth/login", middleware.CORS(controller.Login(db)))
	r.HandleFunc("GET /api/auth/refresh_token", middleware.CORS(controller.RefreshToken(db)))

	err = http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal(err)
	}
}
