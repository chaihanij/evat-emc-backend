package score

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/layers/repositories/files"
	"gitlab.com/chaihanij/evat/app/layers/repositories/members"
	"gitlab.com/chaihanij/evat/app/layers/repositories/users"
	"gitlab.com/chaihanij/evat/app/layers/repositories/userscore"
)

type useCase struct {
	Score       userscore.Repo
	UsersRepo   users.Repo
	MembersRepo members.Repo
	FilesRepo   files.Repo
	ScoreRepo   userscore.Repo
}

type UseCase interface {
	FindAllScore(ctx context.Context, input *entities.ScoreFilter) (*int64, []entities.Score, error) // FindAllScore(ctx context.Context, input *entities.ScoreFilter) (*int64, []entities.Score, error)
	// CreateScore(ctx context.Context, input *entities.Score) (*entities.Score, error)
}

func InitUseCase(ScoreRepo userscore.Repo, usersRepo users.Repo, membersRepo members.Repo, filesRepo files.Repo) UseCase {
	return &useCase{
		ScoreRepo:   ScoreRepo,
		Score:       ScoreRepo,
		UsersRepo:   usersRepo,
		MembersRepo: membersRepo,
		FilesRepo:   filesRepo,
	}
}
