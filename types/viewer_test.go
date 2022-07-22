package viewer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestCreateFilesList test if a list widget can correctly be
// created, when given a slice of strings.
func TestCreateFilesList(t *testing.T) {
	viewer := Viewer{
		files: []string{"A", "B"},
	}
	list := viewer.CreateFilesList()
	require.Equal(t, len(viewer.files), list.Length(), "List item count is incorrect")
}

// TestCreateImage test if a canvas image can correctly be created,
// when given a filepath.
func TestCreateImage(t *testing.T) {
	testcases := []struct {
		name     string
		file     string
		expError bool
	}{
		{"Existing file - pass",
			"../images/Example.png",
			false,
		},
		{"Not a .png file - fail",
			"../images/Example.xlsx",
			true,
		},
		{"Invalid .png file - fail",
			"../images/ExampleInvalid.png",
			true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			viewer := Viewer{}
			image, err := viewer.CreateImage(tc.file)
			if tc.expError {
				require.Error(t, err, "Error should occur while creating image: %v\n", err)
			} else {
				require.Equal(t, tc.file, image.File, "Image filepath is incorrect")
				require.NoError(t, err, "Error creating image: %v\n", err)
			}
		})
	}
}

// TestUpdateImage tests if the image can
// be updated correctly.
func TestUpdateImage(t *testing.T) {
	testcases := []struct {
		name     string
		file     string
		expEqual bool
	}{
		{"Existing file - pass",
			"../images/Example.png",
			true,
		},
		{"Not a .png file - fail",
			"../images/Example.xlsx",
			false,
		},
		{"Invalid .png file - fail",
			"../images/ExampleInvalid.png",
			false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			viewer := Viewer{}
			_, err := viewer.CreateImage("../images/Example2.png")
			require.NoError(t, err, "Error creating image: %v\n", err)

			image := viewer.UpdateImage(tc.file)
			if tc.expEqual {
				require.Equal(t, tc.file, image.File, "Image filepath is incorrect")
			} else {
				require.NotEqual(t, tc.file, image.File, "Image filepath is incorrect")
			}
		})
	}
}
