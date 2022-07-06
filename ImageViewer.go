// ImageViewer is a program, that can be used to preview images from
// a system folder.
package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

// Get image takes a filepath to an image and creates a new image
// canvas. If the file does not exist, it returns nil and an error.
func getImage(filepath string) (*canvas.Image, error) {
	// Check if file exists
	_, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}

	// Get new image
	image := canvas.NewImageFromFile(filepath)

	return image, nil
}

func main() {
	// Initialize a new fyne app
	a := app.New()

	// Open new window
	w := a.NewWindow("Image Viewer")

	// Set window size
	w.Resize(fyne.NewSize(800, 600))

	// Define filename
	filepath := "./Example.png"

	// Get image container
	imageContainer, err := getImage(filepath)
	if err != nil {
		log.Fatalf("Error creating image: %v\n", err)
	}

	//Assign content to window
	w.SetContent(imageContainer)

	// Show window
	w.ShowAndRun()
}
