package members

import (
	"context"
)

func (u useCase) UpdateMemberPullDocument(ctx context.Context, memberUUID string, fileUUID string) error {

	_, err := u.MembersRepo.PullDocument(ctx, memberUUID, fileUUID)
	if err != nil {
		return err
	}
	return nil
}
