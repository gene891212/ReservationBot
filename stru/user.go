package stru

type User struct {
	ID   int
	UID  string
	Name string
}

type UserInfo struct {
	UserID        string `json:"userId"`
	DisplayName   string `json:"displayName"`
	PictureUrl    string `json:"pictureUrl"`
	StatusMessage string `json:"statusMessage"`
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
