package lib

import (
	"database/sql"
	"fmt"
	"linebot-server/stru"

	"github.com/line/line-bot-sdk-go/linebot"
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
				ID:     id,
				UserID: uID,
				Name:   name,
			},
		)
	}
	return allUser
}

func GetUser(db *sql.DB, name string) []stru.UserInfo {
	// stmt, _ := db.Prepare("SELECT (UserID, DisplayName) FROM Users WHERE Name=?")
	// rows, err := stmt.Query(name)
	rows, err := db.Query("SELECT UserID, DisplayName FROM Users WHERE DisplayName=?", name)
	if err != nil {
		panic(err)
	}

	users := []stru.UserInfo{}

	var (
		userID, displayName string
	)
	for rows.Next() {
		err = rows.Scan(&userID, &displayName)
		if err != nil {
			panic(err)
		}
		users = append(
			users,
			stru.UserInfo{
				UserID:      userID,
				DisplayName: displayName,
			},
		)
	}
	return users
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

func InsertReserationToDB(db *sql.DB) {
	db.Exec("")
}
