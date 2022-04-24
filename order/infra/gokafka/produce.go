package gokafka

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/swiggy-2022-bootcamp/cdp-team4/order/infra"

	kafka "github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type OrderOverviewKafkaModel struct {
	OrderID      string         `json:"order_id"`
	ProductIdQty map[string]int `json:"products"`
	Status       string         `json:"status"`
}

func getRandomKey() []byte {
	var src = rand.NewSource(time.Now().UnixNano())
	return []byte(fmt.Sprint(src.Int63()))
}

func getProducer(ctx context.Context, topic string, brokers []string) *kafka.Writer {
	// intialize the writer with the broker addresses, and the topic
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers: brokers,
		Topic:   topic,
	})
}

func WriteMsgToKafka(topic string, msg interface{}) (bool, error) {
	ctx := context.Background()
	jsonString, err := json.Marshal(msg)
	msgString := string(jsonString)

	if err != nil {
		return false, err
	}

	writer := getProducer(ctx, topic, KafkaBrokers)
	defer writer.Close()

	err = writer.WriteMessages(ctx, kafka.Message{
		Key:   getRandomKey(),
		Value: []byte(msgString),
	})

	if err != nil {
		log.WithFields(logrus.Fields{"error": err.Error(), "msg": msg}).
			Error("could not write message ")
		return false, err
	}

	return true, nil
}

func UpdateProductService(topic string, orderid string, db infra.OrderDynamoRepository, status string) {

	res, err := db.GetOrderOverview(orderid)
	if err != nil {
		log.WithFields(logrus.Fields{"error": err.Error(), "msg": "Error reading Order Overview Table"}).
			Error("could not Read order id ")
		return
	}
	newKafaMsg := OrderOverviewKafkaModel{
		OrderID:      res.OrderID,
		ProductIdQty: res.ProductsIdQuantity,
		Status:       status,
	}
	WriteMsgToKafka(topic, newKafaMsg)
}
