package consideration

import (
	"context"
	"encoding/json"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"

	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/consideration/models"
)

func (r repo) FindOneConsideration(ctx context.Context, input *entities.ConsiderationFilter) (*entities.Consideration, error) {
	log.Debugln("Consideration")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	var res []bson.M
	state := []bson.M{

		{
			"$match": bson.M{
				"team_uuid": *input.TeamUUID,
			},
		},

		// {
		// 	"$group": bson.M{
		// 		"_id":             "$team_uuid",
		// 		"indivdual_score": bson.M{"$push": bson.M{"desc": "$desc", "score": "$score"}},
		// 		"update_at":       bson.M{"$last": "$updated_at"},
		// 		"total_score":     bson.M{"$sum": "$score"},
		// 	},
		// },
		{
			"$group": bson.M{
				"_id":             "$team_uuid",
				"indivdual_score":bson.M{"$push": bson.M{"title": "$title", "score": "$full_score"}},
				"update_at":       bson.M{"$last": "$updated_at"},
				"total_score":     bson.M{"$sum": "$full_score"},
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

	var considerations models.Consideration
	err = cursor.All(ctx, &res)
	if err != nil {
		log.WithError(err).Errorln("Consideration Error")
		return nil, err
	}
	var structData []models.Consideration
	jsonData, _ := json.Marshal(res)

	json.Unmarshal(jsonData, &structData)

	considerations.ID = structData[0].ID
	considerations.TotalScore = structData[0].TotalScore
	considerations.No = structData[0].No
	considerations.IndivdualScore = structData[0].IndivdualScore
	considerations.UpdatedAt = structData[0].UpdatedAt

	//fmt.Println("data ", res)

	// considerations.ID = res[0]["_id"].(string)
	// considerations.Score = res[0]["score"].(float64)
	// considerations.UpdatedAt = res[0]["updated_at"].(time.Time)
	// log.Debug(considerations)

	log.WithField("value", considerations).Debugln("Consideration")
	return considerations.ToEntity()
}
