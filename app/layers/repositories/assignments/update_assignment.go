package assignments

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/assignments/models"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r repo) UpdateAssignment(ctx context.Context, input *entities.Assignment) (*entities.Assignment, error) {
	log.Debugln("DB UpdateAssignment")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	filter := models.NewAssignmentFilter(input)
	update := models.UpdateAssignment(input)
	var assignment models.Assignment
	err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAssignments).
		FindOneAndUpdate(ctx, filter, update, opts).
		Decode(&assignment)
	if err != nil {
		log.WithError(err).Errorln("DB UpdateAssignment Error")
		return nil, err
	}
	log.WithField("value", assignment).Debugln("DB UpdateAssignment")
	return assignment.ToEntity()
}
