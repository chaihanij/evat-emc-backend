package assignments

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/assignments/models"
)

func (r repo) FindOneAssignment(ctx context.Context, input *entities.AssignmentFilter) (*entities.Assignment, error) {
	log.Debugln("DB FindOneAssignment")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()

	filter := models.NewAssignmentFilter(input)
	var assignment models.Assignment
	err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAssignments).
		FindOne(ctx, filter, nil).
		Decode(&assignment)
	if err != nil {
		log.WithError(err).Errorln("DB FindOneAssignment Error")
		return nil, err
	}
	// log.WithField("	value", assignment).Debugln("DB FindOneAssignment")
	fmt.Println("assignment :", assignment.Topic)
	return assignment.ToEntity()
}
