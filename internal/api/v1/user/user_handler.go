package user_handler

import "github.com/gorilla/mux"

func SetupUsersRoutes(user_api_router *mux.Router) {
	// POST requests
	user_api_router.HandleFunc("/", registerHandler).Methods("POST")
	user_api_router.HandleFunc("/auth", loginHandler).Methods("POST")
	user_api_router.HandleFunc("/logout", logoutHandler).Methods("POST")

	// GET requests
	user_api_router.HandleFunc("/", getUsersHandler).Methods("GET")
	user_api_router.HandleFunc("/{user_id:[0-9]+}", getUserHandler).Methods("GET")

	// PUT requests
	user_api_router.HandleFunc("/{user_id:[0-9]+}", putUserHandler).Methods("PUT")

	// DELETE requests
	user_api_router.HandleFunc("/{user_id:[0-9]+}", deleteUserHandler).Methods("DELETE")
	user_api_router.HandleFunc("/", deleteUsersHandler).Methods("DELETE")
}
