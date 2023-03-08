package main

import (
	"flag"
	"log"
)

func main() {
	t := mustToken()
	log.Println(t)
}

func mustToken() string {
	token := flag.String("token-bot-token", "", "Telegram bot token")
	flag.Parse()

	if *token == "" {
		log.Fatal("token is required")
	}

	return *token
}
