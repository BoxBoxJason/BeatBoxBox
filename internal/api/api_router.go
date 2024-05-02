package api_init

import (
	music_handler "BeatBoxBox/internal/api/music"
	playlist_handler "BeatBoxBox/internal/api/playlist"

	"github.com/gorilla/mux"
)

func SetupAPIRouter(api_router *mux.Router) {
	music_api_router := api_router.PathPrefix("/musics").Subrouter()
	music_handler.SetupMusicsRoutes(music_api_router)

	playlist_api_router := api_router.PathPrefix("/playlists").Subrouter()
	playlist_handler.SetupPlaylistAPIRoutes(playlist_api_router)
}
