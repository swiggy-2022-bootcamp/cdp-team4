package gokafka

import (
	"context"
	"encoding/json"
	"fmt"

	kafka "github.com/segmentio/kafka-go"
	"github.com/swiggy-2022-bootcamp/cdp-team4/reward/infra"
)

var brokers = []string{
	"localhost:9092",
}

func getKafkaReader(ctx context.Context, topic, groupID string, brokers []string) *kafka.Reader {
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		StartOffset: kafka.LastOffset,
	})
	return r
}

type RewardMessage struct{
	UserID string `json:"user_id"`
	Points int	`json:"points"`
}

func UpdateRewardPoints(ctx context.Context, topic string, db infra.RewardDynamoRepository) {
	reader := getKafkaReader(ctx, topic, "Update-Reward", brokers)
	for {
		msg, err := reader.ReadMessage(ctx)

		var rMsg RewardMessage

		json.Unmarshal([]byte(msg.Value), &rMsg)
		if err != nil {
			fmt.Printf("could not read message %s\n ", err.Error())
		}
		if topic == "payment" {
			UserId := rMsg.UserID
			Points := rMsg.Points
			go db.UpdateRewardByUserId(UserId, Points)
		}
	}
}
