package config

import (
	"os"
	"path"
	"runtime"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type Environment int32

const (
	Development Environment = iota
	Production
	Testing
)

func (e Environment) String() string {
	switch e {
	case Development:
		return "DEVELOPMENT"
	case Production:
		return "PRODUCTION"
	case Testing:
		return "TESTING"
	}
	return "unknown"
}

type RedisConfig struct {
	HostAddress string `yaml:"hostAddress"` // localhost:5672
	DB 			int `yaml:"db"` 
}

type KafkaConfig struct {
	BrokerAddresses string `yaml:"brokerAddresses"` // localhost:9092,localhost:9093
}

type UserClient struct {
	HostAddress string `yaml:"hostAddress"` // localhost:5672
}

type ChatServer struct {
	Port 			int 	`yaml:"port"`
	PresencePath 	string 	`yaml:"presencePath"`
}

type RabbitMQConfig struct {
	HostAddress string `yaml:"hostAddress"` // localhost:5672
}


type Config struct {
	Environment 	string 			`yaml:"environment"`
	Port 			int 			`yaml:"port"`
	LogDir 			string 			`yaml:"logDir"`
	Redis 			RedisConfig 	`yaml:"redis"`
	UserClient 		UserClient 		`yaml:"userClient"`
	ChatServer		ChatServer		`yaml:"chatServer"`
	Kafka 			KafkaConfig 	`yaml:"kafka"`
	RabbitMQ 		RabbitMQConfig 	`yaml:"rabbitmq"`
}

var syncOnceConfig sync.Once
var config *Config

func ProvideConfig() (*Config, error) {
	var err error
	syncOnceConfig.Do(func() {
		err = readConfigFromFile()
	})
	if err != nil {
		return config, err
	}
	return config, nil
}

func readConfigFromFile() error {
	env := strings.ToUpper(os.Getenv("GO_ENV"))
	_, filename, _, _ := runtime.Caller(0)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path.Dir(filename))

	if env == "TESTING" {
		viper.SetConfigName("config.testing")
	} else if env == "PRODUCTION" {
		viper.SetConfigName("config.production")
	} else if env == "STAGING" {
		viper.SetConfigName("config.staging")
	} else {
		viper.SetConfigName("config.development")
	}

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	viper.SetDefault("environment", Development.String())
	viper.SetDefault("port", 80)
	viper.SetDefault("logDir", path.Join(path.Dir(filename), "../../log"))

	err = viper.Unmarshal(&config)
	if err != nil {
		return err
	}
	return nil
}