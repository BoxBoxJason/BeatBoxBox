package album_controller

import (
	db_tables "BeatBoxBox/internal/model"
	album_model "BeatBoxBox/internal/model/album"
	artist_model "BeatBoxBox/internal/model/artist"
	"BeatBoxBox/pkg/db_model"
	file_utils "BeatBoxBox/pkg/utils/fileutils"
	httputils "BeatBoxBox/pkg/utils/httputils"
	"mime/multipart"
)

func PostAlbum(title string, artists_ids []int, description string, release_date string, illustration_file *multipart.FileHeader) ([]byte, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return []byte{}, err
	}
	defer db_model.CloseDB(db)
	artists, err := artist_model.GetArtists(db, artists_ids)
	if err != nil {
		return []byte{}, err
	} else if artists == nil || len(artists) != len(artists_ids) {
		return []byte{}, httputils.NewNotFoundError("some artists do not exist")
	}
	artists_ptr := make([]*db_tables.Artist, len(artists))
	for i, artist := range artists {
		artists_ptr[i] = &artist
	}
	if album_model.AlbumAlreadyExists(db, title, artists_ids) {
		return []byte{}, httputils.NewConflictError("album already exists")
	}
	illustration_file_name, err := file_utils.UploadIllustrationToServer(illustration_file, "albums")
	if err != nil {
		return []byte{}, httputils.NewInternalServerError("could not upload illustration")
	}

	album, err := album_model.CreateAlbum(db, title, description, illustration_file_name, release_date, artists_ptr)
	if err != nil {
		return []byte{}, err
	}
	return ConvertAlbumToJSON(&album)
}
