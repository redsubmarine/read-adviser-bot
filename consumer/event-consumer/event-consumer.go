package event_consumer

import (
	"log"
	"read-adviser-bot/events"
	"time"
)

type Consumer struct {
	fetcher   events.Fetcher
	processor events.Processor
	batchSize int
}

func New(fetcher events.Fetcher, processor events.Processor, batchSize int) *Consumer {
	return &Consumer{
		fetcher:   fetcher,
		processor: processor,
		batchSize: batchSize,
	}
}

func (c Consumer) Start() error {
	for {
		events, err := c.fetcher.Fetch(c.batchSize)
		if err != nil {
			log.Printf("[ERR] consumer: %s", err.Error())
			continue
		}

		if len(events) == 0 {
			time.Sleep(1 * time.Second)
			continue
		}

		if err := c.handleEvents(events); err != nil {
			log.Printf("[ERR] consumer: %s", err.Error())
			continue
		}
	}
}

func (c *Consumer) handleEvents(events []events.Event) error {
	for _, e := range events {
		log.Printf("[INFO] got new event: %s", e.Text)
		if err := c.processor.Process(e); err != nil {
			log.Printf("[ERR] can't handle event: %s", err.Error())
			continue
		}

	}
	return nil
}
