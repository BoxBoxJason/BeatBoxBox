package artist_handler_v1

import "github.com/gorilla/mux"

func SetupArtistsAPIRoutes(artist_api_router *mux.Router) {
	// POST
	artist_api_router.HandleFunc("/", postArtistHandler).Methods("POST")

	// GET
	artist_api_router.HandleFunc("/", getArtistsHandler).Methods("GET")
	artist_api_router.HandleFunc("/{artist_id:[0-9]+}", getArtistHandler).Methods("GET")

	// PATCH
	artist_api_router.HandleFunc("/{artist_id:[0-9]+}", putArtistHandler).Methods("PATCH")

	// DELETE
	artist_api_router.HandleFunc("/{artist_id:[0-9]+}", deleteArtistHandler).Methods("DELETE")
	artist_api_router.HandleFunc("/", deleteArtistsHandler).Methods("DELETE")
}
