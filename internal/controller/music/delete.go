package music_controller

import (
	"fmt"
)

// DeleteMusic deletes a music by its id
func DeleteMusic(music_id int) error {
	music_exists := MusicExists(music_id)
	if !music_exists {
		return fmt.Errorf("music with id %d does not exist", music_id)
	}
	return DeleteMusic(music_id)
}

// DeleteMusics deletes musics by their ids
func DeleteMusics(music_ids []int) error {
	musics_exists := MusicsExists(music_ids)
	if !musics_exists {
		return fmt.Errorf("musics with ids %v do not exist", music_ids)
	}
	return DeleteMusics(music_ids)
}
