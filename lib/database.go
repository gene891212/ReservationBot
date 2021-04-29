package lib

import (
	"database/sql"
	"fmt"
	"linebot-server/stru"

	"github.com/line/line-bot-sdk-go/linebot"
)

func AllUserFromDB(db *sql.DB) []stru.UserInfo {
	rows, err := db.Query("SELECT ID, UserID, displayName FROM Users")
	if err != nil {
		panic(err) // proper error handling instead of panic in your app
	}
	defer rows.Close()

	var (
		user    stru.UserInfo
		allUser []stru.UserInfo
	)

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.UserID, &user.DisplayName)
		if err != nil {
			panic(err)
		}
		allUser = append(allUser, user)
	}
	return allUser
}

func GetUser(db *sql.DB, name string) stru.UserInfo {
	// stmt, _ := db.Prepare("SELECT (UserID, DisplayName) FROM Users WHERE Name=?")
	// rows, err := stmt.Query(name)
	rows := db.QueryRow("SELECT UserID, DisplayName FROM Users WHERE DisplayName=?", name)

	user := stru.UserInfo{}

	err := rows.Scan(&user.UserID, &user.DisplayName)
	if err != nil {
		panic(err)
	}

	return user
}

func InsertUserToDB(db *sql.DB, profile *linebot.UserProfileResponse) {
	stmt, _ := db.Prepare("INSERT INTO Users (UserID, DisplayName, PictureURL) VALUES (?, ?, ?)")
	res, err := stmt.Exec(
		profile.UserID,
		profile.DisplayName,
		profile.PictureURL,
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

func InsertReservationToDB(db *sql.DB) {
	db.Exec("")
}
