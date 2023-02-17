package config

import (
	"github.com/Shopify/sarama"
)

var Producer sarama.SyncProducer

func InitProducer() {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewHashPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	var err error
	Producer, err = sarama.NewSyncProducer(Cfg.KafkaHost, config)
	if err != nil {
		panic(err)
	}
}
