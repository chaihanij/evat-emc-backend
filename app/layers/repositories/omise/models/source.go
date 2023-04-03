package models

import (
	"github.com/omise/omise-go"

	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/entities"
)

func ToSourceEntity(s *omise.Source) (*entities.OmiseSource, error) {
	var omiseSource entities.OmiseSource
	err := copier.Copy(&omiseSource, s)
	return &omiseSource, err
}
