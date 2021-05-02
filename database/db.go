package database

import (
	"database/sql"
	"os"

	"github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

func InitDb() (*sql.DB, error) {
	var config = mysql.Config{
		User:   os.Getenv("user"),
		Passwd: os.Getenv("passwd"),
		Net:    os.Getenv("tcp"),
		Addr:   os.Getenv("addr"),
		DBName: os.Getenv("dbname"),
	}

	db, err := sql.Open("mysql", config.FormatDSN())
	return db, err
}
