package gokafka

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"
	"time"

	kafka "github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/assert"
)

func getRandomKey() []byte {
	src := rand.NewSource(time.Now().UnixNano())
	return []byte(fmt.Sprint(src))
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

	writer := getProducer(ctx, topic, brokers)
	defer writer.Close()

	err = writer.WriteMessages(ctx, kafka.Message{
		Key:   getRandomKey(),
		Value: []byte(msgString),
	})

	if err != nil {
		fmt.Print(err.Error())
		return false, err
	}

	return true, nil
}

func TestProduceKafkaPaymentStatus(t *testing.T) {
	newpayment := PaymentRecord{
		Amount:   133,
		Currency: "INR",
		UserID:   "123",
		OrderID:  "9b7e877e-7995-43fb-a5d3-c7efecd6d32f",
		Status:   "Declined",
		Notes:    []string{"string", "string2"},
	}
	res, err := WriteMsgToKafka("payment", newpayment)
	assert.Equal(t, res, true)
	assert.Nil(t, err)
}
