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
	log.Debugln("Call FindAllFieldRaceTeams")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()

	state := []bson.M{
		{
			"$lookup": bson.M{
				"from":         "field_races",
				"localField":   "field_race_uuid",
				"foreignField": "uuid",
				"as":           "field_races",
			},
		},
		{
			"$lookup": bson.M{
				"from":         "teams",
				"localField":   "team_uuid",
				"foreignField": "uuid",
				"as":           "teams",
			},
		},
		{
			"$unwind": "$teams",
		},
		{
			"$project": bson.M{
				"uuid":                    1,
				"field_race_uuid":         1,
				"team_uuid":               1,
				"score":                   1,
				"updated_at":              1,
				"created_at":              1,
				"created_by":              1,
				"updated_by":              1,
				"field_races.title":       1,
				"field_races.description": 1,
				"field_races.image":       1,
				"field_races.file":        1,
				"field_races.year":        1,
				"field_races.full_score":  1,
				"name":                    "$teams.name",
				"code":                    "$teams.code",
				"type":                    "$teams.team_type",
			},
		},
		{
			"$sort": bson.M{
				"_id": -1,
			},
		},
	}

	cursor, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionFieldRaceTeams).
		Aggregate(ctx, state)

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
