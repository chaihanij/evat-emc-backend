package users

import (
	"context"

	"github.com/AlekSi/pointer"
	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

func (u useCase) ResetPassword(ctx context.Context, input *entities.ResetPassword) (*entities.User, error) {
	log.WithField("input", input).Debugln("ResetPassword UC")
	user, err := u.UsersRepo.FindOneUser(ctx, &entities.UserFilter{ActivateCode: &input.ActivateCode})
	if err != nil {
		if mongo.ErrNoDocuments == err {
			return nil, errors.RecordNotFoundError{Message: constants.DataNotFound}
		}
		return nil, err
	}

	if input.ConfirmPassword != input.NewPassword {
		return nil, errors.ParameterError{Message: "password not match"}
	}

	userPartialUpdate := &entities.UserPartialUpdate{
		UID:      user.UID,
		Password: pointer.ToString(input.ConfirmPassword),
	}

	return u.UsersRepo.PartialUpdateUser(ctx, userPartialUpdate)
}
