package users

import (
	"context"

	"github.com/AlekSi/pointer"
	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func (u useCase) ChangePassword(ctx context.Context, input *entities.ResetPassword) (*entities.User, error) {
	user, err := u.UsersRepo.FindOneUser(ctx, &entities.UserFilter{UID: &input.UID})
	if err != nil {
		if mongo.ErrNoDocuments == err {
			return nil, errors.RecordNotFoundError{Message: constants.DataNotFound}
		}
		return nil, err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.OldPassword)); err != nil {
		log.WithError(err).Errorln("CompareHashAndPassword")
		return nil, errors.Unauthorized{Message: constants.AuthenticationFailed}
	}
	if input.ConfirmPassword != input.NewPassword {
		return nil, errors.ParameterError{Message: "password not match"}
	}
	userPartialUpdate := &entities.UserPartialUpdate{
		UID:      input.UID,
		Password: pointer.ToString(input.ConfirmPassword),
	}
	return u.UsersRepo.PartialUpdateUser(ctx, userPartialUpdate)
}
