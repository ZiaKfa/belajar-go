package repository

import (
	"go-unit-test/entity"

	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositoryMock) FindById(id int) (*entity.Category, error) {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	} else {
		category := arguments.Get(0).(entity.Category)
		return &category, nil
	}
}
