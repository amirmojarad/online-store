// Package db provides crud functions and config attributes for database.
package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type databaseConf struct {
	Username string
	Password string
	Host     string
	Dbname   string
}

// At first, it clean env of os then loading env variables from "controllers/db/db.env"
func init() {
	os.Clearenv()
	if err := godotenv.Load("controllers/db/db.env"); err != nil {
		log.Fatalln("Error loading db.env file", err)
	}
}

// get loaded env variables and create a new databaseConf struct
func GetDatabaseConfig() *databaseConf {
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")
	dbName := os.Getenv("DB_NAME")

	return &databaseConf{
		Username: username,
		Password: password,
		Host:     host,
		Dbname:   dbName,
	}
}
