package consideration

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/layers/repositories/consideration"
)

type useCase struct {
	ConsiderationRepo consideration.Repo
}

type UseCase interface {
	FindOneConsideration(ctx context.Context, input *entities.ConsiderationFilter) ([]entities.AssignmentScore, error)
}

func InitUseCase(ConsiderationRepo consideration.Repo) UseCase {
	return &useCase{
		ConsiderationRepo: ConsiderationRepo,
	}
}
