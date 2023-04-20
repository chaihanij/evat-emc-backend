package consideration

import (
	"context"
	"fmt"

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

	fmt.Println("input :", input)

	page := 1
	if input.Page != 0 {
		page = input.Page
	}

	pagesize := 10
	if input.Pagesize >= 10 {
		pagesize = input.Pagesize
	}

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
			},
		},
		{
			"$group": bson.M{
				"_id":   bson.M{"id": "$consideration.nameteam", "title": "$title"},
				"team":  bson.M{"$first": "$consideration.nameteam"},
				"code":  bson.M{"$first": "$consideration.id"},
				"title": bson.M{"$first": "$title"},
				"considerations": bson.M{"$push": bson.M{
					"id":       "$consideration.id",
					"nameteam": "$consideration.nameteam",
					"title":    "$title",
					"score":    "$consideration.score",
				}},
				"total": bson.M{"$sum": "$consideration.score"},
			},
		},
		{
			"$group": bson.M{
				"_id":  bson.M{"team": "$team"},
				"team": bson.M{"$first": "$team"},
				"code": bson.M{"$first": "$code"},
				"considerations": bson.M{"$push": bson.M{
					"title": "$title",
					"total": bson.M{"$sum": "$considerations.score"},
				}},
				"total": bson.M{"$sum": "$total"},
			},
		},
		filter,
		{
			"$sort": bson.M{
				"total": -1,
			},
		},
		{
			"$skip": page - 1,
		},
		{
			"$limit": pagesize,
		},
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
