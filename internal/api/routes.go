package api

import (
	"github.com/gin-gonic/gin"
	l "github.com/fernandokbs/goimage/internal/logger"
	"github.com/fernandokbs/goimage/internal/images"
	"net/http"
	"time"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", indexHandler)

	api := r.Group("/api")
	{
		api.POST("/upload", uploadHandler)
	}
}

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func uploadHandler(c *gin.Context) {
	form, _ := c.MultipartForm()
	
	files := form.File["images[]"]

	for _, file := range files { 
		l.LogInfo("Se ha cargado archivo", map[string]interface{}{
			"file": file.Filename,
		})
		
		path := "uploads/" + file.Filename
		
		if err := c.SaveUploadedFile(file, path); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot save file"})
			return
		}

		processor, _ := images.NewProcessor(path)

		processor.Watermark("PRUEBA DESDE ENDPOINT")

		time.Sleep(2 * time.Second)
	}

	c.JSON(http.StatusOK, gin.H{"message": "files uploaded"})
}