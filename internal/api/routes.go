package api

import (
	"github.com/gin-gonic/gin"
	l "github.com/fernandokbs/goimage/internal/logger"
	"github.com/fernandokbs/goimage/internal/images"
	"github.com/fernandokbs/goimage/internal/models"
	"github.com/fernandokbs/goimage/internal/database"
	"net/http"
	"time"
	_ "fmt"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", indexHandler)

	api := r.Group("/api")
	{
		api.POST("/upload", uploadHandler)
	}

	r.Static("/files", "./processed_files")
}

func indexHandler(c *gin.Context) {
	db, _ := database.GetConnection() // Handle el error

	var imageRecords []models.Image
	
	db.Find(&imageRecords)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"imageRecords": imageRecords,
	})
}

func uploadHandler(c *gin.Context) {
	form, _ := c.MultipartForm()
	
	files := form.File["images[]"]

	done := make(chan string, len(files))

	var processedLinks []string

	for _, file := range files { 
		l.LogInfo("Se ha cargado archivo", map[string]interface{}{
			"file": file.Filename,
		})
		
		path := "uploads/" + file.Filename
		
		if err := c.SaveUploadedFile(file, path); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot save file"})
			return
		}

		go func(p string) {
			processor, err := images.NewProcessor(path)

			if err != nil {
				l.LogInfo("Error procesando imagen", map[string]interface{}{
					"file": p,
					"error": err.Error(),
				})

				done <- p
				return
			}

			url, _ := processor.Watermark("PRUEBA DESDE ENDPOINT")

			time.Sleep(2 * time.Second) // Para simular que el proceso toma timpo

			l.LogInfo("Imagen procesada", map[string]interface{}{
				"file": url,
			})

			db, _ := database.GetConnection() // Handle el error

			db.Create(&models.Image{Url: url})

			done <- url
		}(path)
	}

	for i := 0; i < len(files); i++ {
		processedLinks = append(processedLinks, <-done)
	}

	c.JSON(http.StatusOK, gin.H{
		"files": processedLinks,
		"message": "files uploaded",
	})
}