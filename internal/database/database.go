package database

import (
	"log"

	"os"

	"github.com/joho/godotenv"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDB() {
	godotenv.Load()
	err := mgm.SetDefaultConfig(nil, os.Getenv("DATABASE_NAME"), options.Client().ApplyURI(os.Getenv("MONOGDB_URI")))
	println(os.Getenv("MONGODB_URI"))
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
}
