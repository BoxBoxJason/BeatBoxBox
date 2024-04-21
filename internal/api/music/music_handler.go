/*
Package music_handler is the handler for music API

This package is responsible for handling all the API requests related to music.
It creates a new router and registers all the handlers for the music API.
*/

package music_handler

import "github.com/gorilla/mux"

var MusicsRouter = mux.NewRouter()

func init() {

	// POST requests
	MusicsRouter.HandleFunc("/musics", uploadHandler).Methods("POST")

	// GET requests
	MusicsRouter.HandleFunc("/musics", getMusicsHandler).Methods("GET")
	MusicsRouter.HandleFunc("/musics/download", downloadMusicsHandler).Methods("GET")
	MusicsRouter.HandleFunc("/musics/download/{music_id}", downloadMusicHandler).Methods("GET")
	MusicsRouter.HandleFunc("/musics/{music_id}", getMusicHandler).Methods("GET")

	// PUT requests
	MusicsRouter.HandleFunc("/musics/{music_id}", putMusicsHandler).Methods("PUT")

	// DELETE requests
	MusicsRouter.HandleFunc("/musics/{music_id}", deleteMusicHandler).Methods("DELETE")
	MusicsRouter.HandleFunc("/musics", deleteMusicsHandler).Methods("DELETE")

}
