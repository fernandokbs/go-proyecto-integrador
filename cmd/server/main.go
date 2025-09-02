package main

import (
	"github.com/gin-gonic/gin"
	"github.com/fernandokbs/goimage/internal/api"
	"github.com/fernandokbs/goimage/internal/database"
)

func main() {
	database.Connect()
	
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	api.RegisterRoutes(r)

	r.Run(":8080")
}
