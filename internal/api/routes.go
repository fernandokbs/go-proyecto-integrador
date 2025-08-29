package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", indexHandler)

	api := r.Group("/api")
	{
		api.POST("/upload", uploadHandler)
		api.GET("/images/:id", statusHandler)
	}
}

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"status": "ONLINE 🚀",
	})
}

func uploadHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no file"})
		return
	}

	// Guardar en carpeta uploads/
	path := "uploads/" + file.Filename
	if err := c.SaveUploadedFile(file, path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot save file"})
		return
	}

	// TODO: registrar en storage (SQLite o JSON)
	c.JSON(http.StatusOK, gin.H{"message": "file uploaded", "path": path})
}

func statusHandler(c *gin.Context) {
	id := c.Param("id")
	// TODO: consultar storage para ver el estado
	c.JSON(http.StatusOK, gin.H{"id": id, "status": "pending"})
}
