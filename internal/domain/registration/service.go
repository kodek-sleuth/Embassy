package registration

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Service interface {
	Create(user *Registration) (*Registration, error)
	FindBy(user *Registration, mode string) (*Registration, error)
	Update(user *Registration) (*Registration, error)
}

type service struct {
	repo Repository
}

// Implement RegistrationHandler Interface
func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (u *service) Create(user *Registration) (*Registration, error) {
	rand.Seed(time.Now().UnixNano())
	max := 1000000
	min := 10000
	user.Code = fmt.Sprintf("CON-AUS-%s-%v", strings.ToUpper(user.City[:3]), max+rand.Intn(max-min))

	result, err := u.repo.Create(user)
	if err != nil{
		return nil, err
	}
	return result, nil
}


func (u *service) Update(user *Registration) (*Registration, error) {
	result, err := u.repo.Update(user)
	if err != nil{
		return nil, err
	}
	return result, nil
}

func (u *service) FindBy(user *Registration, mode string) (*Registration, error) {
	entity, err := u.repo.FindBy(user, mode)
	if err != nil{
		return nil, err
	}

	return entity, nil
}