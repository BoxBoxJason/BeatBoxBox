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

// Logger instances
var (
	debugLogger    *log.Logger
	infoLogger     *log.Logger
	errorLogger    *log.Logger
	criticalLogger *log.Logger
	fatalLogger    *log.Logger
)

func init() {
	// Check if the directory for the log file exists
	if _, err := os.Stat("./log"); os.IsNotExist(err) {
		os.Mkdir("./log", os.ModePerm)
	}

	file, err := os.OpenFile("./log/server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file:", err)
	}

	multi := io.MultiWriter(file, os.Stdout)

	// Initialize loggers
	debugLogger = log.New(multi, "DEBUG: ", log.Ldate|log.Ltime)
	infoLogger = log.New(multi, "INFO: ", log.Ldate|log.Ltime)
	errorLogger = log.New(multi, "ERROR: ", log.Ldate|log.Ltime)
	criticalLogger = log.New(multi, "CRITICAL: ", log.Ldate|log.Ltime)
	fatalLogger = log.New(multi, "FATAL: ", log.Ldate|log.Ltime)
}

// Debug logs a debug message.
func Debug(v ...interface{}) {
	debugLogger.Println(v...)
}

// Info logs an info message.
func Info(v ...interface{}) {
	infoLogger.Println(v...)
}

// Error logs an error message.
func Error(v ...interface{}) {
	errorLogger.Println(v...)
}

// Critical logs a critical message but does not exit the application.
func Critical(v ...interface{}) {
	criticalLogger.Println(v...)
}

// Fatal logs a fatal message and exits the application.
func Fatal(v ...interface{}) {
	fatalLogger.Fatalln(v...)
}
