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

	filter := models.AllScoreFilter(input)
	filterTeamType := models.AllScoreTeamtype(input)

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
				"_id":       bson.M{"id": "$consideration.nameteam", "title": "$title"},
				"team":      bson.M{"$first": "$consideration.nameteam"},
				"code":      bson.M{"$first": "$consideration.id"},
				"title":     bson.M{"$first": "$title"},
				"team_type": bson.M{"$first": "$consideration.teamtype"},
				"no":        bson.M{"$first": "$consideration.no"},
				"considerations": bson.M{"$push": bson.M{
					"id":        "$consideration.id",
					"nameteam":  "$consideration.nameteam",
					"team_type": "$consideration.teamtype",
					"title":     "$title",
					"score":     "$consideration.score",
				}},
				"total": bson.M{"$sum": "$consideration.score"},
			},
		},
		{
			"$group": bson.M{
				"_id":      bson.M{"team": "$team"},
				"team":     bson.M{"$first": "$team"},
				"code":     bson.M{"$first": "$code"},
				"teamtype": bson.M{"$first": "$team_type"},
				"no":       bson.M{"$first": "$no"},
				"considerations": bson.M{"$push": bson.M{
					"title": "$title",
					"total": bson.M{"$sum": "$considerations.score"},
				}},
				"total": bson.M{"$sum": "$total"},
			},
		},
		filter,
		filterTeamType,
		{
			"$sort": bson.M{
				"total": -1,
			},
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

	// idx := 0
	// total := 0.0
	// // var AllScoreAssignments models.AllScoreConsiderations

	// for index, value := range assignments {
	// 	fmt.Println("value :", value.Total)

	// 	if value.Total >= total {
	// 		idx += 1
	// 	}
	// 	// value.No = idx
	// 	assignments[index].No = idx
	// 	fmt.Println("idx", idx)

	// }

	return assignments.ToEntity(), nil
}
