package admins

type AdminCore struct {
	ID       string
	Name     string
	Email    string
	Password string
	Role     string
}

type AdminDataInterface interface {
	Register(AdminCore) error
}

type AdminServiceInterface interface {
	Register(AdminCore) error
}
