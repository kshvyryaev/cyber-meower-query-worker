package search

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/kshvyryaev/cyber-meower-query-worker/pkg"
	"github.com/pkg/errors"
)

func ProvideElastic(config *pkg.Config) (*elasticsearch.Client, error) {
	elasticConfig := elasticsearch.Config{
		Addresses: []string{config.ElasticAddress},
	}
	elastic, err := elasticsearch.NewClient(elasticConfig)
	if err != nil {
		return nil, errors.Wrap(err, "elastic")
	}

	infoResult, err := elastic.Info()
	if err != nil {
		return nil, errors.Wrap(err, "elastic")
	}
	defer infoResult.Body.Close()

	return elastic, nil
}
