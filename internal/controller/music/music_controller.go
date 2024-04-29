/*
package music_controller is the controller for the musics.

Contains the logic for the music handling. Handles the connection between the API and the database.
*/

package music_controller

import (
	"BeatBoxBox/pkg/utils"
	"path/filepath"
)

// Create the music directory if it doesn't exist
func init() {
	go utils.CheckDirExists(filepath.Join("data", "musics"))
	go utils.CheckDirExists(filepath.Join("data", "illustrations"))
}
