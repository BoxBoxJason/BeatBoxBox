/*
package music_controller is the controller for the musics.

Contains the logic for the music handling. Handles the connection between the API and the database.
*/

package music_controller

import (
	"os"
	"path/filepath"
)

// Create the music directory if it doesn't exist
func init() {
	musics_dir := filepath.Join("data", "musics")
	if _, err := os.Stat(musics_dir); os.IsNotExist(err) {
		os.MkdirAll(musics_dir, 0755)
	}
}
