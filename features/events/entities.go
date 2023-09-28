package events

import "time"

type EventCore struct {
	ID               string
	PartnerID        string
	Name             string
	Location         string
	Description      string
	TermCondition    string
	StartDate        time.Time
	EndDate          time.Time
	ValidationStatus string
	ExecutionStatus  string
}
