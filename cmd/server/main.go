package main

import (
	"BeatBoxBox/internal/model/dbinit"
	"BeatBoxBox/pkg/logger"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./frontend/dist"))
	http.Handle("/", fs)

	err := dbinit.CheckDB()
	if err != nil {
		logger.Critical(err)
	} else {
		logger.Info("Database connection is alive !")
	}
	// Start the server
	logger.Info("Server Up & Listening at https://localhost:8080")
	err = http.ListenAndServeTLS(":8080", "secret/cert.pem", "secret/key.pem", nil)
	if err != nil {
		logger.Critical(err)
	}
}
