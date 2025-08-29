package main

import (
	"github.com/gin-gonic/gin"
	"github.com/fernandokbs/goimage/internal/api"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	api.RegisterRoutes(r)

	r.Run(":8080")
}
