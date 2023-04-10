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

func (r repo) FindConsiderationFieldRaceTeam(ctx context.Context, input *entities.ConsiderationFilter) ([]entities.FieldRaceTeamScore, error) {

	log.Debugln("ConsiderationFieldRaceTeam")

	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	state := []bson.M{
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
					"title":    "$consideration.title",
					"score":    "$consideration.score",
				}},
				"total": bson.M{"$sum": "$consideration.score"},
			},
		},
	}

	cursor, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionFieldRaces).
		Aggregate(ctx, state)
	if err != nil {
		log.WithError(err).Errorln("Consideration Error")
		return nil, err
	}

	var considerationFieldRaceTeams models.FieldRaceTeamScores
	err = cursor.All(ctx, &considerationFieldRaceTeams)
	if err != nil {
		log.WithError(err).Errorln("Consideration Error")
		return nil, err
	}
	log.WithField("value", considerationFieldRaceTeams).Debugln("ConsiderationFieldRaceTeam")
	return considerationFieldRaceTeams.ToEntity(), nil

}
