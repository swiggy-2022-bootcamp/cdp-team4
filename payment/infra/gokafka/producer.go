package gokafka

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/swiggy-2022-bootcamp/cdp-team4/payment/infra/logger"

	kafka "github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

var brokers = []string{
	"localhost:9091",
}

var log logrus.Logger = *logger.GetLogger()

// function to generate random key, it's implementation is considering
// nano seconds of time to make strong random key.
func getRandomKey() []byte {
	var src = rand.NewSource(time.Now().UnixNano())
	return []byte(fmt.Sprint(src.Int63()))
}

// constructor to get the kafka writer with given brokers list and topic
func getProducer(ctx context.Context, topic string, brokers []string) *kafka.Writer {
	// intialize the writer with the broker addresses, and the topic
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers: brokers,
		Topic:   topic,
	})
}

// function to produce msg to kafka
func WriteMsgToKafka(topic string, msg interface{}) (bool, error) {
	ctx := context.Background()
	jsonString, err := json.Marshal(msg)
	msgString := string(jsonString)

	if err != nil {
		return false, err
	}

	writer := getProducer(ctx, topic, brokers)
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
