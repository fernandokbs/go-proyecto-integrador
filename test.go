package main

import (
	_ "fmt"
	"github.com/fernandokbs/goimage/internal/images"
)

func main() {
	processor, _ := images.NewProcessor("/home/fernando/proyecto-integrador-go/133968927133882801.jpg")

	processor.Thumbnail()
}
