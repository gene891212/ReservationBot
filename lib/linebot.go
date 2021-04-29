package lib

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

func SetupRichMenu() {
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SCRECT"),
		os.Getenv("ACCESS_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	richMenuConfig := linebot.RichMenu{
		Size: linebot.RichMenuSize{
			Width:  1660,
			Height: 500,
		},
		Selected:    true,
		Name:        "default",
		ChatBarText: "查看更多功能",
		Areas: []linebot.AreaDetail{
			{
				Bounds: linebot.RichMenuBounds{
					X:      0,
					Y:      0,
					Width:  1660,
					Height: 500,
				},
				// Action: linebot.RichMenuAction{
				// 	Type: "message",
				// 	Text: "Good morning",
				// },
				Action: linebot.RichMenuAction{
					Type: "uri",
					URI:  os.Getenv("LIFF_URL"),
				},
			},
		},
	}

	res, err := bot.CreateRichMenu(richMenuConfig).Do()
	if err != nil {
		log.Fatal(err)
	}

	var path string

	if gin.Mode() == gin.TestMode {
		path = "../static/linebot/richmenu.png"
	} else {
		path = "static/linebot/richmenu.png"
	}
	_, err = bot.UploadRichMenuImage(res.RichMenuID, path).Do()
	if err != nil {
		log.Fatal(err)
	}

	_, err = bot.SetDefaultRichMenu(res.RichMenuID).Do()
	if err != nil {
		log.Fatal(err)
	}
}

func PushMessage(db *sql.DB, reciver string, content string) {
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SCRECT"),
		os.Getenv("ACCESS_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	userProfile := GetUser(db, reciver)

	_, err = bot.PushMessage(userProfile.UserID, linebot.NewTextMessage(content)).Do()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
