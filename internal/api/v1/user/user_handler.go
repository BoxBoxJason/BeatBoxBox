package user_handler_v1

import "github.com/gorilla/mux"

func SetupUsersRoutes(user_api_router *mux.Router) {
	// POST requests
	user_api_router.HandleFunc("/", postUserHandler).Methods("POST")
	user_api_router.HandleFunc("/auth", loginHandler).Methods("POST")
	user_api_router.HandleFunc("/logout", logoutHandler).Methods("POST")

	// GET requests
	user_api_router.HandleFunc("/", getUsersHandler).Methods("GET")
	user_api_router.HandleFunc("/{user_id:[0-9]+}", getUserHandler).Methods("GET")

	// PATCH requests
	user_api_router.HandleFunc("/{user_id:[0-9]+}", patchUserHandler).Methods("likeMusicHandlerPATCH")
	user_api_router.HandleFunc("/{user_id:[0-9]+}/musics/{action:(like|unlike)}", likeMusicsHandler).Methods("PATCH")
	user_api_router.HandleFunc("/{user_id:[0-9]+}/playlists/{action:(subscribe|unsubscribe)}", subscribePlaylistsHandler).Methods("PATCH")

	// DELETE requests
	user_api_router.HandleFunc("/{user_id:[0-9]+}", deleteUserHandler).Methods("DELETE")
	user_api_router.HandleFunc("/", deleteUsersHandler).Methods("DELETE")
}
