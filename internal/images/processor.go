package images

import (
	"fmt"
	"os"
	"github.com/h2non/bimg"
	"path/filepath"
	log "github.com/fernandokbs/goimage/internal/logger"
)

type ImageProcessor struct {
	FileName string
	Path string
	Buffer []byte
}

func NewProcessor(path string) (*ImageProcessor, error) {
	buffer, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	return &ImageProcessor{
		FileName: path,
		Path:   path,
		Buffer: buffer,
	}, nil
}

func (p *ImageProcessor) Watermark(text string) (string, error) {
	watermark := bimg.Watermark{
		Text:       text,
		Opacity:    0.25,
		Width:      200,
		DPI:        100,
		Margin:     150,
		Font:       "sans bold 12",
		Background: bimg.Color{255, 255, 255},
	}

	newImage, err := bimg.NewImage(p.Buffer).Watermark(watermark)

	if err != nil {
		return "", err
	}

	return p.Save(newImage)
}

func (p *ImageProcessor) Save(newImage []byte) (string, error) {
	s3Client, err := NewS3Client()

	localPath := filepath.Join("processed_files", filepath.Base(p.FileName))

	if err := bimg.Write(localPath, newImage); err != nil {
		log.LogError("error guardando localmente: %w", map[string]interface{}{
			"error": err,
		})
		return "", err
	}

	key := fmt.Sprintf("imagenes/%s", filepath.Base(p.FileName))

	url, err := s3Client.Upload(localPath, key)

	if err != nil {
		log.LogError("error subiendo a S3:", map[string]interface{}{
			"error": err,
		})
		return "", err
	}
	
	return url, nil
}