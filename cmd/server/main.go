package main

import (
	api_init "BeatBoxBox/internal/api"
	db_model "BeatBoxBox/pkg/db_model"
	"BeatBoxBox/pkg/logger"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Serve the frontend
	fs := http.FileServer(http.Dir("./frontend/dist"))
	http.Handle("/", fs)

	// Check the database connection
	err := db_model.CheckDB()
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
	logger.Info("Server Up & Listening at https://localhost:3000")
	err = http.ListenAndServeTLS(":3000", "secret/cert.pem", "secret/key.pem", nil)
	if err != nil {
		logger.Critical(err)
	}
}
