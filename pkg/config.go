package pkg

import (
	"flag"
)

type Config struct {
	EventStoreAddress string
}

func ProvideConfig() *Config {
	return &Config{
		EventStoreAddress: *flag.String("eventStoreAddress", "127.0.0.1:4222", "Event store address"),
	}
}
