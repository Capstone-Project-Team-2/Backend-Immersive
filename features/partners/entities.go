package partners

import (
	"mime/multipart"
	"time"
)

type PartnerCore struct {
	ID             string
	Name           string
	StartJoin      time.Time
	Email          string
	Password       string
	PhoneNumber    string
	Address        string
	ProfilePicture string
}

type PartnerDataInterface interface {
	SelectAll() ([]PartnerCore, error)
	Select(id string) (PartnerCore, error)
	Insert(input PartnerCore, file multipart.File) error
	Update(id string, input PartnerCore) error
	Delete(id string) error
	Login(email, password string) (string, string, error)
}

type PartnerServiceInterface interface {
	GetAll() ([]PartnerCore, error)
	Get(id string) (PartnerCore, error)
	Add(input PartnerCore, file multipart.File) error
	Update(id string, input PartnerCore) error
	Delete(id string) error
	Login(email, password string) (string, string, error)
}
