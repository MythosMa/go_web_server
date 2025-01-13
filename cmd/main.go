package main

import (
	"log"

	"github.com/joho/godotenv"

	"go_web_server/server"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("env file error!")
	}

	server.StartHttpServer()
}
