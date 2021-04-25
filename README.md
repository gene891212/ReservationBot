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
dbname=
```
Enter the command to start server

```bash
go run server.go
```

## Reference

- [Linebot Messaging API reference](https://developers.line.biz/en/reference/messaging-api/)
- [line-bot-sdk-go github](https://github.com/line/line-bot-sdk-go)
- [line-bot-sdk-go documentation](https://pkg.go.dev/github.com/line/line-bot-sdk-go@v7.8.0+incompatible/linebot)
