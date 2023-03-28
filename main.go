package main

import (
	"book-api-go/database"
	"book-api-go/routers"
)

func main() {
	database.StartDB()
	routers.StartServer().Run(":8080")
}
