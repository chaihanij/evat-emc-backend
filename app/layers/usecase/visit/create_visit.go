package visit

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
)


func (u useCase) CreateVisit(ctx context.Context, data *entities.UpdateVisit ) (*entities.UpdateVisit, error) {
	return u.VisitRepo.CreateVisit(ctx, data)
}