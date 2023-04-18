package assignments

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/assignments/models"
)

func (r repo) FindTopicAssignment(ctx context.Context, input *entities.AssignmentFilter) (*entities.ExportAssignmentTopic, error) {
	log.Debugln("DB FindTopic")
	ctx, cancle := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancle()

	var topic models.ExportAssignmentTopic
	filter := models.NewAssignmentFilter(input)

	err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAssignments).
		FindOne(ctx, filter, nil).
		Decode(&topic)
	if err != nil {
		log.WithError(err).Errorln("DB FindOneAssignment Error")
		return nil, err
	}

	return topic.ToEntity()
}
