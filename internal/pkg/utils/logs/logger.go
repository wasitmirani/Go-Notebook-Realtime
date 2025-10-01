package logs

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

// ANSI colors for console logs
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorCyan   = "\033[36m"
)

// Global logger
var logger *log.Logger

// SetupLogger initializes daily rotating logs and sets global logger
func SetupLogger(logDir string) {
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		os.Mkdir(logDir, 0755)
	}

	appName := "GoNotebookRealtime"
	writer, err := rotatelogs.New(
		fmt.Sprintf("%s/%s-%%Y-%%m-%%d.log", logDir, appName),
		rotatelogs.WithRotationTime(24*time.Hour), // rotate daily
		rotatelogs.WithMaxAge(7*24*time.Hour),     // keep 7 days
	)
	if err != nil {
		log.Fatalf("failed to create rotatelogs: %s", err)
	}

	// MultiWriter → log to file + terminal
	multiWriter := io.MultiWriter(os.Stdout, writer)
	logger = log.New(multiWriter, "", log.Ldate|log.Ltime|log.Lshortfile)
}

// Helper functions for log levels
func Info(msg string) {
	logger.Println(ColorGreen + " INFO: " + ColorReset + msg)
}

func Warning(msg string) {
	logger.Println(ColorYellow + " WARNING: " + ColorReset + msg)
}

func Error(msg string) {
	logger.Println(ColorRed + " ERROR: " + ColorReset + msg)
}

func Danger(msg string) {
	logger.Println(ColorRed + " DANGER: " + ColorReset + msg)
}

func Debug(msg string) {
	logger.Println(ColorBlue + " DEBUG: " + ColorReset + msg)
}

func Success(msg string) {
	logger.Println(ColorCyan + " SUCCESS: " + ColorReset + msg)
}

// General purpose custom log
func Log(level, msg string) {
	logger.Println(fmt.Sprintf(" [%s] %s", level, msg))
}
