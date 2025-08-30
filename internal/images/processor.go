package images

import (
	"fmt"
	"os"
	"github.com/h2non/bimg"
	"path/filepath"
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

func (p *ImageProcessor) Watermark(text string) (error) {
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
		fmt.Println(err);
		return err
	}

	return p.Save(newImage)
}

func (p *ImageProcessor) Thumbnail() (error) {
	options := bimg.Options{
		Width:   200,
		Height:  200,
		Crop:    true, 
		Quality: 90, 
	}

	newImage, err := bimg.NewImage(buffer).Process(options)
	
	if err != nil {
		fmt.Println(err);
		return err
	}

	return p.Save(newImage)
}

func (p *ImageProcessor) Save(newImage []byte) error {
	localPath := filepath.Join("processed_files", filepath.Base(p.FileName))
	return bimg.Write(localPath, newImage)
}