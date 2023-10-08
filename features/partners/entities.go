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
	SelectAll(page, item, search string) ([]PartnerCore, bool, error)
	Select(id string) (PartnerCore, error)
	Insert(input PartnerCore, file multipart.File) error
	Update(id string, input PartnerCore, file multipart.File) error
	Delete(id string) error
	Login(email, password string) (string, string, string, error)
}

type PartnerServiceInterface interface {
	GetAll(page, item, search string) ([]PartnerCore, bool, error)
	Get(id string) (PartnerCore, error)
	Add(input PartnerCore, file multipart.File) error
	Update(id string, input PartnerCore, file multipart.File) error
	Delete(id string) error
	Login(email, password string) (string, string, string, error)
}
