# ReservationBot

Using Golang to build a linebot with liff

## TO-DO

- [ ] Reservation message

## Usage

Create a `.env` file and enter the following below to add environment variables

```
# Linebot

CHANNEL_SCRECT=
ACCESS_TOKEN=
LIFF_URL=

# Mysql

user=
passwd=
net=
addr=
dbname=linebot
```

Database setting:

```sql
-- Create database
CREATE DATABASE linebot;

-- Create table
CREATE TABLE Users (
	ID int NOT NULL AUTO_INCREMENT,
	UserID varchar(50),
	DisplayName varchar(50),
	PictureURL varchar(1000),
	PRIMARY KEY (ID)
);

CREATE TABLE Messages (
	ID int NOT NULL AUTO_INCREMENT,
	Content varchar(4000),
	Sender int,
	Reciver int,
	Time datetime,
	PRIMARY KEY (ID),
	FOREIGN KEY (Sender) REFERENCES Users(ID),
	FOREIGN KEY (Reciver) REFERENCES Users(ID)
);
```

Enter the command to start server

```bash
go run server.go
```

## Reference

- [gin github](https://github.com/gin-gonic/gin)
- [Linebot Messaging API reference](https://developers.line.biz/en/reference/messaging-api/)
- [line-bot-sdk-go github](https://github.com/line/line-bot-sdk-go)
- [line-bot-sdk-go documentation](https://pkg.go.dev/github.com/line/line-bot-sdk-go@v7.8.0+incompatible/linebot)
