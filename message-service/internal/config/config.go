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

type KafkaConfig struct {
	BrokerAddresses string `yaml:"brokerAddresses"` // localhost:9092,localhost:9093
	MessageConsumerGroupID string `yaml:"messageConsumerGroupId"`
}

type CassandraConfig struct {
}

type Config struct {
	Environment string `yaml:"environment"`
	Port int `yaml:"port"`
	LogDir string `yaml:"logDir"`
	Kafka KafkaConfig `yaml:"kafka"`
	Cassandra CassandraConfig `yaml:"cassandra"`
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