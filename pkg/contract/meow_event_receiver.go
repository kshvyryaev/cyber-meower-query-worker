package contract

import eventContract "github.com/kshvyryaev/cyber-meower-event/pkg/event"

type MeowEventReceiver interface {
	Receive() <-chan *eventContract.MeowCreatedEvent
}
