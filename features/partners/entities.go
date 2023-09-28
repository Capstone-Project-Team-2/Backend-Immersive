package partners

import "time"

type PartnerCore struct {
	ID          uint
	Name        string
	StartJoin   time.Time
	Email       string
	Password    string
	PhoneNumber string
	Address     string
}
