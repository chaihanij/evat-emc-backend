package consideration

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"

	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/consideration/models"
)

func (r repo) FindOneConsideration(ctx context.Context, input *entities.ConsiderationFilter) ([]entities.AssignmentScore, error) {
	log.Debugln("Consideration")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	// fmt.Println("id ::", *input.ID)
	// var res []bson.M
	state := []bson.M{

		// {
		// 	"$match": bson.M{
		// 		"team_uuid": *input.TeamUUID,
		// 	},
		// },
		// {
		// 	"$group": bson.M{
		// 		"_id":             "$team_uuid",
		// 		"indivdual_score":bson.M{"$push": bson.M{"title": "$title", "score": "$full_score"}},
		// 		"update_at":       bson.M{"$last": "$updated_at"},
		// 		"total_score":     bson.M{"$sum": "$full_score"},
		// 	},
		// },
		{
			"$match": bson.M{
				"uuid": *input.AssignmentUUID,
			},
		},
		{
			"$unwind": "$consideration",
		},
		{
			"$match": bson.M{
				"consideration.id": *input.ID,
			},
		},
		{
			"$project": bson.M{
				"_id":           1,
				"consideration": 1,
			},
		},
		{
			"$group": bson.M{
				"_id": nil,
				"considerations": bson.M{"$push": bson.M{
					"id":       "$consideration.id",
					"nameteam": "$consideration.nameteam",
					"title":"$consideration.title",
					"score":    "$consideration.score",
				}},
				"total": bson.M{"$sum": "$consideration.score"},
			},
		},
	}

	cursor, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAssignments).
		Aggregate(ctx, state)

	if err != nil {
		log.WithError(err).Errorln("Consideration Error")
		return nil, err
	}

	var considerations models.AssignmentScores
	// var structData []models.AssignmentScore

	err = cursor.All(ctx, &considerations)

	if err != nil {
		log.WithError(err).Errorln("Consideration Error")
		return nil, err
	}
	fmt.Println("considerations :", considerations)
	// jsonData, _ := json.Marshal(res)

	// json.Unmarshal(jsonData, &structData)

	log.WithField("value", considerations).Debugln("Consideration")
	return considerations.ToEntity(), nil
}
