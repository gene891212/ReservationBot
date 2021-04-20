package lib

import (
	"database/sql"
	"linebot-server/stru"
)

func AllUserFromDB(db *sql.DB) []stru.User {
	rows, err := db.Query("SELECT * FROM Users")
	if err != nil {
		panic(err) // proper error handling instead of panic in your app
	}
	defer rows.Close()

	var (
		id        int
		uID, name string
		allUser   []stru.User
	)

	for rows.Next() {
		err = rows.Scan(&id, &uID, &name)
		if err != nil {
			panic(err)
		}
		allUser = append(
			allUser,
			stru.User{
				ID:   id,
				UID:  uID,
				Name: name,
			},
		)
	}
	return allUser
}

func InsertUserToDB(db *sql.DB) {
	db.Exec("INSERT INTO Users (UID, Name) VALUES (?, ?)")
}

func InsertReserationToDB(db *sql.DB) {
	db.Exec("")
}
