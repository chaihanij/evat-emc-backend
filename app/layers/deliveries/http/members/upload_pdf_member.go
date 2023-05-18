package members

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/members/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) UploadPDFMember(c *gin.Context) {
	request, err := new(dtos.UpdatePDFRequest).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}

	// UploadPDFMember(ctx context.Context, memberUUID string, file *entities.File) (*entities.File, error)
	_, err = h.MemberUseCase.UploadPDFMember(c.Request.Context(), request.MemberUUID, request.ToEntity())
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	// responseData := new(dtos.UpdatePDFRespones).Parse(c, res)
	// utils.JSONSuccessResponse(c, responseData)
}
