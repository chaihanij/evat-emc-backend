package field_races

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/field_races/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r repo) getTeamFieldRaceFindOptions(input *entities.FieldRaceFilter) *options.FindOptions {
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
func (r repo) FindTeamAllFieldRace(ctx context.Context, input *entities.FieldRaceFilter) ([]entities.FieldRace, error) {
	log.Debugln("DB FindAllFieldRace")

	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	findOptions := r.getTeamFieldRaceFindOptions(input)
	filter := models.NewFieldRaceFilter(input)

	// filter := bson.M{
	// 	// "uuid": *input.UUID,
	// }
	cursor, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionFieldRaces).
		Find(ctx, filter, findOptions)
	if err != nil {
		log.WithError(err).Errorln("DB FindAllFieldRace Error")
		return nil, err
	}
	var fieldRaces models.FieldRaces
	err = cursor.All(ctx, &fieldRaces)
	if err != nil {
		log.WithError(err).Errorln("DB FindAllFieldRace Error")
		return nil, err
	}
	//log.WithField("value", fieldRaces).Debugln("DB FindAllFieldRace")
	return fieldRaces.ToEntity(), nil
}
