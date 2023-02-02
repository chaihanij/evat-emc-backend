package members

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/members/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) FindAllMember(c *gin.Context) {
	request, err := new(dtos.FindAllMemberRequest).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	totalRecords, members, err := h.MemberUseCase.FindAllMember(c.Request.Context(), request.ToEntity())
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	responseData := new(dtos.FindAllMemberResponseJSON).Parse(members)
	metaData := new(dtos.MetaDataResponse).Parse(request.Page, request.PageSize, totalRecords)
	utils.JSONSuccessCodeWithMetaDataResponse(c, responseData, metaData)

}
