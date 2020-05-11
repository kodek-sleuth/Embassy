package pages

import (
	//"fmt"
	"Embassy/internal/helpers"
	_ "github.com/satori/go.uuid"
)

type Service interface {
	Create(pages *Pages) (*Pages, error)
	Update(pages *Pages) (*Pages, error)
	Delete(pages *Pages) error
	FindAll()([]*Pages, error)
	FindById(pages *Pages)(*Pages, error)
	//CheckUser(pages *Pages) error
}

type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (s service) Create(pages *Pages) (*Pages, error) {
	str, err := helpers.ParseNodes(pages.Body)
	if err != nil{
		return nil, err
	}

	pages.Body = str
	result, err := s.repo.Create(pages)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s service) Update(pages *Pages) (*Pages, error) {
	str, err := helpers.ParseNodes(pages.Body)
	if err != nil{
		return nil, err
	}

	pages.Body = str
	result, err := s.repo.Update(pages)
	if err != nil{
		return result, err
	}
	return result, nil
}

func (s service) Delete(pages *Pages) error {
	err := s.repo.Delete(pages)
	if err != nil{
		return err
	}
	return nil
}

func (s service) FindAll() ([]*Pages, error) {
	result, err := s.repo.FindAll()
	if err != nil{
		return result, err
	}
	return result, nil
}

func (s service) FindById(pages *Pages) (*Pages, error) {
	result, err := s.repo.FindById(pages)
	if err != nil{
		return result, err
	}
	return result, nil
}
