package files

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/files/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) FindOneFile(c *gin.Context) {
	request, err := new(dtos.FindOneFileRequest).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	res, err := h.FilesUseCase.FindOneFile(c.Request.Context(), request.ToEntity())

	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	new(dtos.FindOneFileResponse).Response(c, res)

}
