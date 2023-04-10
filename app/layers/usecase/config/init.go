package config

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/layers/repositories/config"
)

type useCase struct {
	ConfigRepo config.Repo
}

type UseCase interface {
	FindOneConfig(ctx context.Context, input *entities.Config) (*entities.Config, error)
}

func InitUseCase(config config.Repo) UseCase {
	return &useCase{
		ConfigRepo: config,
	}
}
