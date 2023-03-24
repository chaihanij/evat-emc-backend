package assignment_teams

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/assignment_teams/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (r repo) FindAllscoreAssignmentTeam(ctx context.Context, input *entities.AssignmentTeamFilter) ([]entities.AssignmentTeamScore, error) {
	log.Debugln("DB FindAllAssignmentTeam")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	// findOptions := r.getFindOptions(input)
	// log.Debugln("input :: 1234", *input.TeamUUID)
	filter := []bson.M{
		// {
		// 	"$match": bson.M{
		// 		"team_uuid": *input.TeamUUID,
		// 	},
		// },
		// {
		// 	"$lookup": bson.M{
		// 		"from":         "assignments",
		// 		"localField":   "assignment_uuid",
		// 		"foreignField": "uuid",
		// 		"as":           "assignments",
		// 	},
		// },
		// {
		// 	"$unwind": "$assignments",
		// },
		// {
		// 	"$project": bson.M{
		// 		"_id":         1,
		// 		"title":       "$assignments.title",
		// 		"assignments": "$assignments.full_score",
		// 	},
		// },
		{
			"$match": bson.M{
				"team_uuid": *input.TeamUUID,
			},
		},
		//		{
		//			"$lookup": {
		//				"from":         "assignments",
		//				"localField":   "assignment_uuid",
		//				"foreignField": "uuid",
		//				"as":           "assignments",
		//			},
		//		},
		//		{
		//			"$unwind": "$assignments",
		//		},
		{
			"$lookup": bson.M{
				"from":         "field_race_teams",
				"localField":   "team_uuid",
				"foreignField": "team_uuid",
				"as":           "field_race_teams",
			},
		},
		{
			"$unwind": "$field_race_teams",
		},

		{
			"$lookup": bson.M{
				"from":         "field_races",
				"localField":   "field_race_teams.field_race_uuid",
				"foreignField": "uuid",
				"as":           "field_races",
			},
		},
		{
			"$unwind": "$field_races",
		},
		{
			"$project": bson.M{
				"_id": 1,
				// 				"title":       "$assignments.title",
				// 				"assignments": "$assignments.full_score",
				// 				"field_races":1,
				"title":      "$field_races.title",
				"full_score": "$field_races.full_score",
			},
		},
	}
	cursor, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAssignmentTeams).
		Aggregate(ctx, filter)
		// Find(ctx, filter, findOptions)
	if err != nil {
		log.WithError(err).Errorln("DB FindAllAssignmentTeam Error")
		return nil, err
	}
	var fieldRacesTeamScore models.AssignmentTeamScores
	err = cursor.All(ctx, &fieldRacesTeamScore)
	if err != nil {
		log.WithError(err).Errorln("DB FindAllAssignmentTeam Error")
		return nil, err
	}

	var assignment models.AssignmentTeamScores
	for _, value := range fieldRacesTeamScore {
		// var assignmentTeamScores []models.AssignmentTeamScores

		fieldRaceTeamScores := models.AssignmentTeamScore{
			Title: value.Title,
			Score: value.Score,
		}

		assignment = append(assignment, fieldRaceTeamScores)
	}

	filter_assignment := []bson.M{
		{
			"$match": bson.M{
				"team_uuid": *input.TeamUUID,
			},
		},
		{
			"$lookup": bson.M{
				"from":         "assignments",
				"localField":   "assignment_uuid",
				"foreignField": "uuid",
				"as":           "assignments",
			},
		},
		{
			"$unwind": "$assignments",
		},
		{
			"$project": bson.M{
				"_id":        1,
				"full_score": "$assignments.full_score",
				"title":      "$assignments.title",
			},
		},
	}

	cursorassignment, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAssignmentTeams).
		Aggregate(ctx, filter_assignment)
	if err != nil {
		log.WithError(err).Errorln("DB FindAllAssignmentTeam Error")
		return nil, err
	}
	var assignmentTeamsScore models.AssignmentTeamScores
	err = cursorassignment.All(ctx, &assignmentTeamsScore)
	if err != nil {
		log.WithError(err).Errorln("DB FindAllAssignmentTeam Error")
		return nil, err
	}

	for _, value := range assignmentTeamsScore {
		assignmentsTeamScores := models.AssignmentTeamScore{
			Title: value.Title,
			Score: value.Score,
		}
		assignment = append(assignment, assignmentsTeamScores)
	}

	// var assignment models.AssignmentTeamScores
	// log.WithField("value", assignment).Debugln("DB FindAllAssignmentTeam")
	return assignment.ToEntity(), nil
}
