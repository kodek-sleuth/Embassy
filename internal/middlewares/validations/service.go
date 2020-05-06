package validations

import (
	//"fmt"
	_ "github.com/satori/go.uuid"
)

type Service interface {
	Create(chat *Chat) (*Chat, error)
	Update(chat *Chat) (*Chat, error)
	Delete(chat *Chat) error
	FindAll()([]*Chat, error)
	FindById(chat *Chat)(*Chat, error)
	//CheckUser(chat *Chat) error
}

type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (s service) Create(chat *Chat) (*Chat, error) {
	result, err := s.repo.Create(chat)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s service) Update(chat *Chat) (*Chat, error) {
	result, err := s.repo.Update(chat)
	if err != nil{
		return result, err
	}
	return result, nil
}

func (s service) Delete(chat *Chat) error {
	err := s.repo.Delete(chat)
	if err != nil{
		return err
	}
	return nil
}

func (s service) FindAll() ([]*Chat, error) {
	result, err := s.repo.FindAll()
	if err != nil{
		return result, err
	}
	return result, nil
}

func (s service) FindById(chat *Chat) (*Chat, error) {
	result, err := s.repo.FindById(chat)
	if err != nil{
		return result, err
	}
	return result, nil
}
