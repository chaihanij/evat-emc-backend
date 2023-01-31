package users

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) Upload(c *gin.Context) {
	image, err := utils.SaveFile(c, "image")
	// request, err := new(dtos.UpdateUserRequestJSON).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	// data := request.ToEntity()
	// res, err := h.UsersUseCase.UpdateUser(c.Request.Context(), &entities.UserFilter{UID: data.UID}, data)
	// if err != nil {
	// 	utils.JSONErrorResponse(c, err)
	// 	return
	// }
	// responseData := new(dtos.UpdateUserResponseJSON).Parse(res)
	utils.JSONSuccessResponse(c, image)
}
