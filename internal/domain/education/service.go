package education

import (
	//"fmt"
	_ "github.com/satori/go.uuid"
)

type Service interface {
	Create(education *Education) (*Education, error)
	Update(education *Education) (*Education, error)
	Delete(education *Education) error
	FindAll()([]*Education, error)
	FindById(education *Education)(*Education, error)
	//CheckUser(education *Education) error
}

type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (s service) Create(education *Education) (*Education, error) {
	result, err := s.repo.Create(education)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s service) Update(education *Education) (*Education, error) {
	result, err := s.repo.Update(education)
	if err != nil{
		return result, err
	}
	return result, nil
}

func (s service) Delete(education *Education) error {
	err := s.repo.Delete(education)
	if err != nil{
		return err
	}
	return nil
}

func (s service) FindAll() ([]*Education, error) {
	result, err := s.repo.FindAll()
	if err != nil{
		return result, err
	}
	return result, nil
}

func (s service) FindById(education *Education) (*Education, error) {
	result, err := s.repo.FindById(education)
	if err != nil{
		return result, err
	}
	return result, nil
}
