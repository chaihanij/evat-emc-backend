package teams

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/teams/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) SendAssignmentTeamPushDocument(c *gin.Context) {
	request, err := new(dtos.SendAssignmentTeamPushDocumentRequestJSON).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	assignmentTeam, file := request.ToEntity()
	res, err := h.TeamsUseCase.SendAssignmentTeamPushDocument(c.Request.Context(), assignmentTeam, file)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	responseData := new(dtos.FileResponse).Parse(c, res)
	utils.JSONSuccessResponse(c, responseData)
}
