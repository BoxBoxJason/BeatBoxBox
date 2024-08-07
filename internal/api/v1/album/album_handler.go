package album_handler_v1

import "github.com/gorilla/mux"

func SetupAlbumAPIRoutes(album_api_router *mux.Router) {
	// POST requests
	album_api_router.HandleFunc("/", postAlbumHandler).Methods("POST")

	// DELETE requests
	album_api_router.HandleFunc("/{album_id:[0-9]+}", deleteAlbumHandler).Methods("DELETE")
	album_api_router.HandleFunc("/", deleteAlbumsHandler).Methods("DELETE")

	// PATCH requests
	album_api_router.HandleFunc("/{album_id:[0-9]+}", patchAlbumHandler).Methods("PATCH")
	album_api_router.HandleFunc("/{album_id:[0-9]+}/artists/{action}", patchAlbumArtistsHandler).Methods("PATCH")
	album_api_router.HandleFunc("/{album_id:[0-9]+}/musics/{action}", patchAlbumMusicsHandler).Methods("PATCH")

	// GET requests
	album_api_router.HandleFunc("/", getAlbumsHandler).Methods("GET")
	album_api_router.HandleFunc("/{album_id:[0-9]+}", getAlbumHandler).Methods("GET")
}
