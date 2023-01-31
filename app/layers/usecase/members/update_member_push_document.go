package members

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) UpdateMemberPushDocument(ctx context.Context, memberUUID string, file *entities.File) (*entities.File, error) {
	file, err := u.FilesRepo.CreateFile(ctx, file)
	if err != nil {
		return nil, err
	}
	_, err = u.MembersRepo.PushDocument(ctx, memberUUID, file.UUID)
	if err != nil {
		return nil, err
	}
	return file, nil
}
