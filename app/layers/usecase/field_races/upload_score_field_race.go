package fieldraces

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) UploadScoreFieldRace(ctx context.Context, input *entities.FieldRace) (*entities.FieldRace, error) {
	if input.UUID == "" {
		return nil, fmt.Errorf("fail")
	}
	fieldRace, err := u.TeamFieldRacesRepo.UploadScoreFieldRace(ctx, input)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	log.WithField("value", fieldRace).Debugln("UseCase fieldRace")

	return fieldRace, nil
}
