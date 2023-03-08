package main

import (
	"flag"
	"log"

	"read-adviser-bot/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	tgClient := telegram.New(tgBotHost, mustToken())
	log.Println(tgClient)
}

func mustToken() string {
	token := flag.String("token-bot-token", "", "Telegram bot token")
	flag.Parse()

	if *token == "" {
		log.Fatal("token is required")
	}

	return *token
}
