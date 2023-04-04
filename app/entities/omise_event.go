package entities

import (
	"time"

	"github.com/omise/omise-go"
)

type OmiseEvent struct {
	Object   string       `json:"object"`
	ID       string       `json:"id"`
	Live     bool         `json:"livemode"`
	Location *string      `json:"location"`
	Created  time.Time    `json:"created"`
	Key      string       `json:"key"`
	Data     omise.Charge `json:"data"`
}
