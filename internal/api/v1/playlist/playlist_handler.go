package playlist_handler_v1

import "github.com/gorilla/mux"

func SetupPlaylistAPIRoutes(playlist_api_router *mux.Router) {
	// POST requests
	playlist_api_router.HandleFunc("/", postPlaylistHandler).Methods("POST")

	// GET requests
	playlist_api_router.HandleFunc("/", getPlaylistsHandler).Methods("GET")
	playlist_api_router.HandleFunc("/download", downloadPlaylistsHandler).Methods("GET")
	playlist_api_router.HandleFunc("/{playlist_id:[0-9]+}/download", downloadPlaylistHandler).Methods("GET")
	playlist_api_router.HandleFunc("/{playlist_id:[0-9]+}", getPlaylistHandler).Methods("GET")

	// PATCH requests
	playlist_api_router.HandleFunc("/{playlist_id:[0-9]+}/musics/{action:(add|remove)}", updatePlaylistMusicsHandler).Methods("PATCH")
	playlist_api_router.HandleFunc("/{playlist_id:[0-9]+}/owners/{action:(add|remove)}", updatePlaylistOwnersHandler).Methods("PATCH")
	playlist_api_router.HandleFunc("/{playlist_id:[0-9]+}", patchPlaylistHandler).Methods("PATCH")

	// DELETE requests
	playlist_api_router.HandleFunc("/{playlist_id:[0-9]+}", deletePlaylistHandler).Methods("DELETE")
	playlist_api_router.HandleFunc("/", deletePlaylistsHandler).Methods("DELETE")
}
