package album_controller

import (
	db_tables "BeatBoxBox/internal/model"
	album_model "BeatBoxBox/internal/model/album"
	"BeatBoxBox/pkg/db_model"
	file_utils "BeatBoxBox/pkg/utils/fileutils"
	httputils "BeatBoxBox/pkg/utils/httputils"
	"net/http"
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
	album, err := album_model.GetAlbum(db.Preload("Musics").Preload("Artists"), album_id)
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
	albums, err := album_model.GetAlbums(db.Preload("Musics").Preload("Artists"), albums_ids)
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
	albums := album_model.GetAlbumsFromPartialTitle(db.Preload("Musics").Preload("Artists"), map[string]interface{}{}, partial_title)
	albums_ptr := make([]*db_tables.Album, len(albums))
	for i, album := range albums {
		albums_ptr[i] = &album
	}
	return ConvertAlbumsToJSON(albums_ptr)
}

func GetAlbumsJSONFromFilters(titles []string, partial_titles []string, genres []string, artists_names []string, musics_names []string, artists_ids []int, musics_ids []int) ([]byte, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db_model.CloseDB(db)
	if len(titles) > 0 && len(partial_titles) > 0 {
		return nil, httputils.NewBadRequestError("Can't use title with partial_title")
	} else if len(artists_names)*len(artists_ids) > 0 {
		return nil, httputils.NewBadRequestError("Can't use artist with artist_id")
	} else if len(musics_names)*len(musics_ids) > 0 {
		return nil, httputils.NewBadRequestError("Can't use music with music_id")
	}
	albums := album_model.GetAlbumsFromFilters(db.Preload("Musics").Preload("Artists"), titles, partial_titles, genres, artists_names, musics_names, artists_ids, musics_ids)
	albums_ptr := make([]*db_tables.Album, len(albums))
	for i, album := range albums {
		albums_ptr[i] = &album
	}
	return ConvertAlbumsToJSON(albums_ptr)
}

func ServeAlbumFiles(w http.ResponseWriter, album_id int) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	album, err := album_model.GetAlbum(db.Preload("Musics"), album_id)
	if err != nil {
		return err
	}
	paths := make([]string, len(album.Musics))
	for i, music := range album.Musics {
		paths[i] = file_utils.GetAbsoluteMusicPath(music.Path)
	}
	httputils.ServeZip(w, paths, album.Title+".zip")
	return nil
}

func ServeAlbumsFiles(w http.ResponseWriter, albums_ids []int) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	albums, err := album_model.GetAlbums(db.Preload("Musics"), albums_ids)
	if err != nil {
		return err
	}
	paths := make(map[string][]string, len(albums))
	for _, album := range albums {
		paths[album.Title] = make([]string, len(album.Musics))
		for j, music := range album.Musics {
			paths[album.Title][j] = file_utils.GetAbsoluteMusicPath(music.Path)
		}
	}
	httputils.ServeSubdirsZip(w, paths, "albums.zip")
	return nil
}
