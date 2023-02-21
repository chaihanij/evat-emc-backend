package users

import (
	"context"
	"time"

	"github.com/AlekSi/pointer"
	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
	"gitlab.com/chaihanij/evat/app/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func (u useCase) Login(ctx context.Context, input *entities.Login) (*entities.User, error) {

	filter := entities.UserFilter{Email: pointer.ToString(input.Email)}
	user, err := u.UsersRepo.FindOneUser(ctx, &filter)
	if err != nil {
		if mongo.ErrNoDocuments == err {
			return nil, errors.RecordNotFoundError{Message: constants.DataNotFound}
		}
		return nil, err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		log.WithError(err).Errorln("CompareHashAndPassword")
		return nil, errors.Unauthorized{Message: constants.AuthenticationFailed}
	}

	accessToken, err := utils.GetJwtToken(&user.UID, &user.UID, "email", user.Email)
	if err != nil {
		log.WithError(err).Errorln("GetJwtToken")
		return nil, err
	}
	partialUpdateUser := &entities.UserPartialUpdate{
		UID:         user.UID,
		AccessToken: accessToken,
		LastLogin:   pointer.ToTime(time.Now()),
	}
	return u.UsersRepo.PartialUpdateUser(ctx, partialUpdateUser)
}
