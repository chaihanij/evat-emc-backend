package healthcheck

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (handler *handler) Health(c *gin.Context) {
	utils.JSONSuccessResponse(c, nil)
}
