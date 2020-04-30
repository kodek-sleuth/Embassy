package news

import (
	//"fmt"
	_ "github.com/satori/go.uuid"
)

type Service interface {
	Create(news *News) (*News, error)
	Update(news *News) (*News, error)
	Delete(news *News) error
	FindAll()([]*News, error)
	FindById(news *News)(*News, error)
	//CheckUser(news *News) error
}

type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (s service) Create(news *News) (*News, error) {
	result, err := s.repo.Create(news)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s service) Update(news *News) (*News, error) {
	result, err := s.repo.Update(news)
	if err != nil{
		return result, err
	}
	return result, nil
}

func (s service) Delete(news *News) error {
	err := s.repo.Delete(news)
	if err != nil{
		return err
	}
	return nil
}

func (s service) FindAll() ([]*News, error) {
	result, err := s.repo.FindAll()
	if err != nil{
		return result, err
	}
	return result, nil
}

func (s service) FindById(news *News) (*News, error) {
	result, err := s.repo.FindById(news)
	if err != nil{
		return result, err
	}
	return result, nil
}
