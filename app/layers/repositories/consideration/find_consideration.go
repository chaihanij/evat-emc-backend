package consideration

import (
	"context"

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

		{
			"$group": bson.M{
				"_id":        "$team_uuid",
				"updated_at": bson.M{"$last": "$updated_at"},
				"score":      bson.M{"$sum": "$score"},
			},
		},
	}

	cursor, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionFieldRaceTeams).
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

	considerations.ID = res[0]["_id"].(string)
	considerations.Score = res[0]["score"].(float64)
	// considerations.UpdatedAt = res[0]["updated_at"].(time.Time)
	// log.Debug(considerations)

	log.WithField("value", considerations).Debugln("Consideration")
	return considerations.ToEntity()
}
