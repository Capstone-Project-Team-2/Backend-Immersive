package data

import (
	partnerModel "capstone-tickets/features/partners/data"
	"time"

	"gorm.io/gorm"
)

type Event struct {
	ID               string               `gorm:"column:id;type:varchar(191);primaryKey"`
	PartnerID        string               `gorm:"column:partner_id;type:varchar(191)"`
	Name             string               `gorm:"column:name;not null"`
	Location         string               `gorm:"column:location"`
	Description      string               `gorm:"column:description"`
	TermCondition    string               `gorm:"column:term_condition"`
	StartDate        time.Time            `gorm:"column:start_date"`
	EndDate          time.Time            `gorm:"column:end_date"`
	ValidationStatus string               `gorm:"column:validation_status;type:enum('Pending','Valid','Not Valid');default:Pending"`
	ExecutionStatus  string               `gorm:"column:execution_status;type:enum('On Going','Done');default:'On Going'"`
	Partner          partnerModel.Partner `gorm:"foreignKey:PartnerID"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
}
