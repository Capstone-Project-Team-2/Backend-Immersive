package volunteers

type VolunteerCore struct {
	ID       string
	EventID  string
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}
type QueryParam struct {
	Page           int
	LimitPerPage   int
	SearchName     string
	ExistOtherPage bool
}

type Login struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

type VolunteerDataInterface interface {
	Login(email, password string) (string, string, error)
	Insert(input VolunteerCore) error
	SelectAll(eventId string, param QueryParam) (int64, []VolunteerCore, error)
	Select(id string) (VolunteerCore, error)
	Update(id string, input VolunteerCore) error
	Delete(id string) error
}

type VolunteerServiceInterface interface {
	Login(email, password string) (string, string, error)
	Create(input VolunteerCore) error
	GetAll(eventId string, param QueryParam) (bool, []VolunteerCore, error)
	GetById(id string) (VolunteerCore, error)
	UpdateById(id string, input VolunteerCore) error
	DeleteById(id string) error
}
