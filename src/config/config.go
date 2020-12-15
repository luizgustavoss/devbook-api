package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	// database connection string pattern
	connectionStringPattern = "%s:%s@/%s?charset=utf8&parseTime=True&loc=Local"
	// database connection string
	ConnectionString = ""
	// API port
	Port = 0
)

// init environment vars
func Load() {
	var error error

	if error = godotenv.Load(); error != nil {
		log.Fatal(error)
	}

	Port, error = strconv.Atoi(os.Getenv("API_PORT"))
	if error != nil {
		Port = 9000 // assumes default port number
	}

	ConnectionString = fmt.Sprintf(connectionStringPattern, os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
}