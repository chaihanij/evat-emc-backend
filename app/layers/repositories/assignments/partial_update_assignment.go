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

func (r repo) PartialUpdateAssignment(ctx context.Context, input *entities.AssignmentPartialUpdate) (*entities.Assignment, error) {
	log.Debugln("DB PartialUpdateMember")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	filter := models.NewAssignmentFilter(input)
	update := models.PartialUpdateAssignment(input)
	var assignment models.Assignment
	err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAssignments).
		FindOneAndUpdate(ctx, filter, update, opts).
		Decode(&assignment)
	if err != nil {
		log.WithError(err).Errorln("DB PartialUpdateAssignment Error")
		return nil, err
	}
	log.WithField("value", assignment).Debugln("DB PartialUpdateAssignment")
	return assignment.ToEntity()
}
