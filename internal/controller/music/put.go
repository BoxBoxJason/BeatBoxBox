package music_controller

import (
	album_controller "BeatBoxBox/internal/controller/album"
	db_tables "BeatBoxBox/internal/model"
	album_model "BeatBoxBox/internal/model/album"
	artist_model "BeatBoxBox/internal/model/artist"
	music_model "BeatBoxBox/internal/model/music"
	db_model "BeatBoxBox/pkg/db_model"
	httputils "BeatBoxBox/pkg/utils/httputils"
)

func UpdateMusic(music_id int, music_map map[string]interface{}) ([]byte, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return []byte{}, err
	}
	defer db_model.CloseDB(db)
	music, err := music_model.GetMusic(db, music_id)
	if err != nil {
		return []byte{}, err
	}
	err = music_model.UpdateMusic(db, &music, music_map)
	if err != nil {
		return []byte{}, err
	}
	return ConvertMusicToJSON(&music)
}

func AddArtistsToMusic(music_id int, artists_ids []int) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)
	music, err := music_model.GetMusic(db, music_id)
	if err != nil {
		return err
	}
	artists, err := artist_model.GetArtists(db, artists_ids)
	if err != nil {
		return err
	} else if artists == nil || len(artists) != len(artists_ids) {
		return httputils.NewNotFoundError("some artists were not found")
	}
	artists_ptr := make([]*db_tables.Artist, len(artists))
	for i, artist := range artists {
		artists_ptr[i] = &artist
	}
	return music_model.AddArtistsToMusic(db, &music, artists_ptr)
}

func RemoveArtistsFromMusic(music_id int, artists_ids []int) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)
	music, err := music_model.GetMusic(db.Preload("Artists"), music_id)
	if err != nil {
		return err
	}
	artists_ids_map := make(map[int]bool, len(artists_ids))
	for _, artist_id := range artists_ids {
		artists_ids_map[artist_id] = true
	}
	artists_ptr := make([]*db_tables.Artist, len(artists_ids))
	for i, artist := range music.Artists {
		if _, ok := artists_ids_map[artist.Id]; ok {
			artists_ptr[i] = &artist
		}
	}
	if len(artists_ptr) != len(artists_ids) {
		return httputils.NewNotFoundError("some artists were not found")
	}
	return music_model.RemoveArtistsFromMusic(db, &music, artists_ptr)
}

func RemoveAlbumFromMusics(musics_ids []int, album_id int) ([]byte, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return []byte{}, err
	}
	defer db_model.CloseDB(db)
	musics, err := music_model.GetMusics(db, musics_ids)
	if err != nil {
		return []byte{}, err
	} else if len(musics) != len(musics_ids) {
		return []byte{}, httputils.NewNotFoundError("some musics were not found")
	}
	album, err := album_model.GetAlbum(db, album_id)
	if err != nil {
		return []byte{}, err
	}
	musics_ptr := make([]*db_tables.Music, len(musics))
	for i, music := range musics {
		musics_ptr[i] = &music
	}
	err = music_model.RemoveAlbumFromMusics(db, musics_ptr, &album)
	if err != nil {
		return []byte{}, err
	}
	album_json, err := album_controller.ConvertAlbumToJSON(&album)
	if err != nil {
		return []byte{}, err
	}
	return album_json, nil
}

func AddMusicsToAlbum(musics_ids []int, album_id int) ([]byte, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return []byte{}, err
	}
	defer db_model.CloseDB(db)
	musics, err := music_model.GetMusics(db, musics_ids)
	if err != nil {
		return []byte{}, err
	}
	album, err := album_model.GetAlbum(db, album_id)
	if err != nil {
		return []byte{}, err
	}
	musics_ptr := make([]*db_tables.Music, len(musics))
	for i, music := range musics {
		musics_ptr[i] = &music
	}
	err = music_model.AddAlbumToMusics(db, musics_ptr, &album)
	if err != nil {
		return []byte{}, err
	}
	album_json, err := album_controller.ConvertAlbumToJSON(&album)
	if err != nil {
		return []byte{}, err
	}
	return album_json, nil
}
