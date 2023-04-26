package teams

import (
	"context"

	"github.com/AlekSi/pointer"
	"github.com/omise/omise-go"
	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) WebHooks(ctx context.Context, input *entities.OmiseEvent) error {

	if input.Key != "charge.complete" {
		log.WithField("value", input.Key).Infoln("charge.complete")
		return nil
	}
	var charge omise.Charge = input.Data
	if charge.Status != omise.ChargeSuccessful {
		log.WithField("value", charge.Status).Infoln("WebHooks charge status")
		return nil
	}
	teamUUIDData, ok := charge.Metadata["team_uuid"]
	if !ok {
		log.Infoln("WebHooks charge team_uuid not found")
		return nil
	}
	teamUUID, ok := teamUUIDData.(string)
	if !ok {
		log.Infoln("WebHooks charge team_uuid not found")
		return nil
	}
	emailData, ok := charge.Metadata["email"]
	if !ok {
		log.Infoln("WebHooks charge email not found")
		return nil
	}
	email, ok := emailData.(string)
	if !ok {
		log.Infoln("WebHooks charge email not found")
		return nil
	}
	team, err := u.TeamsRepo.PartialUpdateTeam(ctx, &entities.TeamPartialUpdate{UUID: teamUUID, IsPaid: pointer.ToBool(true)})
	if err != nil {
		return err
	}

	activateCode, err := u.UsersRepo.FindOneUser(ctx, &entities.UserFilter{Email: &email})

	err = u.TeamsRepo.SendEmailRegister(email, activateCode.ActivateCode)
	if err != nil {
		return nil
	}

	log.WithFields(log.Fields{
		"emai":   email,
		"charge": charge,
		"team":   team,
	}).Infoln("WebHooks")
	return err
}
