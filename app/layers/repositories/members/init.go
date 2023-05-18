package members

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	MongoDBClient *mongo.Client
}

type Repo interface {
	Config() ([]string, error)
	CountMember(ctx context.Context, input *entities.MemberFilter) (*int64, error)
	CreateMember(ctx context.Context, input *entities.Member) (*entities.Member, error)
	DeleteOneMemeber(ctx context.Context, input *entities.MemberFilter) error
	FindAllMember(ctx context.Context, input *entities.MemberFilter) ([]entities.Member, error)
	FindOneMember(ctx context.Context, input *entities.MemberFilter) (*entities.Member, error)
	PartialUpdateMember(ctx context.Context, input *entities.MemberPartialUpdate) (*entities.Member, error)
	UpdateMember(ctx context.Context, input *entities.Member) (*entities.Member, error)
	PushDocument(ctx context.Context, uuid string, input string) (*entities.Member, error)
	PullDocument(ctx context.Context, uuid string, input string) (*entities.Member, error)
	MemberCheckIn(ctx context.Context, input *entities.MemberCheckIn) (*entities.Member, error)
	UploadPDFMember(ctx context.Context, input *entities.MemberUpdatePDF) (*entities.Member, error)
}

func InitRepo(mongoDBClient *mongo.Client) Repo {
	return &repo{
		MongoDBClient: mongoDBClient,
	}
}
