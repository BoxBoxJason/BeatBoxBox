/*
package logger is the package that contains the logger for the application.

The logger is used to log the application's activity to the console and to a file.
*/

package logger

import (
	"io"
	"log"
	"os"
)

var (
	Debug *log.Logger
	Info  *log.Logger
	Error *log.Logger
	Fatal *log.Logger
)

func init() {
	file, err := os.OpenFile("./log/server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file:", err)
	}

	multi := io.MultiWriter(file, os.Stdout)

	Debug = log.New(multi, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(multi, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(multi, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	Fatal = log.New(multi, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)
}
