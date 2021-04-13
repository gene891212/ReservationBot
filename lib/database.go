package lib

import (
	"database/sql"
	"linebot-server/stru"
	"log"
)

func DataFromDB(db *sql.DB) []stru.User {
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
