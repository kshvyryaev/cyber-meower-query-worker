package event

import eventContract "github.com/kshvyryaev/cyber-meower-event/pkg/event"

type MeowEventReceiver interface {
	Receive() *eventContract.MeowCreatedEvent
}
