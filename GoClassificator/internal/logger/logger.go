package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

// Logger struct holds loggers for different log levels
type Logger struct {
	InfoLog  *log.Logger
	WarnLog  *log.Logger
	ErrorLog *log.Logger
}

var AppLogger *Logger

// InitLogger initializes the logging service
func InitLogger(logDir string) {

	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Fatal("Could not create log directory:", err)
	}

	logFilePath := filepath.Join(logDir, "application.log")
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Could not open log file:", err)
	}

	// Initialize the logger
	AppLogger = &Logger{
		InfoLog:  log.New(logFile, "INFO: ", log.Ldate|log.Ltime),
		WarnLog:  log.New(logFile, "WARNING: ", log.Ldate|log.Ltime),
		ErrorLog: log.New(logFile, "ERROR: ", log.Ldate|log.Ltime),
	}

	log.SetOutput(logFile)

	log.Println("Logging initialized successfully")
}

func getCallerInfo() string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return "unknown:0"
	}
	return fmt.Sprintf("%s:%d", filepath.Base(file), line)
}

func Info(v ...any) {
	AppLogger.InfoLog.Println(append([]any{getCallerInfo()}, v...)...)
}

func Warn(v ...any) {
	AppLogger.WarnLog.Println(append([]any{getCallerInfo()}, v...)...)
}

// Error logs an error message and returns an error object
func Error(v ...any) error {
	err := fmt.Sprintln(v...)
	AppLogger.ErrorLog.Println(append([]any{getCallerInfo()}, err)...)
	return fmt.Errorf("%s", err)
}

func Infof(format string, v ...any) {
	AppLogger.InfoLog.Printf("%s "+format, append([]any{getCallerInfo()}, v...)...)
}

func Warnf(format string, v ...any) {
	AppLogger.WarnLog.Printf("%s "+format, append([]any{getCallerInfo()}, v...)...)
}

// Errorf logs an error message and returns an error object
func Errorf(format string, v ...any) error {
	err := fmt.Sprintf(format, v...)
	AppLogger.ErrorLog.Printf("%s "+format, append([]any{getCallerInfo()}, v...)...)
	return fmt.Errorf("%s", err)
}

func Fatalf(format string, v ...any) {
	AppLogger.ErrorLog.Fatalf("%s "+format, append([]any{getCallerInfo()}, v...)...)
}

// GetLogger returns the global logger instance
func GetLogger() *Logger {
	return AppLogger
}
