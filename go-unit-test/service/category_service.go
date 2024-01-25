package service

import (
	"errors"
	"go-unit-test/entity"
	"go-unit-test/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id int) (*entity.Category, error) {
	category, err := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New(err.Error())
	} else {
		return category, nil
	}
}
