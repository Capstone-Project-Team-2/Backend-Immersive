package volunteers

type VolunteerCore struct {
	ID       string
	EventID  string
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

type VolunteerDataInterface interface {
	Login(email, password string) (VolunteerCore, string, error)
	Insert(input VolunteerCore) error
	SelectAll() ([]VolunteerCore, error)
	Select(id string) (VolunteerCore, error)
	Update(id string, input VolunteerCore) error
	Delete(id string) error
}

type VolunteerServiceInterface interface {
	Login(email, password string) (VolunteerCore, string, error)
	Create(input VolunteerCore) error
	GetAll() ([]VolunteerCore, error)
	GetById(id string) (VolunteerCore, error)
	UpdateById(id string, input VolunteerCore) error
	DeleteById(id string) error
}
