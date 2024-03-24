package main

import (
	"mygram/database"
	"mygram/routers"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	database.StartDB()
	r := routers.StartApp()
	r.Run(":" + os.Getenv("PORT"))
}
