package worker

import (
	"sync"

	"github.com/kshvyryaev/cyber-meower-query-worker/pkg/domain"
	"github.com/kshvyryaev/cyber-meower-query-worker/pkg/event"
	"github.com/kshvyryaev/cyber-meower-query-worker/pkg/search"
	"go.uber.org/zap"
)

type MeowSeederWorker struct {
	meowEventReceiver event.MeowEventReceiver
	meowRepository    search.MeowRepository
	logger            *zap.Logger
}

func ProvideMeowSeederWorker(
	meowEventReceiver event.MeowEventReceiver,
	meowRepository search.MeowRepository,
	logger *zap.Logger) *MeowSeederWorker {
	return &MeowSeederWorker{
		meowEventReceiver: meowEventReceiver,
		meowRepository:    meowRepository,
		logger:            logger,
	}
}

func (worker *MeowSeederWorker) Run(wg *sync.WaitGroup) {
	defer wg.Done()
	for meowEvent := range worker.meowEventReceiver.Receive() {
		worker.logger.Info("Event received from nats",
			zap.Int("id", meowEvent.ID),
			zap.String("body", meowEvent.Body),
			zap.Time("created_on", meowEvent.CreatedOn))

		meow := &domain.Meow{
			ID:        meowEvent.ID,
			Body:      meowEvent.Body,
			CreatedOn: meowEvent.CreatedOn,
		}

		if err := worker.meowRepository.Create(meow); err != nil {
			worker.logger.Error("Event didn't insert into elastic: "+err.Error(),
				zap.Int("id", meowEvent.ID),
				zap.String("body", meowEvent.Body),
				zap.Time("created_on", meowEvent.CreatedOn))
			continue
		}

		worker.logger.Info("Event inserted into elastic",
			zap.Int("id", meowEvent.ID),
			zap.String("body", meowEvent.Body),
			zap.Time("created_on", meowEvent.CreatedOn))
	}
}
