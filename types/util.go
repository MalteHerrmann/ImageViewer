package viewer

import (
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
)

// GetImagesFromFolder returns a list of image files
// from the given directory.
func GetImagesFromFolder(folder string) ([]string, error) {
	supportedFiletypes := []string{
		".png",
		".jpg",
		".jpeg",
	}
	supportedFiletypesPattern := regexp.MustCompile(`(?i)(` + strings.Join(supportedFiletypes, "|") + `)$`)

	// Get list of files in directory
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		return nil, err
	}

	// Create new slice to store file paths
	var filepaths []string

	// Loop through files and add file paths to slice
	for _, file := range files {
		filename := filepath.Join(folder, file.Name())
		if supportedFiletypesPattern.FindString(filename) != "" {
			filepaths = append(filepaths, filename)
		}
	}

	return filepaths, nil
}
