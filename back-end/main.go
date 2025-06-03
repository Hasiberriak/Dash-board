package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber"
	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	host     = "localhost"  
	port     = 5432         
	user     = "postgres"     
	password = "0986803508Aboss66" 
	dbname   = "postgres" 
  )

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	
	sdb ,err := sql.Open("postgres", psqlInfo)

	db = sdb

	
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}
	
	db.Close()

	log.Println("Successfully connected to the database!")

	app := fiber.New()


}