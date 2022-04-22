package gokafka

import (
	"context"
	"encoding/json"
	"fmt"

	kafka "github.com/segmentio/kafka-go"
	"github.com/swiggy-2022-bootcamp/cdp-team4/order/infra"
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
	})
	return r
}

type PaymentRecord struct {
	Amount   int
	Currency string
	UserID   string
	OrderID  string
	Status   string
	Notes    []string
}

func StatusConsumer(ctx context.Context, topic string, db infra.OrderDynamoRepository, db1 infra.OrderDynamoRepository) {

	reader := getKafkaReader(ctx, topic, "id", []string{"localhost:9092"})
	for {
		msg, err := reader.ReadMessage(ctx)

		var statusMsg PaymentRecord

		json.Unmarshal([]byte(msg.Value), &statusMsg)
		if err != nil {
			fmt.Printf("could not read message %s\n ", err.Error())
		}
		fmt.Println(statusMsg, string(msg.Value))
		if topic == "payment" {
			orderid := statusMsg.OrderID
			status := statusMsg.Status
			go db.UpdateOrderStatus(orderid, status)
			go UpdateProductService("product", orderid, db1)
		}
	}
}
