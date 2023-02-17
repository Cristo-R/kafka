package config

import "github.com/Shopify/sarama"

var ConsumerGroup sarama.ConsumerGroup

func InitConsumer() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = false
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	config.Version = sarama.V0_10_2_0
	var err error
	ConsumerGroup, err = sarama.NewConsumerGroup(Cfg.KafkaHost, "testgroup", config)
	if err != nil {
		panic(err)
	}
}
