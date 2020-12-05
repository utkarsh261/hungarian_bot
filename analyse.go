package main

import (
	"log"

	tbot "github.com/go-telegram-bot-api/telegram-bot-api"
)

func analyse(ID int64, msg *tbot.Message) {
	log.Print(msg.Text)

}
