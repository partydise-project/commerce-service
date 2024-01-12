package main

import (
	"commerce-services/database"
	"commerce-services/server"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func initializeEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	initializeEnv()

	dns := os.Getenv("DSN")

	err := database.InicializeDB(dns)
	if err != nil {
		fmt.Println("Error while initializing the database", err)
		return
	}

	r := server.InicializeRouter()

	puerto := 8080
	direccion := fmt.Sprintf(":%d", puerto)

	fmt.Println("Welcome to the commerce-services.", dns)
	log.Fatal(r.Run(direccion))
}
