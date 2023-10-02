package buyers

import (
	"mime/multipart"
)

type BuyerCore struct {
	ID             string
	Name           string `validate:"required"`
	PhoneNumber    string
	Email          string `validate:"required,email"`
	Password       string `validate:"required"`
	Address        string
	ProfilePicture string
}

type BuyerDataInterface interface {
	Login(email, password string) (BuyerCore, string, error)
	Insert(input BuyerCore, file multipart.File) error
	SelectAll() ([]BuyerCore, error)
	Select(id string) (BuyerCore, error)
	Update(input BuyerCore) error
	Delete(id string) error
}

type BuyerServiceInterface interface {
	Login(email, password string) (BuyerCore, string, error)
	Create(input BuyerCore, file multipart.File) error
	GetAll() ([]BuyerCore, error)
	GetById(id string) (BuyerCore, error)
	UpdateById(id string, input BuyerCore) error
	DeleteById(id string) error
}
