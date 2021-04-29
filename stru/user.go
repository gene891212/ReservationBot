package stru

type UserInfo struct {
	ID            int    `json:"id"`
	UserID        string `json:"userId"`
	DisplayName   string `json:"displayName"`
	PictureUrl    string `json:"pictureUrl"`
	StatusMessage string `json:"statusMessage"`
}

type Message struct {
	Content string
	Sender  UserInfo
	Reciver UserInfo
	Time    string
}

type Rerservation struct {
	AccessToken string
	Reciver     string
	SendTime    string
}
