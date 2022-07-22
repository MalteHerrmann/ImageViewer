package viewer

import (
	"fmt"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

// Viewer defines the type for the UI.
type Viewer struct {
	files []string
	image *canvas.Image
	List  *widget.List
}

// SetFiles takes a slice of file paths as an input argument
// and sets the viewer files slice.
func (v *Viewer) SetFiles(files []string) {
	v.files = files
}

// CreateFilesList takes a slice of file paths as an input argument.
// It creates a new surrounding container and fills it with buttons
// for every file.
func (v *Viewer) CreateFilesList() *widget.List {
	// Instantiate new list
	list := widget.NewList(func() int {
		return len(v.files)
	}, func() fyne.CanvasObject {
		return widget.NewLabel("Template")
	}, func(id widget.ListItemID, object fyne.CanvasObject) {
		object.(*widget.Label).SetText(v.files[id])
	})

	list.OnSelected = func(id widget.ListItemID) {
		fmt.Println("Updating: ", v.files[id])
		v.UpdateImage(v.files[id])
	}

	return list
}

// CreateImage takes a filepath to an image and creates a new image
// canvas. If the file does not exist, it returns nil and an error.
func (v *Viewer) CreateImage(filepath string) (*canvas.Image, error) {
	// Check if file exists
	_, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}

	// Get new image
	v.image = canvas.NewImageFromFile(filepath)
	v.image.FillMode = canvas.ImageFillContain

	return v.image, nil
}

// UpdateImage takes a filepath as an input argument and
// checks if the file exists. If it does, it's loaded to
// the image canvas.
func (v *Viewer) UpdateImage(filepath string) *canvas.Image {
	if strings.Contains(filepath, ".png") {
		if _, err := os.Stat(filepath); err == nil {
			v.image.File = filepath
			v.image.Refresh()
		} else {
			fmt.Println("Not a valid filepath: ", filepath)
		}
	}

	return v.image
}
