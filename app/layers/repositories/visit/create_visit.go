package visit

import (
	"context"

	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/visit/models"
)

func (r repo) CreateVisit(ctx context.Context, input *entities.UpdateVisit) (*entities.UpdateVisit, error) {
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()

	visit := models.NewVisit(input)

	_, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.Collectionvisit).
		InsertOne(ctx, visit)

	if err != nil {
		return nil, err
	}

	return visit.ToEntity()

}
