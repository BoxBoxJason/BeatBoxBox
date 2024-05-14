package album_controller

import (
	db_model "BeatBoxBox/internal/model"
	album_model "BeatBoxBox/internal/model/album"
	format_utils "BeatBoxBox/pkg/utils/formatutils"
	"encoding/json"
)

// AlbumExists checks if a album exists in the database
func AlbumExists(album_id int) bool {
	db, err := db_model.OpenDB()
	if err != nil {
		return false
	}
	defer db_model.CloseDB(db)
	_, err = album_model.GetAlbum(db, album_id)
	return err == nil
}

// AlbumsExists checks if albums exist in the database
func AlbumsExists(album_ids []int) bool {
	db, err := db_model.OpenDB()
	if err != nil {
		return false
	}
	defer db_model.CloseDB(db)
	albums, err := album_model.GetAlbums(db, album_ids)
	return err != nil && len(albums) == len(album_ids)
}

func AlbumExistsFromFilters(title string, artists_ids []int) bool {
	db, err := db_model.OpenDB()
	if err != nil {
		return false
	}
	defer db_model.CloseDB(db)

	albums, err := album_model.GetAlbumsFromFilters(db, map[string]interface{}{"title": title})
	valid_albums := []db_model.Album{}
	// Remove albums that do not have the same artists_ids
	for _, album := range albums {
		for _, artist := range album.Artists {
			if format_utils.CheckIntInArray(artists_ids, artist.Id) {
				valid_albums = append(valid_albums, album)
			}
		}
	}

	return err == nil && len(valid_albums) > 0
}

// GetAlbum returns a album from the database
// Selects the album with the given album_id
// Returns the album as a JSON object
func GetAlbum(album_id int) ([]byte, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db_model.CloseDB(db)
	album, err := album_model.GetAlbum(db, album_id)
	if err != nil {
		return nil, err
	}
	return json.Marshal(album)
}

// GetAlbums returns a list of albums from the database
// Selects the albums with the given album_ids
// Returns the albums as a JSON array
func GetAlbums(albums_ids []int) ([]byte, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db_model.CloseDB(db)
	albums, err := album_model.GetAlbums(db, albums_ids)
	if err != nil {
		return nil, err
	}
	return json.Marshal(albums)
}

func GetMusicsPathFromAlbum(album_id int) (string, []string, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return "", nil, err
	}
	defer db_model.CloseDB(db)
	album, err := album_model.GetAlbum(db, album_id)
	if err != nil {
		return "", nil, err
	}
	musics_paths := []string{}
	for _, music := range album.Musics {
		musics_paths = append(musics_paths, music.Path)
	}
	return album.Title, musics_paths, nil
}

func GetMusicsPathFromAlbums(albums_ids []int) (map[string][]string, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db_model.CloseDB(db)
	albums, err := album_model.GetAlbums(db, albums_ids)
	if err != nil {
		return nil, err
	}
	musics_paths := map[string][]string{}
	for _, album := range albums {
		musics_paths[album.Title] = []string{}
		for _, music := range album.Musics {
			musics_paths[album.Title] = append(musics_paths[album.Title], music.Path)
		}
	}
	return musics_paths, nil
}

func GetAlbumsFromPartialTitle(partial_title string) ([]byte, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db_model.CloseDB(db)
	albums, err := album_model.GetAlbumsFromPartialTitle(db, partial_title)
	if err != nil {
		return nil, err
	}
	return json.Marshal(albums)
}
