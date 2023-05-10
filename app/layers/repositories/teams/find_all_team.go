package teams

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/teams/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r repo) getFindOptions(input *entities.TeamFilter) *options.FindOptions {
	findOptions := options.Find().SetSort(bson.D{{Key: "code", Value: -1}})
	if input.PageSize != nil && input.Page != nil {
		pageSize := *input.PageSize
		page := *input.Page
		findOptions.SetLimit(*input.PageSize)
		offset := (page - 1) * pageSize
		findOptions.SetSkip(offset)
	}
	return findOptions
}

func (r repo) FindAllTeam(ctx context.Context, input *entities.TeamFilter) ([]entities.Team, error) {
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	findOptions := r.getFindOptions(input)
	filter := models.NewTeamFilter(input)
	cursor, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionTeams).
		Find(ctx, filter, findOptions)
	if err != nil {
		log.WithError(err).Errorln("DB FindAllTeam Error")
		return nil, err
	}
	var teams models.Teams
	err = cursor.All(ctx, &teams)
	if err != nil {
		log.WithError(err).Errorln("DB FindAllTeam Error")
		return nil, err
	}
	return teams.ToEntity(), nil
}

// func (r repo) FindAllTeam(ctx context.Context, input *entities.TeamFilter) ([]entities.Team, error) {
// 	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
// 	defer cancel()

// 	pagination := bson.M{"$match": bson.M{}}
// 	skip := bson.M{"$match": bson.M{}}
// 	if input.PageSize != nil && input.Page != nil {
// 		pageSize := *input.PageSize
// 		page := *input.Page
// 		offset := (page - 1) * pageSize

// 		pagination = bson.M{
// 			"$limit": *input.PageSize,
// 		}
// 		skip = bson.M{
// 			"$skip": offset,
// 		}

// 	}

// 	match := bson.M{
// 		"$match": bson.M{},
// 	}
// 	if input.Name != nil {
// 		fmt.Println("input :", *input.Name)
// 		match = bson.M{
// 			"$match": bson.M{
// 				"users.tel": bson.M{
// 					"$regex": *input.Name,
// 				},
// 			},
// 		}
// 	}

// 	filter := []bson.M{
// 		{
// 			"$lookup": bson.M{
// 				"from":         "users",
// 				"localField":   "uuid",
// 				"foreignField": "team_uuid",
// 				"as":           "users",
// 			},
// 		},
// 		{
// 			"$unwind": "$users",
// 		},
// 		match,
// 		pagination,
// 		skip,
// 	}
// 	cursor, err := r.MongoDBClient.Database(env.MongoDBName).
// 		Collection(constants.CollectionTeams).
// 		// Find(ctx, filter, findOptions)
// 		Aggregate(ctx, filter)

// 	if err != nil {
// 		log.WithError(err).Errorln("DB FindAllTeam Error")
// 		return nil, err
// 	}
// 	var teams models.Teams
// 	err = cursor.All(ctx, &teams)
// 	if err != nil {
// 		log.WithError(err).Errorln("DB FindAllTeam Error")
// 		return nil, err
// 	}
// 	return teams.ToEntity(), nil
// }
