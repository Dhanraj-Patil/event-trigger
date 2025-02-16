package config

import (
	"log"
	"os"
	"time"
)

func init() {
	os.Setenv("TZ", "Asia/Kolkata")
	ist, err := time.LoadLocation(os.Getenv("TZ"))
	if err != nil {
		log.Fatal("Failed to load timezone:", err)
	}

	time.Local = ist
}
