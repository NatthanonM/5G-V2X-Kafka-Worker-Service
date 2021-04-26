package config

import "github.com/caarlos0/env/v6"

// Config ...
type Config struct {
	Port           string `env:"PORT" envDefault:"18080"`
	Topic          string `env:"TOPIC_ACCIDENT" envDefault:"kafka_accident"`
	TopicDDS       string `env:"TOPIC_DROWSINESS" envDefault:"kafka_dds"`
	Broker1Address string `env:"BROKER_1_ADDRESS" envDefault:"localhost:9093"`
	Broker2Address string `env:"BROKER_2_ADDRESS" envDefault:"localhost:9094"`
	GroupID        string `env:"GROUP_ID" envDefault:"my-group"`
	// Declare Connection
	DataManagementServiceConnection string `env:"DATA_MANAGEMENT_CONNECTION" envDefault:"127.0.0.1:8082"`
	UsernameKafka                   string `env:"USERNAME_KAFKA" envDefault:""`
	PasswordKafka                   string `env:"PASSWORD_KAFKA" envDefault:""`
}

// NewConfig ...
func NewConfig() *Config {
	c := new(Config)
	if err := env.Parse(c); err != nil {
		panic(err)
	}
	return c
}
