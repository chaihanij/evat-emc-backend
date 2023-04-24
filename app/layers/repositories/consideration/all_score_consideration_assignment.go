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

func (r repo) AllScoreConsiderationAssignment(ctx context.Context, input entities.AllScoreFilter) ([]entities.AllScore, error) {

	log.Debugln("Consideration")

	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()

	filter := models.AllScoreFilter(input)

	stateAssignment := []bson.M{

		{
			"$unwind": "$consideration",
		},
		{
			"$project": bson.M{
				"_id":           1,
				"consideration": 1,
				"title":         1,
				"uuid":          1,
			},
		},
		{
			"$group": bson.M{
				"_id":      bson.M{"id": "$consideration.nameteam", "title": "$title"},
				"team":     bson.M{"$first": "$consideration.nameteam"},
				"code":     bson.M{"$first": "$consideration.id"},
				"title":    bson.M{"$first": "$title"},
				"uuid":     bson.M{"$first": "$uuid"},
				"teamtype": bson.M{"$first": "$consideration.teamtype"},
				"considerations": bson.M{
					"$push": bson.M{
						"id":        "$consideration.id",
						"nameteam":  "$consideration.nameteam",
						"team_type": "$consideration.teamtype",
						"title":     "$title",
						"total":     "$consideration.score",
					},
				},
				"total": bson.M{"$sum": "$consideration.score"},
			},
		},
		{
			"$setWindowFields": bson.M{
				"partitionBy": "$teamtype",
				"sortBy": bson.M{
					"total": -1,
				},
				"output": bson.M{
					"no": bson.M{
						"$rank": bson.M{},
					},
				},
			},
		},
		filter,
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

	return assignments.ToEntity(), nil

}
