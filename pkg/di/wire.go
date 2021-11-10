//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/kshvyryaev/cyber-meower-query-worker/pkg"
	"github.com/kshvyryaev/cyber-meower-query-worker/pkg/contract"
	"github.com/kshvyryaev/cyber-meower-query-worker/pkg/event"
	"github.com/kshvyryaev/cyber-meower-query-worker/pkg/search"
	"github.com/kshvyryaev/cyber-meower-query-worker/pkg/worker"
)

func InitializeMeowSeederWorker() (*worker.MeowSeederWorker, func(), error) {
	panic(wire.Build(
		pkg.ProvideConfig,
		pkg.ProvideZap,
		event.ProvideNats,
		event.ProvideNatsMeowEventReceiver,
		wire.Bind(new(contract.MeowEventReceiver), new(*event.NatsMeowEventReceiver)),
		search.ProvideElastic,
		search.ProvideElasticMeowRepository,
		wire.Bind(new(contract.MeowRepository), new(*search.ElasticMeowRepository)),
		worker.ProvideMeowSeederWorker,
	))
}
