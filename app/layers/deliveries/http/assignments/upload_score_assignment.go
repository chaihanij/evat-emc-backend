package assignments

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/assignments/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (H *Handler) UploadScoreAssignment(c *gin.Context) {
	request, err := new(dtos.UploadScoreAssingment).Parse(c)
	if err != nil {
		logrus.Debugln("uploaddata")
		utils.JSONErrorResponse(c, err)
		return
	}
	fmt.Println("request :", request)
}
