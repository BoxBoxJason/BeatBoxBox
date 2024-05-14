package api_init

import (
	album_handler "BeatBoxBox/internal/api/album"
	music_handler "BeatBoxBox/internal/api/music"
	playlist_handler "BeatBoxBox/internal/api/playlist"
	user_handler "BeatBoxBox/internal/api/user"

	"github.com/gorilla/mux"
)

func SetupAPIRouter(api_router *mux.Router) {
	music_api_router := api_router.PathPrefix("/musics").Subrouter()
	music_handler.SetupMusicsRoutes(music_api_router)

	user_api_router := api_router.PathPrefix("/users").Subrouter()
	user_handler.SetupUsersRoutes(user_api_router)

	playlist_api_router := api_router.PathPrefix("/playlists").Subrouter()
	playlist_handler.SetupPlaylistAPIRoutes(playlist_api_router)

	albums_api_router := api_router.PathPrefix("/albums").Subrouter()
	album_handler.SetupAlbumAPIRoutes(albums_api_router)
}
