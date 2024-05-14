package artist_handler

import "github.com/gorilla/mux"

func SetupArtistsRoutes(artist_api_router *mux.Router) {
	// POST
	artist_api_router.HandleFunc("/", postArtistHandler).Methods("POST")

	// GET
	artist_api_router.HandleFunc("/", getArtistsHandler).Methods("GET")
	artist_api_router.HandleFunc("/{artist_id:[0-9]+}", getArtistHandler).Methods("GET")

	// PUT
	artist_api_router.HandleFunc("/{artist_id:[0-9]+}", putArtistHandler).Methods("PUT")

	// DELETE
	artist_api_router.HandleFunc("/{artist_id:[0-9]+}", deleteArtistHandler).Methods("DELETE")
	artist_api_router.HandleFunc("/", deleteArtistsHandler).Methods("DELETE")
}
