package lib

import (
	"database/sql"
	"time"
)

func ReserveMessage(db *sql.DB, reserveTime time.Time, reciver string, content string) {
	for {
		now := time.Now().Format("2006-01-02 15:04")
		if now == reserveTime.Format("2006-01-02 15:04") {
			PushMessage(db, reciver, content)
			return
		}
	}
	// duration := time.Duration(time.Now().Sub(reserveTime))
	// fmt.Println(duration)
}
