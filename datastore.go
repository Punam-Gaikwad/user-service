package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func CreateConnection() (*gorm.DB, error) {
	// Get database details from environment variables
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	DBName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	//host := "localhost"
	//user := "postgres"
	//DBName := "mydb"
	//password := ""

	uri := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		user, password, host, DBName,
	)
	log.Println("postgres is listening on: ", uri)

	return gorm.Open(
		"postgres",
		uri,
	)
}
