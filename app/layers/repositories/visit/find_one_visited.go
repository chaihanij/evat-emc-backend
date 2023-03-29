package visit

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/layers/repositories/visit/models"
	"go.mongodb.org/mongo-driver/bson"

	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
)

func (r repo) FindOneVisited(ctx context.Context) (*entities.Visited, error) {
	log.Debugln("DB Visit")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	t := time.Now()

	timestart_date := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	timeend_date := time.Date(t.Year(), t.Month(), t.Day(), 17, 59, 59, 59, time.UTC)

	filter := bson.M{
		"create_at": bson.M{
			"$gte": timestart_date,
			"$lt":  timeend_date,
		},
	}
	filterTotal := bson.M{}

	connect := r.MongoDBClient.Database(env.MongoDBName).Collection(constants.Collectionvisit)

	todayVisit, err := connect.CountDocuments(ctx, filter)
	if err != nil {
		log.WithError(err).Errorln("DB FindVisitToday Error")
		return nil, err
	}
	totalVisited, err := connect.CountDocuments(ctx, filterTotal)

	if err != nil {
		log.WithError(err).Errorln("DB FindVisitTotal Error")
		return nil, err
	}

	var visit models.Visit

	visit.TodayVisit = int(todayVisit)
	visit.TotalVisited = int(totalVisited)

	log.WithField("value", visit).Debugln("DB FindVisit")
	return visit.ToEntity()
}
