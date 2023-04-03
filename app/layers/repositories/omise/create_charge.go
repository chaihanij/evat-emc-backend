package omise

import (
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/layers/repositories/omise/models"
)

func (r repo) CreateCharge(amount int64, sourceID string, metadata map[string]interface{}) (*entities.OmiseCharge, error) {
	charge, createCharge := &omise.Charge{}, &operations.CreateCharge{
		Amount:   amount,
		Currency: "thb",
		Source:   sourceID,
		Metadata: metadata,
	}
	if err := r.Client.Do(charge, createCharge); err != nil {
		log.WithError(err).Warnln("CreateCharge Do Error")
		return nil, err
	}

	return models.ToChargeEntity(charge)
}
