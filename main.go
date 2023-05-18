package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/rfauzi44/up/db"
	"github.com/rfauzi44/up/routers"
)

func main() {
	db.Init()
	e := routers.New()
	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}
