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
		log.Println("ERROR: Could not get the working dir: ", err)
	}
}

func LogById(id string, data string) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Println("ERROR: Could not create log file:", err)
	}
	logFilePath := generateLogFilePath(id)
	logDir := filepath.Join(cwd, "..", logFilePath) // Go to parent dir and find event-logs
	file, err := os.OpenFile(logDir, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("ERROR: Could not create log file:", err)
	}
	defer file.Close()

	// Set output to the new log file
	log.SetOutput(file)

	log.Println(data)
	go deleteOldLogs(48 * time.Hour)

}

func LogEvent(data models.EventLog) {
	logFilePath := generateLogFilePath(data.Trigger.ID)
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("ERROR: Could not create log file:", err)
	}
	defer file.Close()

	log.SetOutput(file)

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(jsonData))
}

func generateLogFilePath(eventId string) string {
	return filepath.Join(logDir, fmt.Sprintf("%s.log", eventId))
}

func deleteOldLogs(retention time.Duration) {
	ticker := time.NewTicker(1 * time.Hour)
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

			info, err := os.Stat(filePath)
			if err != nil {
				log.Println("ERROR: Could not get file info:", err)
				continue
			}

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
