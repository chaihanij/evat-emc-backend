package consideration

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/consideration/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r repo) AllScore(ctx context.Context, input entities.AllScoreFilter) ([]entities.AllScore, error) {
	log.Debugln("Consideration")

	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()

	filter := models.AllScoreFilter(input)
	filterTeamType := models.AllScoreTeamtype(input)

	oo, _ := primitive.ObjectIDFromHex("63eb10556bcc684ff7cb7210")
	ooa, _ := primitive.ObjectIDFromHex("6458c099f27f388312ca12e9")

	stateAssignment := []bson.M{
		{
			"$match": bson.M{
				"_id": bson.M{
					"$nin": bson.A{oo, ooa},
				},
			},
		},
		{
			"$unwind": "$consideration",
		},
		{
			"$project": bson.M{
				"_id":           1,
				"consideration": 1,
				"title":         1,
			},
		},
		{
			"$group": bson.M{
				"_id":       bson.M{"id": "$consideration.nameteam", "title": "$title"},
				"team":      bson.M{"$first": "$consideration.nameteam"},
				"code":      bson.M{"$first": "$consideration.id"},
				"title":     bson.M{"$first": "$title"},
				"team_type": bson.M{"$first": "$consideration.teamtype"},
				"considerations": bson.M{
					"$push": bson.M{
						"id":        "$consideration.id",
						"nameteam":  "$consideration.nameteam",
						"team_type": "$consideration.teamtype",
						"title":     "$title",
						"score":     "$consideration.score",
					},
				},
				"total": bson.M{"$sum": "$consideration.score"},
			},
		},
		{
			"$group": bson.M{
				"_id":      bson.M{"team": "$team"},
				"team":     bson.M{"$first": "$team"},
				"code":     bson.M{"$first": "$code"},
				"teamtype": bson.M{"$first": "$team_type"},
				"considerations": bson.M{
					"$push": bson.M{
						"title": "$title",
						"total": bson.M{"$sum": "$considerations.score"},
					},
				},
				"total": bson.M{"$sum": "$total"},
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
						"$denseRank": bson.M{},
					},
				},
			},
		},
		{
			"$sort": bson.M{
				"teamtype": -1,
			},
		},
		filter,
		filterTeamType,
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
