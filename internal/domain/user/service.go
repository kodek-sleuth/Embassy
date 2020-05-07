package user

import "Embassy/internal/helpers"

type Service interface {
	Create(user *User) (*User, error)
	FindAll() ([]*User, error)
	FindBy(user *User, mode string) (*User, error)
	Login(user *User, password string) (*User, error)
	Update(user *User) (*User, error)
	Delete(user *User) error
	GetAll() (map[string]interface{}, error)
}

type service struct {
	repo Repository
}

// Implement UserHandler Interface
func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (u *service) Create(user *User) (*User, error) {
	hash, err := helpers.GenerateHash([]byte(user.Password))
	if err != nil{
		return nil, err
	}

	user.Password = hash

	result, err := u.repo.Create(user)
	if err != nil{
		return nil, err
	}
	return result, nil
}

func(u *service) Login(user *User, password string) (*User, error) {
	entity, err := u.repo.FindBy(user, "email")
	if err != nil{
		return nil, err
	}

	err = helpers.CompareHash(entity.Password, password)
	if err != nil{
		return entity, err
	}

	return entity, nil
}

func (u *service) Update(user *User) (*User, error) {
	result, err := u.repo.Update(user)
	if err != nil{
		return nil, err
	}
	return result, nil
}

func (u *service) FindAll() ([]*User, error) {
	result, err := u.repo.FindAll()
	if err != nil{
		return nil, err
	}
	return result, nil
}

func (u *service) FindBy(user *User, mode string) (*User, error) {
	entity, err := u.repo.FindBy(user, mode)
	if err != nil{
		return nil, err
	}

	return entity, nil
}

func (u *service) Delete(user *User) error {
	err := u.repo.Delete(user)
	if err != nil{
		return err
	}
	return nil
}

func (u *service) GetAll() (map[string]interface{}, error) {
	entity, err := u.repo.GetAll()
	if err != nil{
		return nil, err
	}

	return entity, nil
}