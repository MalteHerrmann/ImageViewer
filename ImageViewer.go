// ImageViewer is a program, that can be used to preview images from
// a system folder.
package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
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
	image.FillMode = canvas.ImageFillContain

	return image, nil
}

// updateImage takes the image canvas and a filepath. It checks if the
// file exists, has the .png extension and loads it into the canvas if
// that is the case.
func updateImage(image *canvas.Image, filepath string) *canvas.Image {
	if strings.Contains(filepath, ".png") {
		if _, err := os.Stat(filepath); err == nil {
			image.File = filepath
			image.Refresh()
		}
	}

	return image
}

// fillList takes a List type and fills it with files contained in the
// given directory.
func fillList(image *canvas.Image, files []string) *widget.List {
	// Instantiate new list
	list := widget.NewList(func() int {
		return len(files)
	}, func() fyne.CanvasObject {
		return widget.NewLabel("Template")
	}, func(id widget.ListItemID, object fyne.CanvasObject) {
		object.(*widget.Label).SetText(files[id])
	})

	return list
}

// Get folder contents
func getFolderContents(folder string) ([]string, error) {
	// Get list of files in directory
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatalf("Error reading directory: %v\n", err)
	}

	// Create new slice to store filepaths
	var filepaths []string

	// Loop through files and add filepaths to slice
	for _, file := range files {
		filepaths = append(filepaths, file.Name())
	}

	return filepaths, nil
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
	filepath2 := "./Example2.png"

	// Get image container
	imageContainer, err := getImage(filepath)
	if err != nil {
		log.Fatalf("Error creating image: %v\n", err)
	}

	// Get folder contents
	files, err := getFolderContents("./")
	if err != nil {
		log.Fatalf("Error getting folder contents: %v\n", err)
	}

	// Fill list container
	list := fillList(imageContainer, files)

	// Assign new image to container
	updateImage(imageContainer, filepath2)

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
