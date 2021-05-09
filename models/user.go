package models

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
)

type User struct {
	ID            int    `json:"id"`
	UserID        string `json:"userId"`
	DisplayName   string `json:"displayName"`
	PictureUrl    string `json:"pictureUrl"`
	StatusMessage string `json:"statusMessage"`
}

func GetUser(db *sql.DB, name string) (User, error) {
	// stmt, _ := db.Prepare("SELECT (UserID, DisplayName) FROM Users WHERE Name=?")
	// rows, err := stmt.Query(name)
	rows := db.QueryRow("SELECT ID, UserID, DisplayName FROM Users WHERE DisplayName=?", name)

	user := User{}
	err := rows.Scan(&user.ID, &user.UserID, &user.DisplayName)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func GetUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT ID, UserID, DisplayName FROM Users")
	if err != nil {
		return []User{}, err
	}
	defer rows.Close()

	user := User{}
	allUsers := []User{}

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.UserID, &user.DisplayName)
		if err != nil {
			return []User{}, err
		}
		allUsers = append(allUsers, user)
	}
	return allUsers, nil
}

func CreateUser(db *sql.DB, profile *linebot.UserProfileResponse) error {
	stmt, _ := db.Prepare("INSERT INTO Users (UserID, DisplayName, PictureURL) VALUES (?, ?, ?)")
	_, err := stmt.Exec(
		profile.UserID,
		profile.DisplayName,
		profile.PictureURL,
	)
	if err != nil {
		return err
	}
	return nil
}

func GetUserByAccessToken(accessToken string) (User, error) {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", "https://api.line.me/v2/profile", nil)

	req.Header.Add("Authorization", "Bearer "+accessToken)
	resp, err := client.Do(req)
	if err != nil {
		return User{}, nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return User{}, nil
	}

	user := User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return User{}, nil
	}

	return user, nil
}
