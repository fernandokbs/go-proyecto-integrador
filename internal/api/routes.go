package api

import (
	"github.com/gin-gonic/gin"
	l "github.com/fernandokbs/goimage/internal/logger"
	"mime/multipart"
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
	type UploadInput struct {
		Action string                `form:"action" binding:"required"`
		File   *multipart.FileHeader `form:"image" binding:"required"`
	}

	var input UploadInput

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	file := input.File

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