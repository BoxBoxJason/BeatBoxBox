package main

import (
	api_init "BeatBoxBox/internal/api"
	db_model "BeatBoxBox/pkg/db_model"
	"BeatBoxBox/pkg/logger"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Check the database connection
	err := db_model.CheckDB()
	if err != nil {
		logger.Critical(err)
		return
	}
	logger.Info("Database connection is alive !")

	// Setup the main router
	main_router := mux.NewRouter()

	// Serve the frontend
	fs := http.FileServer(http.Dir("./frontend/dist"))
	main_router.PathPrefix("/").Handler(fs)

	// Setup the API routes
	api_router := main_router.PathPrefix("/api").Subrouter()
	api_init.SetupAPIRouter(api_router)
	logger.Info("API routes are set up")

	// Start the server
	logger.Info("Server Up & Listening at https://localhost:3000")
	err = http.ListenAndServeTLS(":3000", "secret/cert.pem", "secret/key.pem", main_router)
	if err != nil {
		logger.Critical("Unable to start the server using TLS: ", err)
		logger.Debug("Starting the server without TLS")
		err = http.ListenAndServe(":3000", main_router)
		if err != nil {
			logger.Fatal("Unable to start the server", err)
		} else {
			logger.Info("Server started successfully on http://localhost:3000")
		}
	} else {
		logger.Info("Server started successfully on https://localhost:3000")
	}
}
