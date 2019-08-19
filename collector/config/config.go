package config

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type Configuration struct {
	Server ServerConfiguration
	Kafka  KafkaConfiguration
}

type ServerConfiguration struct {
	Port int
}

type KafkaConfiguration struct {
	Hosts string
	Topic KafkaTopicConfiguration
}

type KafkaTopicConfiguration struct {
	Name              string `mapstructure:"name"`
	NumPartitions     int    `mapstructure:"num-partitions"`
	ReplicationFactor int    `mapstructure:"replication-factor"`
}

func LoadConfiguration(log echo.Logger) Configuration {
	viper.SetConfigName("config")
	viper.AddConfigPath("config")

	var configuration Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	return configuration
}
