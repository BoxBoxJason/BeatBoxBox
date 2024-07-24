/*
Package music_handler_v1 is the handler for music API

This package is responsible for handling all the API requests related to music.
It creates a new router and registers all the handlers for the music API.
*/

package music_handler_v1

import "github.com/gorilla/mux"

func SetupMusicsRoutes(music_api_router *mux.Router) {
	// POST requests
	music_api_router.HandleFunc("/", uploadHandler).Methods("POST")

	// GET requests
	music_api_router.HandleFunc("/", getMusicsHandler).Methods("GET")
	music_api_router.HandleFunc("/download", downloadMusicsHandler).Methods("GET")
	music_api_router.HandleFunc("/download/{music_id:[0-9]+}", downloadMusicHandler).Methods("GET")
	music_api_router.HandleFunc("/{music_id:[0-9]+}", getMusicHandler).Methods("GET")

	// PUT requests
	music_api_router.HandleFunc("/{music_id:[0-9]+}", putMusicsHandler).Methods("PUT")

	// DELETE requests
	music_api_router.HandleFunc("/{music_id:[0-9]+}", deleteMusicHandler).Methods("DELETE")
	music_api_router.HandleFunc("/", deleteMusicsHandler).Methods("DELETE")
}
