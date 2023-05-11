package dtos

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

type MemberCheckInRequestJSON struct {
	Member_uuid    string `uri:"member_uuid"`
	Is_checkin     bool   `json:"is_checkin"`
	Check_national bool   `json:"is_national"`
	Is_check_data  bool   `json:"is_data"`
	Is_Check_image bool   `json:"is_image"`
}

func (req *MemberCheckInRequestJSON) Parse(c *gin.Context) (*MemberCheckInRequestJSON, error) {
	if err := c.ShouldBindJSON(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	if err := c.ShouldBindUri(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	return req, nil
}

func (req *MemberCheckInRequestJSON) ToEntity() *entities.MemberCheckIn {

	memberChectIn := &entities.MemberCheckIn{
		Member_uuid:    &req.Member_uuid,
		Is_checkin:     &req.Is_checkin,
		Check_national: &req.Check_national,
		Is_check_data:  &req.Is_check_data,
		Is_Check_image: &req.Is_Check_image,
	}

	return memberChectIn
}
