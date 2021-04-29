package lib

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
)

func ReserveMessage(reserveTime time.Time, reciver string, content string) {
	for {
		now := time.Now().Format("2006-01-02 15:04")
		if now == reserveTime.Format("2006-01-02 15:04") {
			pushMessage()
			return
		}
	}
	// duration := time.Duration(time.Now().Sub(reserveTime))
	// fmt.Println(duration)
}

func pushMessage() {
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SCRECT"),
		os.Getenv("ACCESS_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	_, err = bot.PushMessage("Ub1eff16cb01f4343694423cba8c74e52", linebot.NewTextMessage("hello")).Do()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("okok")
}
