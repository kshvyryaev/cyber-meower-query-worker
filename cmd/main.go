package main

import (
	"fmt"

	"github.com/kshvyryaev/cyber-meower-query-worker/pkg/config"
	"github.com/kshvyryaev/cyber-meower-query-worker/pkg/event"
)

func main() {
	config := config.ProvideConfig()
	connection, _, _ := event.ProvideNatsConnection(config)
	receiver, _, _ := event.ProvideNatsMeowEventReceiver(connection)

	for {
		meowEvent := receiver.Receive()
		fmt.Println(meowEvent)
	}
}
