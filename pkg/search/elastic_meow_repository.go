package search

import (
	"bytes"
	"context"
	"encoding/json"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/google/wire"
	"github.com/kshvyryaev/cyber-meower-query-worker/pkg/domain"
	"github.com/pkg/errors"
)

const _meowIndex = "meows"

type ElasticMeowRepository struct {
	client *elasticsearch.Client
}

func ProvideElasticMeowRepository(client *elasticsearch.Client) *ElasticMeowRepository {
	return &ElasticMeowRepository{
		client: client,
	}
}

func (repository *ElasticMeowRepository) Create(meow *domain.Meow) error {
	json, err := json.Marshal(meow)
	if err != nil {
		return errors.Wrap(err, "elastic meow repository")
	}

	request := esapi.IndexRequest{
		Index:      _meowIndex,
		DocumentID: strconv.Itoa(meow.ID),
		Body:       bytes.NewReader(json),
		Refresh:    "true",
	}

	response, err := request.Do(context.Background(), repository.client)
	if err != nil {
		return errors.Wrap(err, "elastic meow repository")
	}
	defer response.Body.Close()

	return nil
}

var ElasticMeowRepositorySet = wire.NewSet(
	ProvideElasticMeowRepository,
	wire.Bind(new(MeowRepository), new(*ElasticMeowRepository)),
)
