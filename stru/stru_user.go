package stru

type User struct {
	ID     int    `db:"id"`
	UserId string `db:"userId"`
	Name   string `db:"name"`
}

type Message struct {
	Content string `db:"content"`
	Sender  User   `db:"sender"`
	Reciver User   `db:"reciver"`
	Time    string `db:"time"`
}
