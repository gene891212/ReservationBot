package handler

import (
	"linebot-server/stru"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

var (
	config = mysql.Config{
		User:   os.Getenv("user"),
		Passwd: os.Getenv("passwd"),
		Net:    os.Getenv("tcp"),
		Addr:   os.Getenv("addr"),
		DBName: os.Getenv("dbname"),
	}
)

func LiffPage(c *gin.Context) {
	now := time.Now()
	// db, err := sql.Open("mysql", config.FormatDSN())
	// if err != nil {
	// 	panic(err)
	// }
	// allUser := lib.AllUserFromDB(db)
	// c.HTML(200, "index.tmpl", gin.H{
	// 	"users": allUser,
	// 	"now": struct {
	// 		Date string
	// 		Time string
	// 	}{
	// 		Date: time.Now().Format("2006-01-02"),
	// 		Time: time.Now().Format("15:04"),
	// 	},
	// })

	// Dev data
	c.HTML(200, "index.tmpl", gin.H{
		"users": []stru.User{
			{
				UID:  "something",
				Name: "Ian",
			},
			{
				UID:  "ok",
				Name: "Gene",
			},
		},
		"now": struct {
			Date string
			Time string
		}{
			Date: now.Format("2006-01-02"),
			Time: now.Format("15:04"),
		},
	})
}
