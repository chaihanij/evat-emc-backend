package assignments

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/assignments/models"
)

func (r repo) CreateAssignment(ctx context.Context, input *entities.Assignment) (*entities.Assignment, error) {
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	assignment := models.NewAssignment(input)
	result, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAssignments).
		InsertOne(ctx, assignment)
	if err != nil {
		log.WithError(err).Errorln("DB CreateAssignment Error")
		return nil, err
	}
	if result.InsertedID == 0 {
		return nil, err
	}
	log.WithField("value", assignment).Debugln("DB CreateAssignment")
	return assignment.ToEntity()
}
