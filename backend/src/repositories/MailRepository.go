package repositories

import (
	. "api/src/models"
	"encoding/json"

	"gopkg.in/validator.v2"
)

type MailRepository struct {
	Model Mail
	MailRepositoryInterface
}

type MailRepositoryInterface interface {
	Send(params *json.Decoder) (Mail, error)
}

func (r *MailRepository) Send(params *json.Decoder) (Mail, error) {
	var mail Mail

	err := params.Decode(&mail)

	if err != nil {
		return mail, err
	}

	err = validator.Validate(&mail)

	if err != nil {
		return mail, err
	}

	return r.Model.Send(mail)
}
