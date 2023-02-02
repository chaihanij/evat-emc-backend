package files

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) FindOneFile(ctx context.Context, input *entities.FileFilter) (*entities.File, error) {

	return u.FilesRepo.FindOneFile(ctx, input)
}
