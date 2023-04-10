package assignments

import (
	"context"
	// "time"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/assignments/models"
	// "go.mongodb.org/mongo-driver/bson"
)

func (r repo) UploadScoreAssignment(ctx context.Context, input *entities.Assignment) (*entities.Assignment, error) {
	log.Debugln("DB UpdateAssignment")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	filter := models.NewAssignmentFilter(input)
	update := models.UploadScoreAssignment(input)
	var assignment models.Assignment
	err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAssignments).
		FindOneAndUpdate(ctx, filter, update, nil).
		Decode(&assignment)
	if err != nil {
		log.WithError(err).Errorln("DB UpdateAssignmentError Error")
		return nil, err
	}
	log.WithField("value", assignment).Debugln("DB UpdateAssignment")
	return assignment.ToEntity()
}
