package main

import (
	"kredit_plus/connection"
	"kredit_plus/router"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	e := echo.New()

	db, err := connection.GetConnection()
	if err != nil {
		panic(err)
	}

	router.Router(router.Setups{Echo: e, DB: db})

	port := ":" + os.Getenv("PORT")
	e.Logger.Fatal(e.Start(port))
}
