package user

type Repository interface {
	Create(user *User) (*User, error)
	Delete(user *User) error
	FindBy(user *User, mode string) (*User, error)
	FindAll()([]*User, error)
	Update(user *User) (*User, error)
}
