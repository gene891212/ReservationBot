package stru

type User struct {
	ID   int
	UID  string
	Name string
}

type Message struct {
	Content string
	Sender  User
	Reciver User
	Time    string
}

type Rerservation struct {
	AccessToken string
	Reciver     string
	SendTime    string
}
