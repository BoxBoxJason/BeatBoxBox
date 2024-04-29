package album_controller

import "fmt"

// DeleteAlbum deletes a album by its id
func DeleteAlbum(album_id int) error {
	album_exists := AlbumExists(album_id)
	if !album_exists {
		return fmt.Errorf("album with id %d does not exist", album_id)
	}
	return DeleteAlbum(album_id)
}

// DeleteAlbums deletes albums by their ids
func DeleteAlbums(album_ids []int) error {
	albums_exists := AlbumsExists(album_ids)
	if !albums_exists {
		return fmt.Errorf("albums with ids %v do not exist", album_ids)
	}
	return DeleteAlbums(album_ids)
}
