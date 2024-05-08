package playlist_handler

import "github.com/gorilla/mux"

func SetupPlaylistAPIRoutes(playlist_api_router *mux.Router) { // TODO
	// POST requests
	playlist_api_router.HandleFunc("/", createPlaylistHandler).Methods("POST")

	// GET requests
	playlist_api_router.HandleFunc("/", getPlaylistsHandler).Methods("GET")
	playlist_api_router.HandleFunc("/download", downloadPlaylistsHandler).Methods("GET")
	playlist_api_router.HandleFunc("/download/{playlist_id:[0-9]+}", downloadPlaylistHandler).Methods("GET")
	playlist_api_router.HandleFunc("/{playlist_id:[0-9]+}", getPlaylistHandler).Methods("GET")

	// PUT requests
	playlist_api_router.HandleFunc("/{playlist_id:[0-9]+}/add", addMusicsToPlaylistHandler).Methods("PUT")
	playlist_api_router.HandleFunc("/{playlist_id:[0-9]+/add/{music_id:[0-9]+}}", addMusicToPlaylistHandler).Methods("PUT")
	playlist_api_router.HandleFunc("/{playlist_id:[0-9]+/remove}", removeMusicsFromPlaylistHandler).Methods("PUT")
	playlist_api_router.HandleFunc("/{playlist_id:[0-9]+/remove/{music_id:[0-9]+}}", removeMusicFromPlaylistHandler).Methods("PUT")
	playlist_api_router.HandleFunc("/{playlist_id:[0-9]+}", putPlaylistHandler).Methods("PUT")

	// DELETE requests
	playlist_api_router.HandleFunc("/{playlist_id:[0-9]+}", deletePlaylistHandler).Methods("DELETE")
	playlist_api_router.HandleFunc("/", deletePlaylistsHandler).Methods("DELETE")
}
