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
		ID           int
		userID, name string
		allUser      []stru.User
	)

	for rows.Next() {
		err = rows.Scan(&ID, &userID, &name)
		if err != nil {
			panic(err)
		}
		allUser = append(
			allUser,
			stru.User{
				ID:     ID,
				UserId: userID,
				Name:   name,
			},
		)
	}
	return allUser
}
