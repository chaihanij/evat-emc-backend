package teams

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/mongo"
)

func (u useCase) FindOneAssignmentTeam(ctx context.Context, input *entities.AssignmentTeamFilter) (*entities.AssignmentTeam, error) {
	assignmentTeam, err := u.AssignmentTeamsRepo.FindOneAssignmentTeam(ctx, input)
	if err != nil {
		return nil, err
	}
	if val, ok := assignmentTeam.Documents.([]string); ok {
		log.WithField("value", val).Debugln("FindOneMember documents")
		documents, err := u.FilesRepo.FindAllFile(ctx, &entities.FileFilter{UUIDs: val})
		if err != nil && mongo.ErrNoDocuments != err {
			return nil, err
		}
		assignmentTeam.Documents = documents
	}
	return assignmentTeam, nil
}
