package assignments

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/mongo"
)

func (u useCase) UpdateAssignment(ctx context.Context, input *entities.AssignmentPartialUpdate) (*entities.Assignment, error) {
	assignment, err := u.AssignmentsRepo.PartialUpdateAssignment(ctx, input)
	if err != nil {
		return nil, err
	}
	if val, ok := assignment.Image.(string); ok {
		img, err := u.FilesRepo.FindOneFile(ctx, &entities.FileFilter{UUID: &val})
		if err != nil && mongo.ErrNoDocuments != err {
			return nil, err
		}
		if img != nil {
			assignment.Image = *img
		}
	}

	if val, ok := assignment.Document.(string); ok {
		doc, err := u.FilesRepo.FindOneFile(ctx, &entities.FileFilter{UUID: &val})
		if err != nil && mongo.ErrNoDocuments != err {
			return nil, err
		}
		if doc != nil {
			assignment.Document = *doc
		}
	}

	log.WithField("value", assignment).Debugln("UseCase FindOneAssignments")
	return assignment, nil
}
