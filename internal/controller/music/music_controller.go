/*
package music_controller is the controller for the musics.

Contains the logic for the music handling. Handles the connection between the API and the database.
*/

package music_controller

import (
	file_utils "BeatBoxBox/pkg/utils/fileutils"
	"path/filepath"
)

// Create the music directory & musics illustrations directory if it doesn't exist
func init() {
	go file_utils.CheckDirExists(filepath.Join("data", "musics"))
	go file_utils.CheckDirExists(filepath.Join("data", "illustrations", "musics"))
}
