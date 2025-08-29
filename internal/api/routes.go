package api

import (
	"github.com/gin-gonic/gin"
	l "github.com/fernandokbs/goimage/internal/logger"
	"net/http"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", indexHandler)

	api := r.Group("/api")
	{
		api.POST("/upload", uploadHandler)
	}
}

func indexHandler(c *gin.Context) {
	l.LogInfo("Servidor iniciado", nil)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"status": "ONLINE 🚀",
	})
}

func uploadHandler(c *gin.Context) {
	file, err := c.FormFile("image")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no file"})
		return
	}

	l.LogInfo("Se ha cargado archivo", map[string]interface{}{
		"file": file.Filename,
	})

	// Guardar en carpeta uploads/
	path := "uploads/" + file.Filename
	if err := c.SaveUploadedFile(file, path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot save file"})
		return
	}

	// TODO: registrar en storage (SQLite o JSON)
	c.JSON(http.StatusOK, gin.H{"message": "file uploaded", "path": path})
}