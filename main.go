package main

import (
	"api-books/database"
	"api-books/server"
)

func main() {
	database.StartDB()
	server := server.NewServer()
	server.Run()
}
