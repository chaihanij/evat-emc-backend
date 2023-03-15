package userscore

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	models "gitlab.com/chaihanij/evat/app/layers/repositories/userscore/models"

	// "gitlab.com/chaihanij/evat/app/layers/repositories/score/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r repo) GetFindOptions(input *entities.ScoreFilter) *options.FindOptions {
	findOptions := options.Find().SetSort(bson.D{{Key: "createdate", Value: -1}})
	if input.PageSize != nil && input.Page != nil {
		pageSize := *input.PageSize
		page := *input.Page
		findOptions.SetLimit(*input.PageSize)
		offset := (page - 1) * pageSize
		findOptions.SetSkip(offset)
	}
	// fmt.Println("pageSize", input.PageSize)
	return findOptions
}

func (r repo) FindAllScore(ctx context.Context, input *entities.ScoreFilter) ([]entities.Score, error) {
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	findOptions := r.GetFindOptions(input)

	filter := models.NewScoreFilter(input)

	cursor, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionScore).
		Find(ctx, filter, findOptions)

	if err != nil {
		log.WithError(err).Errorln("DB FindAllscore Error -- --  - at")
		return nil, err
	}

	var score models.Scores

	err = cursor.All(ctx, &score)
	if err != nil {
		log.WithError(err).Errorln("DB FindAllscore Error |||||||| at")
		return nil, err
	}

	return score.ToEntity(), nil
}

// func (r repo) CountScore(ctx context.Context, input *entities.ScoreFilter) (*int64, error) {
// 	log.Debugln("DB Score")
// 	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
// 	defer cancel()

// 	filter := models.NewScoreFilter(input)
// 	count, err := r.MongoDBClient.Database(env.MongoDBName).
// 		Collection(constants.CollectionScore).
// 		CountDocuments(ctx, filter)

// 	if err != nil {
// 		log.WithError(err).Errorln("DB Countscore Error")
// 		return nil, err
// 	}
// 	log.WithField("value", count).Debugln("DB Score")
// 	return &count, nil
// }
