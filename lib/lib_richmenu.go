package lib

import (
	"log"
	"os"

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
				Action: linebot.RichMenuAction{
					Type: "message",
					Text: "Good morning",
				},
			},
		},
	}

	res, err := bot.CreateRichMenu(richMenuConfig).Do()
	if err != nil {
		log.Fatalln(err)
	}

	_, err = bot.UploadRichMenuImage(res.RichMenuID, "static/richmenu.png").Do()
	if err != nil {
		log.Fatalln(err)
	}

	_, err = bot.SetDefaultRichMenu(res.RichMenuID).Do()
	if err != nil {
		log.Fatalln(err)
	}
}
