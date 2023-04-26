package certificate

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/certificate/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func CreateCertificate(c *gin.Context) {
	logrus.Debugln("call api")
	request, err := new(dtos.CreateCertificateRequestJSON).Parse(c)

	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	logrus.Debugln("request :", request)
}
