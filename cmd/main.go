package main

import (
	"log"

	"github.com/joho/godotenv"

	"go_web_server/db"
	"go_web_server/server"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("env file error!")
	}

	if err := db.InitDB(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer db.DB.Close()

	server.StartHttpServer()
}
