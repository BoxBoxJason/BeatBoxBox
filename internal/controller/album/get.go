package album_controller

import (
	db_tables "BeatBoxBox/internal/model"
	album_model "BeatBoxBox/internal/model/album"
	"BeatBoxBox/pkg/db_model"
)

// AlbumExists checks if an album exists in the database
func AlbumExists(album_id int) bool {
	db, err := db_model.OpenDB()
	if err != nil {
		return false
	}
	defer db_model.CloseDB(db)
	_, err = album_model.GetAlbum(db, album_id)
	return err == nil
}

// AlbumsExist checks if albums exist in the database
func AlbumsExist(album_ids []int) bool {
	db, err := db_model.OpenDB()
	if err != nil {
		return false
	}
	defer db_model.CloseDB(db)
	albums, err := album_model.GetAlbums(db, album_ids)
	return err == nil && len(albums) == len(album_ids)
}

// GetAlbum returns an album from the database in JSON format
func GetAlbumJSON(album_id int) ([]byte, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db_model.CloseDB(db)
	album, err := album_model.GetAlbum(db, album_id)
	if err != nil {
		return nil, err
	}
	return ConvertAlbumToJSON(&album)
}

// GetAlbums returns a list of albums from the database in JSON format
func GetAlbumsJSON(albums_ids []int) ([]byte, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db_model.CloseDB(db)
	albums, err := album_model.GetAlbums(db, albums_ids)
	if err != nil {
		return nil, err
	}
	albums_ptr := make([]*db_tables.Album, len(albums))
	for i, album := range albums {
		albums_ptr[i] = &album
	}
	return ConvertAlbumsToJSON(albums_ptr)
}

func GetMusicsPathFromAlbum(album_id int) (string, []string, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return "", nil, err
	}
	defer db_model.CloseDB(db)
	album, err := album_model.GetAlbum(db.Preload("Musics"), album_id)
	if err != nil {
		return "", nil, err
	}
	musics_paths := []string{}
	for _, music := range album.Musics {
		musics_paths = append(musics_paths, music.Path)
	}
	return album.Title, musics_paths, nil
}

// GetAlbumsFromFilters returns a list of albums from the database in JSON format
func GetMusicsPathFromAlbums(albums_ids []int) (map[string][]string, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db_model.CloseDB(db)
	albums, err := album_model.GetAlbums(db.Preload("Musics"), albums_ids)
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

// GetAlbumsFromPartialTitle returns a list of albums from the database in JSON format
func GetAlbumsJSONFromPartialTitle(partial_title string) ([]byte, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db_model.CloseDB(db)
	albums := album_model.GetAlbumsFromPartialTitle(db, map[string]interface{}{}, partial_title)
	albums_ptr := make([]*db_tables.Album, len(albums))
	for i, album := range albums {
		albums_ptr[i] = &album
	}
	return ConvertAlbumsToJSON(albums_ptr)
}
