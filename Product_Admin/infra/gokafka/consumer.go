package gokafka

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	kafka "github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Product_Admin/infra"
	"github.com/swiggy-2022-bootcamp/cdp-team4/Product_Admin/infra/logger"
)

var log logrus.Logger = *logger.GetLogger()

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

type ProductQuantityRecord struct {
	ProductID           string
	QuantityToBeReduced int64
}

func ProductConsumer(ctx context.Context, topic string, db infra.ProductAdminDynamoRepository) {

	reader := getKafkaReader(ctx, topic, "id", []string{"localhost:9090"})
	for {
		msg, err := reader.ReadMessage(ctx)

		var message ProductQuantityRecord

		json.Unmarshal([]byte(msg.Value), &message)
		if err != nil {
			fmt.Printf("could not read message %s\n ", err.Error())
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusBadRequest})
		}
		if topic == "cart" {
			productId := message.ProductID
			quantityToBeReduced := message.QuantityToBeReduced
			go db.UpdateQuantity(productId, quantityToBeReduced)
		}
	}
}
