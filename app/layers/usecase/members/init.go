package members

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/layers/repositories/files"
	"gitlab.com/chaihanij/evat/app/layers/repositories/logsetting"
	"gitlab.com/chaihanij/evat/app/layers/repositories/members"
)

type useCase struct {
	MembersRepo    members.Repo
	FilesRepo      files.Repo
	LogsettingRepo logsetting.Repo
}

type UseCase interface {
	CreateMember(ctx context.Context, input *entities.Member) (*entities.Member, error)
	DeleteMember(ctx context.Context, input *entities.MemberFilter) error
	FindAllMember(ctx context.Context, input *entities.MemberFilter) (*int64, []entities.Member, error)
	FindOneMember(ctx context.Context, input *entities.MemberFilter) (*entities.Member, error)
	UpdateMember(ctx context.Context, input *entities.MemberPartialUpdate) (*entities.Member, error)
	UpdateMemberImage(ctx context.Context, memberUUID string, newFile *entities.File) (*entities.File, error)
	UpdateMemberPullDocument(ctx context.Context, memberUUID string, fileUUID string) error
	UpdateMemberPushDocument(ctx context.Context, memberUUID string, file *entities.File) (*entities.File, error)
	UseCaseMemberCheckeIn(ctx context.Context, input *entities.MemberCheckIn) (*entities.Member, error)
}

func InitUseCase(membersRepo members.Repo, filesRepo files.Repo, logSettingRepo logsetting.Repo) UseCase {
	return &useCase{
		MembersRepo:    membersRepo,
		FilesRepo:      filesRepo,
		LogsettingRepo: logSettingRepo,
	}
}
