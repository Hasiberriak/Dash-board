package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"strconv"

	_ "github.com/lib/pq"
)

var db *sql.DB





func main() {
	err := godotenv.Load(".env") // โหลดไฟล์ .env
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

	host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT") 
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")

	portInt, err := strconv.Atoi(port)

	if err != nil {
		log.Fatalf("Invalid port number: %v", err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, portInt, user, password, dbname)
	
	sdb ,err := sql.Open("postgres", psqlInfo)

	db = sdb

	
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}
	

	log.Println("Successfully connected to the database!")

	

}

