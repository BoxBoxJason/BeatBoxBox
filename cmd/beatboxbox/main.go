package main

import (
	db_controller "BeatBoxBox/internal/controller/db"
	"BeatBoxBox/pkg/logger"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./frontend/dist"))
	http.Handle("/", fs)

	db_controller.CreateDB()
	logger.Info("Server Up & Listening at https://localhost:8080")

	err := http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)
	if err != nil {
		logger.Fatal(err)
	}

}
