package entities

import "time"

type Otp struct {
	Id          string
	Uid         *string
	PhoneNumber *string
	Email       *string
	Otp         string
	RefCode     *string
	ExpireAt    time.Time
}

type OtpFilter struct {
	Id          *string
	Uid         *string
	PhoneNumber *string
	Email       *string
	Otp         *string
	RefCode     *string
}
