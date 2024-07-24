package api_init

import (
	album_handler_v1 "BeatBoxBox/internal/api/v1/album"
	artist_handler_v1 "BeatBoxBox/internal/api/v1/artist"
	music_handler_v1 "BeatBoxBox/internal/api/v1/music"
	playlist_handler_v1 "BeatBoxBox/internal/api/v1/playlist"
	user_handler_v1 "BeatBoxBox/internal/api/v1/user"

	"github.com/gorilla/mux"
)

func SetupAPIRouter(api_router *mux.Router) {
	setupAPIV1Router(api_router.PathPrefix("/v1").Subrouter())
}

func setupAPIV1Router(api_router *mux.Router) {
	music_api_router := api_router.PathPrefix("/musics").Subrouter()
	music_handler_v1.SetupMusicsRoutes(music_api_router)

	user_api_router := api_router.PathPrefix("/users").Subrouter()
	user_handler_v1.SetupUsersRoutes(user_api_router)

	playlist_api_router := api_router.PathPrefix("/playlists").Subrouter()
	playlist_handler_v1.SetupPlaylistAPIRoutes(playlist_api_router)

	albums_api_router := api_router.PathPrefix("/albums").Subrouter()
	album_handler_v1.SetupAlbumAPIRoutes(albums_api_router)

	artists_api_router := api_router.PathPrefix("/artists").Subrouter()
	artist_handler_v1.SetupArtistsAPIRoutes(artists_api_router)
}
