package cmd

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"math/rand"
	"time"

	"mykafka/config"
)

const (
	MESSAGE = "hello %d"
	TOPIC   = "test"
)

func StartProducer() error {
	producer := config.Producer
	defer func() {
		_ = producer.Close()
	}()
	for {
		randNum := rand.Uint64()
		oriValue, _ := fmt.Printf(MESSAGE, randNum)
		message := sarama.ProducerMessage{
			Topic:     TOPIC,
			Value:     sarama.StringEncoder(oriValue),
			Key:       sarama.StringEncoder(randNum % 10),
			Timestamp: time.Time{},
		}
		t1 := time.Now().Nanosecond()
		partition, offset, err := producer.SendMessage(&message)
		t2 := time.Now().Nanosecond()

		if err == nil {
			fmt.Println("produce success, partition:", partition, ",offset:", offset, ",cost:", (t2-t1)/(1000*1000), " ms")
		} else {
			fmt.Println(err.Error())
		}
		time.Sleep(1 * time.Second)
	}
}

type consumerHandler struct{}

func (consumerHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}
func (consumerHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}
func (c consumerHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d  value:%s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Value))
		//todo 持久化？

		sess.MarkMessage(msg, "")
	}
	return nil
}
func StartConsumer(ctx context.Context) error {
	consumerGroup := config.ConsumerGroup
	defer func() {
		_ = consumerGroup.Close()
	}()
	for {
		err := consumerGroup.Consume(ctx, []string{TOPIC}, consumerHandler{})
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}

}
