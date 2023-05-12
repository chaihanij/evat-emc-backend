package config

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) UpdateConfigs(ctx context.Context, input *entities.Config) (*entities.Config, error) {
	return u.ConfigRepo.UpdateConfig(ctx, input)
}
