package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/Dhanraj-Patil/event-trigger/internal/models"
)

var logDir = "event-logs"

func InitLogger() {
	err := os.MkdirAll("event-logs", os.ModePerm)
	if err != nil {
		log.Fatal("ERROR: Could not create logs directory:", err)
	}
}

func LogEvent(data models.EventLog) {
	logFilePath := generateLogFilePath(data.Trigger.ID)
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("ERROR: Could not create log file:", err)
	}
	defer file.Close()

	// Set output to the new log file
	log.SetOutput(file)

	// Write a sample log
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(jsonData))

	// Start cleanup scheduler in a separate goroutine
	go deleteOldLogs(48 * time.Hour)

	// Simulate long-running process
	// select {}
}

func generateLogFilePath(eventId string) string {
	return filepath.Join(logDir, fmt.Sprintf("%s.log", eventId))
}

// deleteOldLogs deletes log files older than the specified duration
func deleteOldLogs(retention time.Duration) {
	ticker := time.NewTicker(1 * time.Hour) // Run cleanup every hour
	defer ticker.Stop()

	for range ticker.C {
		files, err := os.ReadDir(logDir)
		if err != nil {
			log.Println("ERROR: Could not read logs directory:", err)
			continue
		}

		now := time.Now()

		for _, file := range files {
			filePath := filepath.Join(logDir, file.Name())

			// Get file info
			info, err := os.Stat(filePath)
			if err != nil {
				log.Println("ERROR: Could not get file info:", err)
				continue
			}

			// Check if file is older than retention period
			if now.Sub(info.ModTime()) > retention {
				err := os.Remove(filePath)
				if err != nil {
					log.Println("ERROR: Could not delete old log file:", err)
				} else {
					log.Println("INFO: Deleted old log file:", filePath)
				}
			}
		}
	}
}
