package assignment_teams

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/assignment_teams/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r repo) UpdateAssignmentTeamPushDocument(ctx context.Context, input *entities.AssignmentTeamPartialUpdate, documentUUID string) (*entities.AssignmentTeam, error) {
	log.Debugln("DB UpdateAssignmentTeamPushDocument")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()

	var dliveryTime models.DliveryTime
	filterOverDue := bson.M{
		"uuid": input.AssignmentUUID,
	}
	errOverdue := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAssignments).
		FindOne(ctx, filterOverDue, nil).
		Decode(&dliveryTime)
	if errOverdue != nil {
		log.WithError(errOverdue).Errorln("DB FindOneAssignmentTeam Error")
		return nil, errOverdue
	}

	dliveryTime.TimeNow = time.Now()
	log.Debugln("dliveryTime", dliveryTime)
	dlivery := dliveryTime.DliveryTime.Unix()
	overdue := dliveryTime.DueDate.Unix()
	timeNow := dliveryTime.TimeNow.Unix()

	if dlivery >= timeNow {

		return nil, fmt.Errorf("ยังไม่ถึงเวลาส่งงาน")
	}
	if overdue <= timeNow {
		return nil, fmt.Errorf("เกินกำหนดเวลาส่งงาน")
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After).SetUpsert(true)
	filter := models.NewAssignmentTeamFilter(input)
	update := models.UpdateAssignmentTeamPushDocument(documentUUID, input.Topic)
	var assignmentTeam models.AssignmentTeam
	err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAssignmentTeams).
		FindOneAndUpdate(ctx, filter, update, opts).
		Decode(&assignmentTeam)
	if err != nil {
		log.WithError(err).Errorln("DB UpdateAssignmentTeamPushDocument Error")
		return nil, err
	}
	log.WithField("value", assignmentTeam).Debugln("DB UpdateAssignmentTeamPushDocument")
	return assignmentTeam.ToEntity()
}
