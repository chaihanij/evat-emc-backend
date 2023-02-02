package members

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/members/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) UpdateMemberPullDocument(c *gin.Context) {
	request, err := new(dtos.UpdateMemberPullDocumentRequest).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	memberUUID, documentUUID := request.ToEntity()
	err = h.MemberUseCase.UpdateMemberPullDocument(c.Request.Context(), memberUUID, documentUUID)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	utils.JSONSuccessResponse(c, nil)
}
