package omise

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) WebHooks(c *gin.Context) {
	var event entities.OmiseEvent
	err := c.ShouldBind(&event)
	if err != nil {
		utils.JSONErrorResponse(c, errors.ParameterError{Message: err.Error()})
	}
	log.WithField("value", event).Debugln("WebHooks")
	err = h.TeamsUseCase.WebHooks(c.Request.Context(), &event)
	utils.JSONSuccessResponse(c, err)
}
