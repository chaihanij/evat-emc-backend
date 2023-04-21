package assignments

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/assignments/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) UploadFileScore(c *gin.Context) {

	request, err := new(dtos.FileRequestJSON).Pares(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}

	// fmt.Println("req :", request)
	// filescore, file := request.ToEntity()

	// reqponseData := new(dtos.FileRequestJSON).Pares(c, request)
	// utils.JSONSuccessResponse(c, request)

	_, err = h.AssignmentsUseCase.UploadFileScore(c.Request.Context(), request.AssignmentUUID, request.ToEntity())
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}

	// responseData := new(dtos.FileResponse).Parse(c, res)
	// utils.JSONSuccessResponse(c, responseData)

}
