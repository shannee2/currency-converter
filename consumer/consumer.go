package consumer

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"grpc_currency_converter/dao"
)

type CurrencyUpdateMessage struct {
	Currency string  `json:"currency"`
	NewRate  float64 `json:"newRate"`
}

func StartKafkaConsumer(broker, topic, groupID string, dao *dao.CurrencyDAOImpl) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
	}

	err = consumer.Subscribe(topic, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to topic %s: %v", topic, err)
	}

	fmt.Println("Kafka Consumer started, listening for messages...")

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			var update CurrencyUpdateMessage
			if err := json.Unmarshal(msg.Value, &update); err != nil {
				log.Printf("Failed to parse message: %v", err)
				continue
			}

			err := dao.UpdateConversionRate(update.Currency, update.NewRate)
			if err != nil {
				log.Printf("Error updating database: %v", err)
			}
		} else {
			log.Printf("Consumer error: %v\n", err)
		}
	}
}
