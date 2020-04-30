package registration

type Repository interface {
	Create(user *Registration) (*Registration, error)
	Delete(user *Registration) error
	FindBy(user *Registration, mode string) (*Registration, error)
	FindAll()([]*Registration, error)
	Update(user *Registration) (*Registration, error)
}
