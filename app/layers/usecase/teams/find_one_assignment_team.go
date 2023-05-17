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

	var doc []entities.File
	for _, val := range assignmentTeam.Document {
		documents, _ := u.FilesRepo.FindAllFile(ctx, &entities.FileFilter{UUID: &val.DocumentUUID})
		doc = append(doc, documents...)
		assignmentTeam.Documents = doc
	}

	if val, ok := assignmentTeam.Documents.([]string); ok {
		log.WithField("value", val).Debugln("FindOneMember documents")
		if len(val) > 0 {
			documents, err := u.FilesRepo.FindAllFile(ctx, &entities.FileFilter{UUIDs: val})
			if err != nil && mongo.ErrNoDocuments != err {
				return nil, err
			}
			assignmentTeam.Documents = documents
		}
	}

	return assignmentTeam, nil
}
