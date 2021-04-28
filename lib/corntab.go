package lib

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func ReserveMessage(reserveTime time.Time) {
	c := cron.New()
	c.AddFunc("@every 1s", func() {
		fmt.Println(reserveTime)
	})
	c.Start()
}
