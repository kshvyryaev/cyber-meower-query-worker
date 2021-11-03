package pkg

import (
	"flag"
)

type Config struct {
	EventStoreAddress string
	ElasticAddress    string
}

func ProvideConfig() *Config {
	return &Config{
		EventStoreAddress: *flag.String("eventStoreAddress", "127.0.0.1:4222", "Event store address"),
		ElasticAddress:    *flag.String("elasticAddress", "http://127.0.0.1:9200", "Elastic search address"),
	}
}
