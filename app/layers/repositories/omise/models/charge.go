package models

import (
	"github.com/omise/omise-go"

	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/entities"
)

func ToChargeEntity(s *omise.Charge) (*entities.OmiseCharge, error) {
	var omiseCharge entities.OmiseCharge
	err := copier.Copy(&omiseCharge, s)
	return &omiseCharge, err
}
