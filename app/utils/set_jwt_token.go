package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/errors"
)

func GetJwtToken(id *string, uuid *string, jwtMapName string, jwtMapValue string) (*string, error) {
	if id == nil || uuid == nil {
		return nil, errors.InternalError{Message: constants.UserIdMissing}
	}
	stringId := *id
	idEnc, err := EncryptAES(stringId, env.EncryptKey)
	if err != nil {
		return nil, errors.InternalError{Message: err.Error()}
	}
	stringUUID := *uuid
	uuidEnc, err := EncryptAES(stringUUID, env.EncryptKey)
	if err != nil {
		return nil, errors.InternalError{Message: err.Error()}
	}
	// Default 0 Jwt Token Life Is An Never Expired
	jwtTokenLife, err := strconv.Atoi(env.JwtTokenLife)
	if err != nil {
		jwtTokenLife = 0
	}
	var jwtMapClaims jwt.MapClaims

	if jwtTokenLife > 0 {
		jwtMapClaims = jwt.MapClaims{
			"iat":      time.Now().Unix(),
			"exp":      time.Now().Add(time.Minute * time.Duration(jwtTokenLife)).Unix(),
			"guid":     fmt.Sprintf("%x", idEnc),
			"uuid":     fmt.Sprintf("%x", uuidEnc),
			jwtMapName: jwtMapValue,
		}
	} else {
		jwtMapClaims = jwt.MapClaims{
			"iat":      time.Now().Unix(),
			"guid":     fmt.Sprintf("%x", idEnc),
			"uuid":     fmt.Sprintf("%x", uuidEnc),
			jwtMapName: jwtMapValue,
		}
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodRS512, jwtMapClaims)
	privateKey, err := GetPrivateKey()
	if err != nil {
		return nil, errors.InternalError{Message: err.Error()}
	}
	jwtTokenString, err := jwtToken.SignedString(privateKey)
	if err != nil {
		return nil, errors.InternalError{Message: err.Error()}
	}
	return &jwtTokenString, nil
}
