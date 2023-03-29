package visit

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/layers/repositories/visit"
)

// visit "gitlab.com/chaihanij/evat/app/layers/repositories/visit"

type useCase struct {
	VisitRepo visit.Repo
}

type UseCase interface {
	FindVisit(ctx context.Context) (*entities.Visited, error)
	CreateVisit(ctx context.Context, data *entities.UpdateVisit) (*entities.UpdateVisit, error)
}

func InitUseCase(VisitRepo visit.Repo) UseCase {
	return &useCase{
		VisitRepo: VisitRepo,
	}

}
