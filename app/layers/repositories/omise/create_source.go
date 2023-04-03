package omise

import (
	"github.com/AlekSi/pointer"
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
	log "github.com/sirupsen/logrus"
)

func (r repo) CreateSource(amount int64) (*string, error) {

	source, createSource := &omise.Source{}, &operations.CreateSource{
		Amount:   amount,
		Currency: "thb",
		Type:     "promptpay",
	}
	if err := r.Client.Do(source, createSource); err != nil {
		log.WithError(err).Warnln("CreateSource Do Error")
		return nil, err
	}
	log.WithField("sourceID", source.ID)
	return pointer.ToString(source.ID), nil
}
