package main

import (
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

const accessToken = "LBP_ACCESS_TOKEN"
const secretToken = "LBP_SECRET_TOKEN"
const envPort = "PORT"

var bot *linebot.Client

func init() {
	st := os.Getenv(secretToken)
	at := os.Getenv(accessToken)
	botTmp, err := linebot.New(st, at)
	if err != nil {
		panic("init error")
	}

	bot = botTmp
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			message := linebot.NewTextMessage("Hello World")
			_, err := bot.ReplyMessage(event.ReplyToken, message).Do()
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func run() error {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/ping", pingHandler)

	port := os.Getenv(envPort)
	return http.ListenAndServe(":"+port, nil)
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}
