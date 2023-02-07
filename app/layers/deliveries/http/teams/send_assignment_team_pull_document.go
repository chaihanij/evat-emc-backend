package teams

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/teams/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) SendAssignmentTeamPullhDocument(c *gin.Context) {
	request, err := new(dtos.SendAssignmentTeamPullDocumentRequestJSON).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	assignmentTeam, documentID := request.ToEntity()
	err = h.TeamsUseCase.SendAssignmentTeamPullDocument(c.Request.Context(), assignmentTeam, documentID)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	utils.JSONSuccessResponse(c, nil)
}
