package score

// import (
// 	"github.com/gin-gonic/gin"
// 	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/score/dtos"
// 	"gitlab.com/chaihanij/evat/app/utils"
// )

// func (h *Handler) CreateScore(c *gin.Context) {

// 	request, err := new(dtos.CreateScoreRequestJSON).Parse(c)
// 	if err != nil {
// 		utils.JSONErrorResponse(c, err)
// 	}

// 	res, err := h.ScoreUseCase.CreateScore(c.Request.Context(), request.ToEntity())
// 	if err != nil {
// 		utils.JSONErrorResponse(c, err)
// 		return
// 	}
// 	responseData := new(dtos.CreateScoreResponseJSON).Parse(res)
// 	utils.JSONSuccessResponse(c, responseData)

// }
