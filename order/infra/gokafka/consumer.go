package gokafka

import (
	"context"
	"encoding/json"

	kafka "github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"github.com/swiggy-2022-bootcamp/cdp-team4/order/infra"
	"github.com/swiggy-2022-bootcamp/cdp-team4/order/infra/logger"
)

var KafkaBrokers = []string{
	"localhost:9092",
}

var log logrus.Logger = *logger.GetLogger()

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

	reader := getKafkaReader(ctx, topic, "id", KafkaBrokers)
	for {
		msg, err := reader.ReadMessage(ctx)

		var statusMsg PaymentRecord

		json.Unmarshal([]byte(msg.Value), &statusMsg)
		if err == nil {
			//fmt.Println(statusMsg, string(msg.Value))
			if topic == "payment" {
				userid := statusMsg.OrderID
				status := statusMsg.Status
				if status != "" && userid != "" {
					go db.UpdateOrderStatus(userid, status)
					go UpdateProductService("product", userid, db1, status)
				}
			}
		}
	}
}
