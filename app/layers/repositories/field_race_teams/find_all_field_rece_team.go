package field_race_teams

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/field_race_teams/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r repo) getFindOptions(input *entities.FieldRaceTeamFilter) *options.FindOptions {
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

func (r repo) FindAllFieldRaceTeams(ctx context.Context, input *entities.FieldRaceTeamFilter) ([]entities.FieldRaceTeam, error) {
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	findOptions := r.getFindOptions(input)
	filter := models.NewFieldRaceTeamFilter(input)
	cursor, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionFieldRaceTeams).
		Find(ctx, filter, findOptions)
	if err != nil {
		log.WithError(err).Errorln("DB CollectionFieldRaceTeams Error")
		return nil, err
	}
	var FieldRaceTeams models.FieldRaceTeams
	err = cursor.All(ctx, &FieldRaceTeams)
	if err != nil {
		log.WithError(err).Errorln("DB CollectionFieldRaceTeams Error")
		return nil, err
	}
	return FieldRaceTeams.ToEntity(), nil
}
