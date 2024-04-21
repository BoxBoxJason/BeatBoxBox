/*
Contains the handler for the GET requests to the /api/musics endpoint
*/
package music_handler

import (
	music_controller "BeatBoxBox/internal/controller/music"
	"archive/zip"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func getMusicHandler(w http.ResponseWriter, r *http.Request) {
	// Get the music ID from the URL
	music_id_str := mux.Vars(r)["id"]
	music_id, err := strconv.Atoi(music_id_str)
	if err != nil {
		http.Error(w, "Invalid music ID provided, please use a valid integer music ID", http.StatusBadRequest)
		return
	}

	// Get the music from the database
	music, err := music_controller.GetMusic(music_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the music as a JSON response
	w.Header().Set("Content-Type", "application/json")
	w.Write(music)
	w.WriteHeader(http.StatusOK)
}

// GET musics handler, checks if the request is a GET request and then
// Returns all the musics in the database as a JSON response
// Can filter musics ids with the music_ids request parameter
func getMusicsHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve requested music IDs from the URL
	query_params := r.URL.Query()
	music_ids := []int{}

	if query_params.Get("music_ids") != "" {
		music_ids_str := strings.Split(query_params.Get("music_ids"), ",")
		for _, music_id_str := range music_ids_str {
			music_id, err := strconv.Atoi(music_id_str)
			if err != nil {
				http.Error(w, "Invalid music ID provided, please use a valid integer music ID", http.StatusBadRequest)
				return
			}
			music_ids = append(music_ids, music_id)
		}
	}

	// Get the musics from the database
	musics_json, err := music_controller.GetMusics(music_ids)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the musics as a JSON response
	w.Header().Set("Content-Type", "application/json")
	w.Write(musics_json)
	w.WriteHeader(http.StatusOK)
}

// Download music handler
// Returns the music file corresponding to the music ID
func downloadMusicHandler(w http.ResponseWriter, r *http.Request) {
	// Get the music ID from the URL & retrieve the corresponding music path
	music_id_str := mux.Vars(r)["music_id"]
	music_id, err := strconv.Atoi(music_id_str)
	if err != nil {
		http.Error(w, "Invalid music ID provided, please use a valid integer music ID", http.StatusBadRequest)
		return
	}

	music_path, err := music_controller.GetMusicPathFromId(music_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Open the music file
	music_file, err := os.Open(music_path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer music_file.Close()

	// Serve the file
	http.ServeFile(w, r, music_path)
	w.WriteHeader(http.StatusOK)
}

// Download musics handler
// Returns the music(s) file(s) corresponding to the music ID(s)
func downloadMusicsHandler(w http.ResponseWriter, r *http.Request) {
	// Get the music ID from the URL & retrieve the corresponding music path
	music_ids_requested := r.URL.Query().Get("music_ids")
	if music_ids_requested == "" {
		http.Error(w, "No music ID provided, please use music_ids request parameter", http.StatusBadRequest)
		return
	}

	music_ids_str := strings.Split(music_ids_requested, ",")
	music_ids := []int{}
	for _, music_id_str := range music_ids_str {
		music_id, err := strconv.Atoi(music_id_str)
		if err != nil {
			http.Error(w, "Invalid music ID provided, please use a valid music ID", http.StatusBadRequest)
			return
		}
		music_ids = append(music_ids, music_id)
	}

	musics_paths, err := music_controller.GetMusicsPathFromIds(music_ids)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = zipMusics(w, musics_paths)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Zip musics into a single zip file and write it to the response writer
// Takes a list of music paths and writes them to the zip file
func zipMusics(w http.ResponseWriter, musics_paths []string) error {
	zipped_musics := zip.NewWriter(w)
	defer zipped_musics.Close()

	for _, music_path := range musics_paths {
		music_file, err := os.Open(music_path)
		if err != nil {
			return err
		}
		defer music_file.Close()
		music_info, err := music_file.Stat()
		if err != nil {
			return err
		}
		music_header, err := zip.FileInfoHeader(music_info)
		if err != nil {
			return err
		}
		music_header.Name = music_path
		music_header.Method = zip.Deflate
		writer, err := zipped_musics.CreateHeader(music_header)
		if err != nil {
			return err
		}
		if _, err := io.Copy(writer, music_file); err != nil {
			return err
		}
	}

	if err := zipped_musics.Close(); err != nil {
		return err
	}

	return nil
}