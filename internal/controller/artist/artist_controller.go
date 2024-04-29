package artist_controller

import (
	"BeatBoxBox/pkg/utils"
	"path/filepath"
)

// Create the albums illustrations directory if it doesn't exist
func init() {
	go utils.CheckDirExists(filepath.Join("data", "illustrations", "artists"))
}
