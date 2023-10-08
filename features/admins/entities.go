package admins

type AdminCore struct {
	ID          string
	Name        string
	Email       string
	Password    string
	PhoneNumber string
	Address     string
	Role        string
}

type AdminDataInterface interface {
	Register(AdminCore) error
	Login(email, password string) (string, string, string, error)
	Get(admin_id string) (AdminCore, error)
	GetAll() ([]AdminCore, error)
	Update(admin_id string, input AdminCore) error
	Delete(admin_id string) error
}

type AdminServiceInterface interface {
	Register(AdminCore) error
	Login(email, password string) (string, string, string, error)
	Get(admin_id string) (AdminCore, error)
	GetAll() ([]AdminCore, error)
	Update(admin_id string, input AdminCore) error
	Delete(admin_id string) error
}
