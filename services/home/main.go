package main

import (
	"github.com/gin-gonic/gin"
	"github.com/terenceodonoghue/golang/services/home/internal/routes"
)

func main() {
	r := gin.Default()
	routes.Climate(r.Group("/climate"))
	r.Run()
}
