package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

const accessToken = "LBP_ACCESS_TOKEN"
const secretToken = "LBP_SECRET_TOKEN"

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
			leftBtn := linebot.NewMessageAction("left", "left clicked")
			rightBtn := linebot.NewMessageAction("right", "right clicked")
			template := linebot.NewConfirmTemplate("Hello World", leftBtn, rightBtn)
			message := linebot.NewTemplateMessage("Hi", template)

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

func run(port string) error {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/ping", pingHandler)
	return http.ListenAndServe(":"+port, nil)
}

func main() {
	port := flag.String("p", "8080", "port")
	flag.Parse()
	if err := run(*port); err != nil {
		log.Fatalln(err)
	}
}
