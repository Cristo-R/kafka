package config

import (
	"fmt"
	env "github.com/caarlos0/env/v6"
)

const ServiceName = "Mykafka"

type cfg struct {
	Env      string `env:"ENV" envDefault:"dev"`
	LogLevel string `env:"LOG_LEVEL" envDefault:"debug"`

	KafkaHost []string `env:"KAFKA_HOST" envDefault:"localhost:9092"`
}

var Cfg *cfg

func init() {

	Cfg = &cfg{}
	if err := env.Parse(Cfg); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", Cfg)
}
