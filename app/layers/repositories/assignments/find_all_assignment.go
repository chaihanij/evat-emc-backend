package assignments

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/assignments/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r repo) getFindOptions(input *entities.AssignmentFilter) *options.FindOptions {
	findOptions := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})
	if input.PageSize != nil && input.Page != nil {
		pageSize := *input.PageSize
		page := *input.Page
		findOptions.SetLimit(*input.PageSize)
		offset := (page - 1) * pageSize
		findOptions.SetSkip(offset)
	}
	return findOptions
}

func (r repo) FindAllAssignment(ctx context.Context, input *entities.AssignmentFilter) ([]entities.Assignment, error) {
	log.Debugln("DB FindAllAssignment")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	findOptions := r.getFindOptions(input)
	filter := models.NewAssignmentFilter(input)
	cursor, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAssignments).
		Find(ctx, filter, findOptions)
	if err != nil {
		log.WithError(err).Errorln("DB FindAllAssignment Error")
		return nil, err
	}
	var assignments models.Assignments
	err = cursor.All(ctx, &assignments)
	if err != nil {
		log.WithError(err).Errorln("DB FindAllAssignment Error")
		return nil, err
	}

	log.WithField("value", assignments).Debugln("DB FindAllAssignment")
	return assignments.ToEntity(), nil
}
