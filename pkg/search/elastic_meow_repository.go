package search

import (
	"bytes"
	"encoding/json"
	"strconv"

	elastic "github.com/elastic/go-elasticsearch/v7"
	"github.com/google/wire"
	"github.com/kshvyryaev/cyber-meower-query-worker/pkg/domain"
	"github.com/pkg/errors"
)

const _meowIndex = "meows"

type ElasticMeowRepository struct {
	client *elastic.Client
}

func ProvideElasticMeowRepository(client *elastic.Client) *ElasticMeowRepository {
	return &ElasticMeowRepository{
		client: client,
	}
}

func (repository *ElasticMeowRepository) Create(meow *domain.Meow) error {
	json, err := json.Marshal(meow)
	if err != nil {
		return errors.Wrap(err, "elastic meow repository")
	}

	_, err = repository.client.Index(
		_meowIndex,
		bytes.NewReader(json),
		repository.client.Index.WithDocumentID(strconv.Itoa(meow.ID)),
		repository.client.Index.WithRefresh("wait_for"),
	)

	if err != nil {
		return errors.Wrap(err, "elastic meow repository")
	}

	return nil
}

var ElasticMeowRepositorySet = wire.NewSet(
	ProvideElasticMeowRepository,
	wire.Bind(new(MeowRepository), new(*ElasticMeowRepository)),
)
