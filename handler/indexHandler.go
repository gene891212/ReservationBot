package handler

import (
	"linebot-server/stru"
	"os"

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
	// db, err := sql.Open("mysql", config.FormatDSN())
	// if err != nil {
	// 	panic(err)
	// }
	// allUser := lib.AllUserFromDB(db)
	// c.HTML(200, "index.tmpl", gin.H{
	// 	"items": allUser,
	// })
	c.HTML(200, "index.tmpl", gin.H{
		"items": []stru.User{
			{
				UID:  "something",
				Name: "Ian",
			},
			{
				UID:  "ok",
				Name: "Gene",
			},
		},
	})
}
