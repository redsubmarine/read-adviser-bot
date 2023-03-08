package config

import (
	"flag"
	"log"
	"path/filepath"
)

const sqliteStoragePath = "data/sqlite/"

type Config struct {
	TgBotToken string
	DBFilePath string
}

func MustLoad() Config {
	tgBotToken := flag.String("tg-bot-token", "", "Telegram bot token")

	dbFileName := flag.String("db-file-name", "db-test.sqlite", "DB file name")

	flag.Parse()

	if *tgBotToken == "" {
		log.Fatal("token is required")
	}

	return Config{
		TgBotToken: *tgBotToken,
		DBFilePath: filepath.Join(sqliteStoragePath, *dbFileName),
	}
}
