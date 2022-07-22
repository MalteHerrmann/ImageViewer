// ImageViewer is a program, that can be used to preview images from
// a system folder.
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	viewer "github.com/MalteHerrmann/ImageViewer/types"
	"log"
)

func main() {
	// Initialize a new fyne app
	a := app.New()

	// Open new window
	w := a.NewWindow("Image Viewer")

	// Set window size
	w.Resize(fyne.NewSize(800, 600))

	// Define filename
	defaultImage := "./images/Example.png"

	// Define viewer struct
	imageViewer := viewer.Viewer{}

	// Assign image files from folder
	filesInFolder, err := viewer.GetImagesFromFolder("./images")
	if err != nil {
		log.Fatalf("Error reading directory: %v\n", err)
	}
	imageViewer.SetFiles(filesInFolder)

	// Get image container
	imageContainer, err := imageViewer.CreateImage(defaultImage)
	if err != nil {
		log.Fatalf("Error creating image: %v\n", err)
	}

	// Fill list container
	list := imageViewer.CreateFilesList()

	// Define horizontal split
	hSplit := container.NewHSplit(
		list,
		imageContainer,
	)
	hSplit.SetOffset(0.2)

	//Assign content to window
	w.SetContent(hSplit)

	// Show window
	w.ShowAndRun()
}
