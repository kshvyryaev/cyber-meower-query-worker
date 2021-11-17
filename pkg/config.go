package pkg

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type Config struct {
	EventStoreAddress string `envconfig:"EVENT_STORE_ADDRESS"`
	ElasticAddress    string `envconfig:"ELASTIC_ADDRESS"`
}

func ProvideConfig() (*Config, error) {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		return nil, errors.Wrap(err, "config")
	}
	return &config, nil
}
