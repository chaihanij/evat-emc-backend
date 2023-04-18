package consideration

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/consideration/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (r repo) AllScore(ctx context.Context, input entities.AllScoreFilter) ([]entities.AllScore, error) {
	log.Debugln("Consideration")

	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()

	// filter := models.AllScoreFilter(input)
	// statefieldRaces := []bson.M{

	// 	// {
	// 	// 	"$unwind": "$consideration",
	// 	// },
	// 	// {
	// 	// 	"$group": bson.M{
	// 	// 		"_id": "$consideration.id",
	// 	// 		"name": bson.M{
	// 	// 			"$first": "$consideration.nameteam",
	// 	// 		},
	// 	// 		"considerations": bson.M{
	// 	// 			"$push": bson.M{
	// 	// 				"title": "$consideration.title",
	// 	// 				"score": "$consideration.score",
	// 	// 				"type":  "field_races",
	// 	// 			}},
	// 	// 	},
	// 	// },

	// 	{
	// 		"$unwind": "$consideration",
	// 	},
	// 	{
	// 		"$group": bson.M{
	// 			"_id": "$consideration.id",
	// 			"name": bson.M{
	// 				"$first": "$consideration.nameteam",
	// 			},
	// 			"total":bson.M{
	// 				"$sum":"$consideration.score",
	// 			 },
	// 			"considerations": bson.M{
	// 				"$push": bson.M{
	// 					"title": "$consideration.title",
	// 					"score": "$consideration.score",
	// 					"type":  "field_races",
	// 				}},
	// 		},
	// 	},

	// 	// filter,
	// }

	// var fieldRaces models.AllScoreConsiderations
	// FieldRace, err := r.MongoDBClient.Database(env.MongoDBName).
	// 	Collection(constants.CollectionFieldRaces).
	// 	Aggregate(ctx, statefieldRaces)
	// if err != nil {
	// 	log.WithError(err).Errorln("FieldRace Error")
	// 	return nil, err
	// }
	// err = FieldRace.All(ctx, &fieldRaces)

	// if err != nil {
	// 	log.WithError(err).Errorln("FieldRace-Error")
	// 	return nil, err
	// }

	stateAssignment := []bson.M{

		{
			"$unwind": "$consideration",
		},
		{
			"$group": bson.M{
				"_id": "$consideration.id",
				"name": bson.M{
					"$first": "$consideration.nameteam",
				},
				"total": bson.M{
					"$sum": "$consideration.score",
				},
				"considerations": bson.M{
					"$push": bson.M{
						"title": "$consideration.title",
						"score": "$consideration.score",
						"type":  "assignment",
					}},
			},
		},
		// filter,
	}

	var assignments models.AllScoreConsiderations
	assignment, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAssignments).
		Aggregate(ctx, stateAssignment)
	if err != nil {
		log.WithError(err).Errorln("assignment Error")
		return nil, err
	}
	err = assignment.All(ctx, &assignments)

	if err != nil {
		log.WithError(err).Errorln("assignment Error")
		return nil, err
	}

	var AllScore models.AllScoreConsiderations

	// for _, value := range fieldRaces {
	// 	AllScore = append(AllScore, value)
	// }
	for _, value := range assignments {
		AllScore = append(AllScore, value)
	}

	return AllScore.ToEntity(), nil
}
