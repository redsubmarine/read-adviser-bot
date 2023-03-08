package main

import (
	"context"
	"log"

	tgClient "read-adviser-bot/clients/telegram"
	"read-adviser-bot/config"
	event_consumer "read-adviser-bot/consumer/event-consumer"
	"read-adviser-bot/events/telegram"
	"read-adviser-bot/storage/sqlite"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "files_storage"
	batchSize   = 100
)

func main() {
	cfg := config.MustLoad()
	// s := files.New(storagePath)
	s, err := sqlite.New(cfg.DBFilePath)
	if err != nil {
		log.Fatal("can't connect to storage: ", err)
	}

	if err := s.Init(context.TODO()); err != nil {
		log.Fatal("can't init storage: ", err)
	}

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, cfg.TgBotToken),
		s,
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}
