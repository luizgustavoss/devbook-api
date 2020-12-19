package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	// connectionStringPattern database connection string pattern
	connectionStringPattern = "%s:%s@/%s?charset=utf8&parseTime=True&loc=Local"
	// ConnectionString database connection string
	ConnectionString string
	// Port API port number
	Port int
	// SecretKey key used to sign token
	SecretKey []byte
)

// Load inits environment variables
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

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}