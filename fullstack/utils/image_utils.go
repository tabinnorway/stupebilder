package utils

import (
	"fmt"
	"log"

	"github.com/disintegration/imaging"
)

func ResizeImage(inputPath, outputPath string, maxLength int) error {
	// Open the source image
	fmt.Printf("Resizing from: %s\n", inputPath)
	fmt.Printf("           to: %s\n", outputPath)
	src, err := imaging.Open(inputPath)
	if err != nil {
		return err
	}

	// Get original dimensions
	srcWidth := src.Bounds().Dx()
	srcHeight := src.Bounds().Dy()

	// Calculate new dimensions
	var newWidth, newHeight int
	if srcWidth > srcHeight {
		newWidth = maxLength
		newHeight = (srcHeight * maxLength) / srcWidth
	} else {
		newHeight = maxLength
		newWidth = (srcWidth * maxLength) / srcHeight
	}

	// Resize the image
	dst := imaging.Resize(src, newWidth, newHeight, imaging.Lanczos)

	// Save the resized image
	err = imaging.Save(dst, outputPath)
	if err != nil {
		log.Fatalf("Got error converting image: %+v", err)
		return err
	}

	return nil
}
