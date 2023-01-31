package members

import (
	"context"

	"github.com/AlekSi/pointer"
	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) UpdateMemberImage(ctx context.Context, memberUUID string, newFile *entities.File) (*entities.File, error) {
	file, err := u.FilesRepo.CreateFile(ctx, newFile)
	if err != nil {
		return nil, err
	}
	_, err = u.MembersRepo.PartialUpdateMember(ctx, &entities.MemberPartialUpdate{UUID: memberUUID, Image: pointer.ToString(file.UUID)})
	if err != nil {
		return nil, err
	}
	return file, nil
}
