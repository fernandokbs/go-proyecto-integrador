package main

import (
	_ "fmt"
	"github.com/fernandokbs/goimage/internal/images"
)

func main() {
	processor, _ := images.NewProcessor("133968927133882801.jpg")

	processor.Watermark("Hola")
}
