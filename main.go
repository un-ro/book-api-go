package main

import "book-api-go/routers"

func main() {
	routers.StartServer().Run(":8080")
}
