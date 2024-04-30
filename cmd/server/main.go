package main

import (
	api_init "BeatBoxBox/internal/api"
	"BeatBoxBox/internal/model/dbinit"
	"BeatBoxBox/pkg/logger"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Serve the frontend
	fs := http.FileServer(http.Dir("./frontend/dist"))
	http.Handle("/", fs)

	// Check the database connection
	err := dbinit.CheckDB()
	if err != nil {
		logger.Critical(err)
		return
	}
	logger.Info("Database connection is alive !")

	// Setup the routes
	main_router := mux.NewRouter()
	api_router := main_router.PathPrefix("/api").Subrouter()
	api_init.SetupAPIRouter(api_router)
	logger.Info("API routes are set up")

	// Start the server
	logger.Info("Server Up & Listening at https://localhost:8080")
	err = http.ListenAndServeTLS(":8080", "secret/cert.pem", "secret/key.pem", nil)
	if err != nil {
		logger.Critical(err)
	}
}
