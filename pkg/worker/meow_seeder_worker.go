package worker

import (
	"github.com/kshvyryaev/cyber-meower-query-worker/pkg/event"
	"go.uber.org/zap"
)

type MeowSeederWorker struct {
	meowEventReceiver event.MeowEventReceiver
	logger            *zap.Logger
}

func ProvideMeowSeederWorker(meowEventReceiver event.MeowEventReceiver, logger *zap.Logger) *MeowSeederWorker {
	return &MeowSeederWorker{
		meowEventReceiver: meowEventReceiver,
		logger:            logger,
	}
}

func (worker *MeowSeederWorker) Run() {
	for meowEvent := range worker.meowEventReceiver.Receive() {
		worker.logger.Info("Event received from nats",
			zap.Int("id", meowEvent.ID),
			zap.String("body", meowEvent.Body),
			zap.Time("created_on", meowEvent.CreatedOn))
	}
}
