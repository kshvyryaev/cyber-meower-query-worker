//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/kshvyryaev/cyber-meower-query-worker/pkg"
	"github.com/kshvyryaev/cyber-meower-query-worker/pkg/event"
	"github.com/kshvyryaev/cyber-meower-query-worker/pkg/search"
	"github.com/kshvyryaev/cyber-meower-query-worker/pkg/worker"
)

func InitializeMeowSeederWorker() (*worker.MeowSeederWorker, func(), error) {
	panic(wire.Build(
		pkg.ProvideConfig,
		pkg.ProvideZap,
		event.ProvideNats,
		event.NatsMeowEventPublisherSet,
		search.ProvideElastic,
		search.ElasticMeowRepositorySet,
		worker.ProvideMeowSeederWorker,
	))
}
