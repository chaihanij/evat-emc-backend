package members

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/members/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) CreateCertificate(c *gin.Context) {
	req, err := new(dtos.RequestCertificate).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}

	res, err := h.MemberUseCase.CreateCertificate(c.Request.Context(), *req.ToEntity().UUID)
	responseData := new(dtos.ResponseCertificate).Parse(c, res)

	utils.JSONSuccessResponse(c, responseData)

}
