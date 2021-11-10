package contract

import "github.com/kshvyryaev/cyber-meower-query-worker/pkg/domain"

type MeowRepository interface {
	Create(meow *domain.Meow) error
}
