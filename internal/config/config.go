package config

import "github.com/caarlos0/env/v6"

// Config ...
type Config struct {
	Port           string `env:"PORT" envDefault:"18080"`
	Topic          string `env:"Topic" envDefault:"kafka_accident"`
	TopicDDS       string `env:"Topic" envDefault:"kafka_dds"`
	Broker1Address string `env:"Broker1Address" envDefault:"localhost:9092"`
	Broker2Address string `env:"Broker2Address" envDefault:"localhost:9093"`
	Broker3Address string `env:"Broker3Address" envDefault:"localhost:9094"`
	GroupID        string `env:"GroupID" envDefault:"my-group"`
	// Declare Connection
	DataManagementServiceConnection string `env:"DATA_MANAGEMENT_CONNECTION" envDefault:"127.0.0.1:8082"`
}

// NewConfig ...
func NewConfig() *Config {
	c := new(Config)
	if err := env.Parse(c); err != nil {
		panic(err)
	}
	return c
}
