package album_handler_v1

import (
	album_controller "BeatBoxBox/internal/controller/album"
	music_controller "BeatBoxBox/internal/controller/music"
	custom_errors "BeatBoxBox/pkg/errors"
	file_utils "BeatBoxBox/pkg/utils/fileutils"
	format_utils "BeatBoxBox/pkg/utils/formatutils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// addMusicsToAlbumHandler adds musics to an album
// @ Summary: Adds musics to an album
// @ Description: Adds musics to an album
// @ Tags: albums
// @ Accept: json
// @ Produces: json
// @ Param: album_id path int true "Album ID"
// @ Param: musics_ids query []int true "Music IDs"
// @ Failure 400 {string} string "Invalid album ID provided, please use a valid integer album ID"
// @ Failure 404 {string} string "Album does not exist"
// @ Failure 400 {string} string "musics_ids must be comma separated integers"
// @ Failure 404 {string} string "At least one music does not exist"
// @ Success 200 {string} string "OK"
// @ Router /api/albums/{album_id}/add [put]
func addMusicsToAlbumHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// Get album ID from URL
	album_id_str := mux.Vars(r)["album_id"]
	album_id, err := strconv.Atoi(album_id_str)
	if err != nil {
		http.Error(w, "Invalid album ID provided, please use a valid integer album ID", http.StatusBadRequest)
		return
	}
	if !album_controller.AlbumExists(album_id) {
		http.Error(w, "Album does not exist", http.StatusNotFound)
		return
	}

	// Get music IDs from the request body
	musics_ids_str := r.URL.Query().Get("musics_ids")
	musics_ids, err := format_utils.ConvertStringToIntArray(musics_ids_str, ",")
	if err != nil {
		http.Error(w, "musics_ids must be comma separated integers", http.StatusBadRequest)
		return
	}
	if len(musics_ids) > 0 {
		if !music_controller.MusicsExists(musics_ids) {
			http.Error(w, "At least one music does not exist", http.StatusNotFound)
			return
		}
		err = album_controller.AddMusicsToAlbum(album_id, musics_ids)
		if err != nil {
			custom_errors.SendErrorToClient(err, w, "")
		}
	}
	w.WriteHeader(http.StatusOK)
}

// addMusicToAlbumHandler adds a music to an album
// @ Summary: Adds a music to an album
// @ Description: Adds a music to an album
// @ Tags: albums
// @ Accept: json
// @ Produces: json
// @ Param: album_id path int true "Album ID"
// @ Param: music_id path int true "Music ID"
// @ Failure 400 {string} string "Invalid album ID provided, please use a valid integer album ID"
// @ Failure 404 {string} string "Album does not exist"
// @ Failure 400 {string} string "Invalid music ID provided, please use a valid integer music ID"
// @ Failure 404 {string} string "Music does not exist"
// @ Success 200 {string} string "OK"
// @ Router /api/albums/{album_id}/add/{music_id} [put]
func addMusicToAlbumHandler(w http.ResponseWriter, r *http.Request) {
	// Get album ID from URL
	album_id_str := mux.Vars(r)["album_id"]
	album_id, err := strconv.Atoi(album_id_str)
	if err != nil {
		http.Error(w, "Invalid album ID provided, please use a valid integer album ID", http.StatusBadRequest)
		return
	}
	if !album_controller.AlbumExists(album_id) {
		http.Error(w, "Album does not exist", http.StatusNotFound)
		return
	}

	// Get music ID from URL
	music_id_str := mux.Vars(r)["music_id"]
	music_id, err := strconv.Atoi(music_id_str)
	if err != nil {
		http.Error(w, "Invalid music ID provided, please use a valid integer music ID", http.StatusBadRequest)
		return
	}
	if !music_controller.MusicExists(music_id) {
		http.Error(w, "Music does not exist", http.StatusNotFound)
		return
	}
	err = album_controller.AddMusicsToAlbum(album_id, []int{music_id})
	if err != nil {
		http.Error(w, "Error when adding music to album: "+err.Error(), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
}

// removeMusicsFromAlbumHandler removes musics from an album
// @ Summary: Removes musics from an album
// @ Description: Removes musics from an album
// @ Tags: albums
// @ Accept: json
// @ Produces: json
// @ Param: album_id path int true "Album ID"
// @ Param: musics_ids query []int true "Music IDs"
// @ Failure 400 {string} string "Invalid album ID provided, please use a valid integer album ID"
// @ Failure 404 {string} string "Album does not exist"
// @ Failure 400 {string} string "musics_ids must be integers"
// @ Failure 404 {string} string "At least one music does not exist"
// @ Success 200 {string} string "OK"
// @ Router /api/albums/{album_id}/remove [put]
func removeMusicsFromAlbumHandler(w http.ResponseWriter, r *http.Request) {
	// Get album ID from URL
	album_id_str := mux.Vars(r)["album_id"]
	album_id, err := strconv.Atoi(album_id_str)
	if err != nil {
		http.Error(w, "Invalid album ID provided, please use a valid integer album ID", http.StatusBadRequest)
		return
	}
	if !album_controller.AlbumExists(album_id) {
		http.Error(w, "Album does not exist", http.StatusNotFound)
		return
	}

	// Get music IDs from the request body
	musics_ids_str := r.URL.Query().Get("musics_ids")
	musics_ids, err := format_utils.ConvertStringToIntArray(musics_ids_str, ",")
	if err != nil {
		http.Error(w, "musics_ids must be integers", http.StatusBadRequest)
		return
	}
	if len(musics_ids) > 0 {
		if !music_controller.MusicsExists(musics_ids) {
			http.Error(w, "At least one music does not exist", http.StatusNotFound)
			return
		}
		err = album_controller.RemoveMusicsFromAlbum(album_id, musics_ids)
		if err != nil {
			http.Error(w, "Error when removing musics from album: "+err.Error(), http.StatusBadRequest)
		}
	}
	w.WriteHeader(http.StatusOK)
}

// removeMusicFromAlbumHandler removes a music from an album
// @ Summary: Removes a music from an album
// @ Description: Removes a music from an album
// @ Tags: albums
// @ Accept: json
// @ Produces: json
// @ Param: album_id path int true "Album ID"
// @ Param: music_id path int true "Music ID"
// @ Failure 400 {string} string "Invalid album ID provided, please use a valid integer album ID"
// @ Failure 404 {string} string "Album does not exist"
// @ Failure 400 {string} string "Invalid music ID provided, please use a valid integer music ID"
// @ Failure 404 {string} string "Music does not exist"
// @ Success 200 {string} string "OK"
// @ Router /api/albums/{album_id}/remove/{music_id} [put]
func removeMusicFromAlbumHandler(w http.ResponseWriter, r *http.Request) {
	// Get album ID from URL
	album_id_str := mux.Vars(r)["album_id"]
	album_id, err := strconv.Atoi(album_id_str)
	if err != nil {
		http.Error(w, "Invalid album ID provided, please use a valid integer album ID", http.StatusBadRequest)
		return
	}
	if !album_controller.AlbumExists(album_id) {
		http.Error(w, "Album does not exist", http.StatusBadRequest)
		return
	}

	// Get music ID from URL
	music_id_str := mux.Vars(r)["music_id"]
	music_id, err := strconv.Atoi(music_id_str)
	if err != nil {
		http.Error(w, "Invalid music ID provided, please use a valid integer music ID", http.StatusBadRequest)
		return
	}
	if !music_controller.MusicExists(music_id) {
		http.Error(w, "Music does not exist", http.StatusBadRequest)
		return
	}
	err = album_controller.RemoveMusicsFromAlbum(album_id, []int{music_id})
	if err != nil {
		http.Error(w, "Error when removing music from album: "+err.Error(), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
}

// putAlbumHandler updates an album
// @ Summary: Updates an album
// @ Description: Updates an album
// @ Tags: albums
// @ Accept: json
// @ Produces: json
// @ Param: album_id path int true "Album ID"
// @ Param: title formData string false "Album title"
// @ Param: description formData string false "Album description"
// @ Param: illustration formData file false "Album illustration"
// @ Failure 400 {string} string "Invalid album ID provided, please use a valid integer album ID"
// @ Failure 404 {string} string "Album does not exist"
// @ Success 200 {string} string "OK"
// @ Router /api/albums/{album_id} [put]
func putAlbumHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(file_utils.MAX_REQUEST_SIZE)
	if err != nil {
		http.Error(w, "Request too large", http.StatusRequestEntityTooLarge)
		return
	}

	album_id_str := mux.Vars(r)["album_id"]
	album_id, err := strconv.Atoi(album_id_str)
	if err != nil {
		http.Error(w, "Invalid album ID provided, please use a valid integer album ID", http.StatusBadRequest)
		return
	}
	if !album_controller.AlbumExists(album_id) {
		http.Error(w, "Album does not exist", http.StatusBadRequest)
		return
	}

	update_dict := make(map[string]interface{})

	title := r.FormValue("title")
	if title != "" {
		update_dict["title"] = title
	}

	description := r.FormValue("description")
	if description != "" {
		update_dict["description"] = description
	}

	illustration_file, illustration_file_header, err := r.FormFile("illustration")
	if err == nil {
		illustration_file_name, err := file_utils.UploadIllustrationToServer(illustration_file_header, illustration_file, "albums")
		if err == nil && illustration_file_name != file_utils.DEFAULT_ILLUSTRATION_FILE {
			update_dict["illustration"] = illustration_file_name
		}
	}

	err = album_controller.UpdateAlbum(album_id, update_dict)
	if err != nil {
		custom_errors.SendErrorToClient(err, w, "")
		return
	}
	w.WriteHeader(http.StatusOK)
}
