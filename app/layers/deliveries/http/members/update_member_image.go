package members

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/members/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) UpdateMemberImage(c *gin.Context) {
	request, err := new(dtos.UpdateMemberImageRequest).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}

	res, err := h.MemberUseCase.UpdateMemberImage(c.Request.Context(), request.MemberUUID, request.ToEntity())
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	responseData := new(dtos.FileResponse).Parse(res)
	utils.JSONSuccessResponse(c, responseData)
}
