package members

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/members/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) CreateCertificate(c *gin.Context) {
	var w http.ResponseWriter = c.Writer

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
	responseData, err := new(dtos.ResponseCertificate).Parse(c, res)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	//res
	w.Header().Set("Content-Disposition", `attachment; filename=`+res.FirstName+`.pdf`)
	w.Header().Set("Content-Type", "application/pdf")
	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
}
