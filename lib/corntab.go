package lib

import (
	"database/sql"
	"linebot-server/models"
	"time"
)

func ReserveMessage(db *sql.DB, msg models.Message) {
	for {
		now := time.Now().Format("2006-01-02 15:04")
		if now == msg.Time.Format("2006-01-02 15:04") {
			PushMessage(db, msg)
			return
		}
	}
	// duration := time.Duration(time.Now().Sub(reserveTime))
	// fmt.Println(duration)
}
