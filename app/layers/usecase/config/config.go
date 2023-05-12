package config

import (
	"context"
	"fmt"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) FindOneConfig(ctx context.Context, input *entities.Config) (*entities.Config, error) {
	fmt.Println("eeeadsa :", input)
	conFig, err := u.ConfigRepo.FindOneConfig(ctx, input)
	if err != nil {
		return nil, err
	}

	return conFig, nil
}
