package playlist_controller

import "errors"

// DeletePlaylist deletes an playlist from the database
// Selects the playlist with the given playlist_id
func DeletePlaylist(playlist_id int) error {
	if !PlaylistExists(playlist_id) {
		return errors.New("playlist does not exist")
	}
	return DeletePlaylist(playlist_id)
}

// DeletePlaylists deletes a list of playlists from the database
// Selects the playlists with the given playlist_ids
func DeletePlaylists(playlist_ids []int) error {
	if !PlaylistsExist(playlist_ids) {
		return errors.New("at least one playlist does not exist")
	}
	return DeletePlaylists(playlist_ids)
}
