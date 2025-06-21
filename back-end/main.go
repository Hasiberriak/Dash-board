package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"dashboard/handlers"
	"dashboard/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

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
	apiPort := os.Getenv("API_PORT")

	portInt, err := strconv.Atoi(port)

	if err != nil {
		log.Fatalf("Invalid port number: %v", err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, portInt, user, password, dbname)

	sdb, err := sql.Open("postgres", psqlInfo)

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

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Origin,Content-Type,Accept",
	}))

	transactionHandler := handlers.NewTransactionHandler(db)

	routes.SetupRoutes(app, transactionHandler)

	log.Printf("Server starting on port %s\n", apiPort)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", apiPort)))

}
