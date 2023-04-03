package omise

import (
	opn "github.com/omise/omise-go"
	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
)

type repo struct {
	Client *opn.Client
}

type Repo interface {
	CreateCharge(amount int64, sourceID string, metadata map[string]interface{}) (*entities.OmiseCharge, error)
	CreateSource(amount int64) (*string, error)
}

func InitRepo(public string, private string) Repo {
	client, err := opn.NewClient(public, private)
	if err != nil {
		log.WithError(err).Warnln("InitRepo omise errr")
	}
	return &repo{
		Client: client,
	}
}
