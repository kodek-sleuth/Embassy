package tourism

import (
	//"fmt"
	_ "github.com/satori/go.uuid"
)

type Service interface {
	Create(tourism *Tourism) (*Tourism, error)
	Update(tourism *Tourism) (*Tourism, error)
	Delete(tourism *Tourism) error
	FindAll()([]*Tourism, error)
	FindById(tourism *Tourism)(*Tourism, error)
	//CheckUser(tourism *Tourism) error
}

type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (s service) Create(tourism *Tourism) (*Tourism, error) {
	result, err := s.repo.Create(tourism)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s service) Update(tourism *Tourism) (*Tourism, error) {
	result, err := s.repo.Update(tourism)
	if err != nil{
		return result, err
	}
	return result, nil
}

func (s service) Delete(tourism *Tourism) error {
	err := s.repo.Delete(tourism)
	if err != nil{
		return err
	}
	return nil
}

func (s service) FindAll() ([]*Tourism, error) {
	result, err := s.repo.FindAll()
	if err != nil{
		return result, err
	}
	return result, nil
}

func (s service) FindById(tourism *Tourism) (*Tourism, error) {
	result, err := s.repo.FindById(tourism)
	if err != nil{
		return result, err
	}
	return result, nil
}
