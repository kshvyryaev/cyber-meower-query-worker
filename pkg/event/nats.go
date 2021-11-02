package event

import (
	"github.com/kshvyryaev/cyber-meower-query-worker/pkg/config"
	"github.com/nats-io/nats.go"
	"github.com/pkg/errors"
)

func ProvideNatsConnection(config *config.Config) (*nats.EncodedConn, func(), error) {
	connection, err := nats.Connect(config.EventStoreAddress)
	if err != nil {
		return nil, nil, errors.Wrap(err, "nats connection")
	}

	encodedConnection, err := nats.NewEncodedConn(connection, nats.JSON_ENCODER)
	if err != nil {
		return nil, nil, errors.Wrap(err, "nats connection")
	}

	cleanup := func() {
		encodedConnection.Close()
	}

	return encodedConnection, cleanup, nil
}
