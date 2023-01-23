package auth

import (
	"context"
	"log"

	"github.com/AlekSi/pointer"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func (u useCase) Auth(ctx context.Context, input *entities.Auth) (*string, error) {

	filter := entities.UserFilter{Email: pointer.ToString(input.Email)}
	user, err := u.UsersRepo.FindOneUser(ctx, &filter)
	if err != nil {
		if mongo.ErrNoDocuments == err {
			return nil, errors.RecordNotFoundError{Message: constants.DataNotFound}
		}
		return nil, err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		log.Print("UC LoginUser step CompareHashAndPassword error :", err)
		return nil, errors.Unauthorized{Message: constants.AuthenticationFailed}
	}

	accessToken, err := utils.GetJwtToken(res.Id, res.Uuid, "username", *input.Username)
	if err != nil {
		log.Print("UC LoginByUsername step GetJwtToken error :", err)
		return nil, err
	}

	return nil, nil
	// if err = bcrypt.CompareHashAndPassword([]byte(*res.Password), []byte(*input.Password)); err != nil {
	// 	log.Print("UC LoginByUsername step CompareHashAndPassword error :", err)
	// 	return nil, errors.Unauthorized{Message: constants.AuthenticationFailed}
	// }
	// accessToken, err := utils.GetJwtToken(res.Id, res.Uuid, "username", *input.Username)
	// if err != nil {
	// 	log.Print("UC LoginByUsername step GetJwtToken error :", err)
	// 	return nil, err
	// }
	// token := entities.Users{Id: res.Id, AccessToken: accessToken}
	// _, err = u.UsersRepo.UpdateUser(ctx, &token)
	// if err != nil {
	// 	log.Print("UC LoginByUsername step UpdateUser error :", err)
	// 	return nil, err
	// }
	// utils.FromUserIdToGinContext(ctx, *res.Id)
	// return accessToken, nil
}
