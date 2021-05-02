package models

type Message struct {
	Content string
	Sender  User
	Reciver User
	Time    string
}

// func GetMessage(db *sql.DB, user User) []Message {

// }

// func CreateMessage(db *sql.DB, msg Message) {

// }
