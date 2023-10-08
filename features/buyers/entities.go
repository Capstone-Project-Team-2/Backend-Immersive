package buyers

import "mime/multipart"

type BuyerCore struct {
	ID             string
	Name           string `validate:"required"`
	PhoneNumber    string
	Email          string `validate:"required,email"`
	Password       string `validate:"required"`
	Address        string
	ProfilePicture string
}
type Login struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}
type QueryParam struct {
	Page           int
	LimitPerPage   int
	SearchName     string
	ExistOtherPage bool
}
type BuyerDataInterface interface {
	Login(email, password string) (string, string, string, error)
	Insert(input BuyerCore, file multipart.File) error
	SelectAll(param QueryParam) (int64, []BuyerCore, error)
	Select(id string) (BuyerCore, error)
	Update(id string, input BuyerCore, file multipart.File) error
	Delete(id string) error
}

type BuyerServiceInterface interface {
	Login(email, password string) (string, string, string, error)
	Create(input BuyerCore, file multipart.File) error
	GetAll(param QueryParam) (bool, []BuyerCore, error)
	GetById(id string) (BuyerCore, error)
	UpdateById(id string, input BuyerCore, file multipart.File) error
	DeleteById(id string) error
}
