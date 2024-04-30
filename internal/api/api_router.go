package api_init

import (
	music_handler "BeatBoxBox/internal/api/music"

	"github.com/gorilla/mux"
)

func SetupAPIRouter(api_router *mux.Router) {
	music_api_router := api_router.PathPrefix("/musics").Subrouter()
	music_handler.SetupMusicsRoutes(music_api_router)
}
