package main

import (
	"mygram/database"
	"mygram/routers"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	database.StartDB()
	r := routers.StartApp()
	r.Run(":8080")
}
