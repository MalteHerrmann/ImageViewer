package viewer

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// TestGetImagesFromFolder tests the GetImagesFromFolder function.
// It does so by checking two reference folders - one
// containing images and one without.
func TestGetImagesFromFolder(t *testing.T) {
	testcases := []struct {
		name      string
		folder    string
		expImages int
		expError  bool
	}{
		{"Existing folder - pass",
			"../images",
			2,
			false,
		},
		{"Folder without images - pass",
			"./",
			0,
			false,
		},
		{"Not a folder - fail",
			"../abcdef",
			0,
			true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			images, err := GetImagesFromFolder(tc.folder)
			if tc.expError {
				require.Error(t, err, "Error should occur while getting images: %v\n", err)
			} else {
				require.NoError(t, err, "Error getting images: %v\n", err)
				require.Equal(t, len(images), tc.expImages, "Number of images is incorrect")
			}
		})
	}
}
