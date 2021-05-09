package models

import (
	"database/sql"
	"time"
)

type Message struct {
	Content string    `json:"content"`
	Sender  User      `json:"sender"`
	Reciver User      `json:"reciver"`
	Time    time.Time `json:"time"`
}

// func GetMessage(db *sql.DB, user User) []Message {

// }

func CreateMessage(db *sql.DB, msg Message) error {
	stmt, err := db.Prepare(`INSERT INTO Messages (Content, Sender, Reciver, Time) VALUES (?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		msg.Content,
		msg.Sender.ID,
		msg.Reciver.ID,
		msg.Time,
	)
	if err != nil {
		return err
	}
	return nil
}
