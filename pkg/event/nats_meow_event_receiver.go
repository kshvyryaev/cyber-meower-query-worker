package event

import (
	eventContract "github.com/kshvyryaev/cyber-meower-event/pkg/event"
	"github.com/nats-io/nats.go"
	"github.com/pkg/errors"
)

type NatsMeowEventReceiver struct {
	connection *nats.EncodedConn
	channel    <-chan *eventContract.MeowCreatedEvent
}

func ProvideNatsMeowEventReceiver(connection *nats.EncodedConn) (*NatsMeowEventReceiver, func(), error) {
	channel := make(chan *eventContract.MeowCreatedEvent)

	_, err := connection.BindRecvChan(eventContract.MeowCreatedEventSubject, channel)
	if err != nil {
		return nil, nil, errors.Wrap(err, "nats meow event receiver")
	}

	cleanup := func() {
		close(channel)
	}

	return &NatsMeowEventReceiver{connection: connection, channel: channel}, cleanup, nil
}

func (receiver *NatsMeowEventReceiver) Receive() <-chan *eventContract.MeowCreatedEvent {
	return receiver.channel
}
