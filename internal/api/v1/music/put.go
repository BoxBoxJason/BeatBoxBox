/*
Contains the handler for the PUT requests to the /musics endpoint
*/
package music_handler

import (
	album_controller "BeatBoxBox/internal/controller/album"
	music_controller "BeatBoxBox/internal/controller/music"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// putMusicsHandler updates the music with the given ID
// @Summary Update a music by its ID
// @Description Update a music by its ID
// @Tags musics
// @Accept json
// @Produce json
// @Param music_id path int true "Music Id"
// @Param title formData string false "Title"
// @Param description formData string false "Description"
// @Param album_id formData int false "Album Id"
// @Param genres formData []string false "Genres"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Invalid music ID provided, please use a valid integer music ID"
// @Failure 404 {string} string "Music does not exist"
// @Failure 500 {string} string "Internal server error when updating music"
// @Router /api/musics/{music_id} [put]
func putMusicsHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// Get the music ID from the URL
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

	// Parse the url parameters and retrieve only authorized ones
	update_dict := make(map[string]interface{})
	title := r.FormValue("title")
	if title != "" {
		update_dict["Title"] = title
	}
	description := r.FormValue("description")
	if description != "" {
		update_dict["Description"] = description
	}

	album_id_str := r.FormValue("album_id")
	if album_id_str != "" {
		album_id, err := strconv.Atoi(album_id_str)
		if err != nil {
			http.Error(w, "Invalid album ID provided, please use a valid integer album ID", http.StatusBadRequest)
			return
		}
		if !album_controller.AlbumExists(album_id) {
			http.Error(w, "Album does not exist", http.StatusNotFound)
			return
		}
		update_dict["AlbumId"] = album_id
	}
	genres := r.Form["genres"]
	if len(genres) > 0 {
		update_dict["Genres"] = genres
	}

	// Update the music in the database
	err = music_controller.UpdateMusic(music_id, update_dict)
	if err != nil {
		http.Error(w, "Internal server error when updating music: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
