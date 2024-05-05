package playlist_controller

import (
	file_utils "BeatBoxBox/pkg/utils/fileutils"
	"path/filepath"
)

// Create the music directory if it doesn't exist
func init() {
	go file_utils.CheckDirExists(filepath.Join("data", "illustrations", "playlists"))
}
