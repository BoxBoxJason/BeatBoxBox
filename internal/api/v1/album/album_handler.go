package album_handler

import "github.com/gorilla/mux"

func SetupAlbumAPIRoutes(album_api_router *mux.Router) {
	// POST requests
	album_api_router.HandleFunc("/", createAlbumHandler).Methods("POST")

	// GET requests
	album_api_router.HandleFunc("/{album_id:[0-9]+}", getAlbumHandler).Methods("GET")
	album_api_router.HandleFunc("/", getAlbumsHandler).Methods("GET")
	album_api_router.HandleFunc("/{album_partial_title:[a-zA-Z0-9_\\-\\s'\"&éèêëàçù]+}", getAlbumsByPartialTitleHandler).Methods("GET")
	album_api_router.HandleFunc("/download", downloadAlbumsHandler).Methods("GET")
	album_api_router.HandleFunc("/download/{album_id:[0-9]+}", downloadAlbumHandler).Methods("GET")

	// PUT requests
	album_api_router.HandleFunc("/{album_id:[0-9]+}/add", addMusicsToAlbumHandler).Methods("PUT")
	album_api_router.HandleFunc("/{album_id:[0-9]+/add/{music_id:[0-9]+}}", addMusicToAlbumHandler).Methods("PUT")
	album_api_router.HandleFunc("/{album_id:[0-9]+/remove}", removeMusicsFromAlbumHandler).Methods("PUT")
	album_api_router.HandleFunc("/{album_id:[0-9]+/remove/{music_id:[0-9]+}}", removeMusicFromAlbumHandler).Methods("PUT")
	album_api_router.HandleFunc("/{album_id:[0-9]+}", putAlbumHandler).Methods("PUT")

	// DELETE requests
	album_api_router.HandleFunc("/{album_id:[0-9]+}", deleteAlbumHandler).Methods("DELETE")
	album_api_router.HandleFunc("/", deleteAlbumsHandler).Methods("DELETE")
}
