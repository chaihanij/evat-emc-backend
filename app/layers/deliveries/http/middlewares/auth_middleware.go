package middlewares

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/errors"
	"gitlab.com/chaihanij/evat/app/layers/usecase/users"
	"gitlab.com/chaihanij/evat/app/utils"
)

type AuthMiddleware interface {
	Authentication(c *gin.Context)
}

type authMiddleware struct {
	UsersUseCase users.UseCase
}

func InitAuthMiddleware(usersUseCase users.UseCase) AuthMiddleware {
	return &authMiddleware{
		UsersUseCase: usersUseCase,
	}
}

func (middleware *authMiddleware) authentication(c *gin.Context, jwtAccessToken string) {

	publicKey, err := utils.GetPublicKey()
	if err != nil {
		err = errors.Unauthorized{Message: err.Error()}
		utils.JSONErrorResponse(c, err)
		return
	}

	claims := jwt.MapClaims{}
	jwtToken, err := jwt.ParseWithClaims(jwtAccessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil {
		err = errors.Unauthorized{Message: err.Error()}
		utils.JSONErrorResponse(c, err)
		return
	}

	if jwtToken == nil {
		err = errors.Unauthorized{Message: "invalid JWT token"}
		utils.JSONErrorResponse(c, err)
		return
	}

	jwtIssueAt := fmt.Sprintf("%v", claims["iat"])
	jwtIssueAtValue, err := strconv.ParseFloat(jwtIssueAt, 64)
	if err != nil {
		err = errors.Unauthorized{Message: fmt.Sprintf("JWT iat '%v' not found", jwtIssueAt)}
		utils.JSONErrorResponse(c, err)
		return
	}

	jwtGuid := fmt.Sprintf("%v", claims["guid"])
	if claims["guid"] == nil || jwtGuid == "" {
		err = errors.Unauthorized{Message: fmt.Sprintf("JWT guid '%v' not found", jwtGuid)}
		utils.JSONErrorResponse(c, err)
		return
	}

	guid, err := utils.DecryptAES(jwtGuid, env.EncryptKey)
	if err != nil {
		err = errors.Unauthorized{Message: fmt.Sprintf("decrypt AES guid: %s", err.Error())}
		utils.JSONErrorResponse(c, err)
		return
	}

	jwtUuid := fmt.Sprintf("%v", claims["uuid"])
	if claims["uuid"] == nil || jwtGuid == "" {
		err = errors.Unauthorized{Message: fmt.Sprintf("JWT uuid '%v' not found", jwtGuid)}
		utils.JSONErrorResponse(c, err)
		return
	}

	uuid, err := utils.DecryptAES(jwtUuid, env.EncryptKey)
	if err != nil {
		err = errors.Unauthorized{Message: fmt.Sprintf("decrypt AES uuid: %s", err.Error())}
		utils.JSONErrorResponse(c, err)
		return
	}

	jwtData := entities.JwtData{
		ID:      *guid,
		UID:     *uuid,
		IssueAt: jwtIssueAtValue,
	}

	c.Set(constants.JWTDataKey, jwtData)

	c.Header(constants.Authorization, jwtAccessToken)

	c.Next()
}

func (middleware *authMiddleware) Authentication(c *gin.Context) {

	s := c.GetHeader(constants.Authorization)
	jwtAccessToken := strings.TrimPrefix(s, "Bearer ")
	if jwtAccessToken == "" {
		err := errors.Unauthorized{Message: constants.MissingAuthorization}
		utils.JSONErrorResponse(c, err)
		return
	}

	middleware.authentication(c, jwtAccessToken)
}
