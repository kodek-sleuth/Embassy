package notice

import (
	//"fmt"
	_ "github.com/satori/go.uuid"
)

type Service interface {
	Create(notice *Notice) (*Notice, error)
	Update(notice *Notice) (*Notice, error)
	Delete(notice *Notice) error
	FindAll()([]*Notice, error)
	FindById(notice *Notice)(*Notice, error)
	//CheckUser(notice *Notice) error
}

type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (s service) Create(notice *Notice) (*Notice, error) {
	result, err := s.repo.Create(notice)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s service) Update(notice *Notice) (*Notice, error) {
	result, err := s.repo.Update(notice)
	if err != nil{
		return result, err
	}
	return result, nil
}

func (s service) Delete(notice *Notice) error {
	err := s.repo.Delete(notice)
	if err != nil{
		return err
	}
	return nil
}

func (s service) FindAll() ([]*Notice, error) {
	result, err := s.repo.FindAll()
	if err != nil{
		return result, err
	}
	return result, nil
}

func (s service) FindById(notice *Notice) (*Notice, error) {
	result, err := s.repo.FindById(notice)
	if err != nil{
		return result, err
	}
	return result, nil
}
