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



type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}


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

	
	err = createProductTable(&Product{
		Name:  "Sample Product",
		Price: 100,
	})

	if err != nil {
		log.Fatalf("Error inserting product: %v", err)
	}


}

func createProductTable(product *Product) error {
	
	_, err := db.Exec("INSERT INTO data(name, pirce) VALUES ($1, $2);",
		product.Name, product.Price)

	return err
}