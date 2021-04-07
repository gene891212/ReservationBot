package lib

import (
	"database/sql"
	"linebot-server/stru"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv"
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

func DataFromDB() []stru.User {
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		log.Fatal(err) // proper error handling instead of panic in your app
	}
	defer rows.Close()

	var (
		id           int
		userId, name string
		allUser      []stru.User
	)

	for rows.Next() {
		err = rows.Scan(&id, &userId, &name)
		if err != nil {
			log.Fatal(err)
		}
		allUser = append(
			allUser,
			stru.User{
				ID:     id,
				UserId: userId,
				Name:   name,
			},
		)
	}
	return allUser
}
