package search

import (
	elastic "github.com/elastic/go-elasticsearch/v7"
	"github.com/kshvyryaev/cyber-meower-query-worker/pkg"
	"github.com/pkg/errors"
)

func ProvideElastic(config *pkg.Config) (*elastic.Client, func(), error) {
	elasticConfig := elastic.Config{
		Addresses: []string{config.ElasticAddress},
	}
	elastic, err := elastic.NewClient(elasticConfig)
	if err != nil {
		return nil, nil, errors.Wrap(err, "elastic")
	}

	infoResult, err := elastic.Info()
	if err != nil {
		return nil, nil, errors.Wrap(err, "elastic")
	}

	cleanup := func() {
		infoResult.Body.Close()
	}

	return elastic, cleanup, nil
}
