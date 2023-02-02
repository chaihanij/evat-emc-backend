package files

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/layers/repositories/files"
)

type useCase struct {
	FilesRepo files.Repo
}

type UseCase interface {
	FindOneFile(ctx context.Context, input *entities.FileFilter) (*entities.File, error)
}

func InitUseCase(filesRepo files.Repo) UseCase {
	return &useCase{
		FilesRepo: filesRepo,
	}
}
