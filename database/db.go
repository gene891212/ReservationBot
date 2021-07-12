package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

func InitDb() (*sql.DB, error) {
	var config = mysql.Config{
		User:   os.Getenv("user"),
		Passwd: os.Getenv("passwd"),
		Net:    os.Getenv("net"),
		Addr:   os.Getenv("addr"),
		DBName: os.Getenv("dbname"),
	}
	fmt.Println(os.Getenv("addr"))
	db, err := sql.Open("mysql", config.FormatDSN())
	// db, err := sql.Open("mysql", "root:password@tcp(192.168.0.10:3306)/linebot")
	return db, err
}
