package dtos

import (
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/entities"
)

type CreateVisitRequestJSON struct {
	IP        string    `json:"ip" bson:"ip"`
	Create_at time.Time `json:"create_at" bson:"create_at"`
}

func (req *CreateVisitRequestJSON) Parse(c *gin.Context) (*CreateVisitRequestJSON, error) {

	req.IP = c.ClientIP()
	req.Create_at = time.Now()

	return req, nil

}

func (req *CreateVisitRequestJSON) ToEntity() *entities.UpdateVisit {
	return &entities.UpdateVisit{
		IP:        req.IP,
		Create_at: req.Create_at,
	}
}

type CreateVisitResponseJSON struct {
	IP        string    `json:"ip" bson:"ip"`
	Create_at time.Time `json:"create_at" bson:"create_at"`
}

func (m *CreateVisitResponseJSON) Parse(data *entities.UpdateVisit) *CreateVisitResponseJSON {
	visit := &CreateVisitResponseJSON{
		IP:        data.IP,
		Create_at: data.Create_at,
	}
	return visit
}
