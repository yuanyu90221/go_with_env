package main

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error Loading .env file")
	// }
	log.Println(os.Getenv("DB_PASSWD"))
}
