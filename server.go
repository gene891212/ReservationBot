package main

import (
	"linebot-server/lib"
	"linebot-server/router"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

var (
	config = mysql.Config{
		User:   os.Getenv("user"),
		Passwd: os.Getenv("passwd"),
		Net:    os.Getenv("tcp"),
		Addr:   os.Getenv("addr"),
		DBName: os.Getenv("dbname"),
	}
)

func main() {
	lib.SetupRichMenu()

	router := router.SetupRouter()
	err := router.Run(":7000")
	if err != nil {
		log.Fatal(err)
	}
}
