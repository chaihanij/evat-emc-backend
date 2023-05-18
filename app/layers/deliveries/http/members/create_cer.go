package members

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/members/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) CreateC(c *gin.Context) {
	// var w http.ResponseWriter = c.Writer

	req, err := new(dtos.RequestCertificate).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}

	res, err := h.MemberUseCase.CreateCertificate(c.Request.Context(), *req.ToEntity().UUID)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	responseData := new(dtos.ResponseCertificateNew).Parse(c, res)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}

	utils.JSONSuccessResponse(c, responseData)

}
