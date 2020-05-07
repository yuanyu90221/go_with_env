package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error Loading .env file")
	// }
	log.Println(os.Getenv("DB_PASSWD"))
}
