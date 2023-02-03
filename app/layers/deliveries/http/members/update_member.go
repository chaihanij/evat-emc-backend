package members

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/members/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) UpdateMember(c *gin.Context) {
	request, err := new(dtos.UpdateMemberRequest).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}

	res, err := h.MemberUseCase.UpdateMember(c.Request.Context(), request.ToEntity())
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	responseData := new(dtos.UpdateMemberResponseJSON).Parse(c, res)
	utils.JSONSuccessResponse(c, responseData)
}
