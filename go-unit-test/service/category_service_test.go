package service

import (
	"errors"
	"go-unit-test/entity"
	"go-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	// Program the behavior of the "mock" object
	categoryRepository.Mock.On("FindById", 1).Return(nil, errors.New("NOT FOUND"))

	// Perform the test
	result, err := categoryService.Get(1)
	assert.Nil(t, result)
	assert.NotNil(t, err)

	// Make sure the expectations were met
	categoryRepository.Mock.AssertExpectations(t)
}

func TestCategoryService_GetFound(t *testing.T) {
	category := entity.Category{
		ID:   2,
		Name: "Pikachu",
	}
	categoryRepository.Mock.On("FindById", 2).Return(category, nil)

	// Perform the test
	result, err := categoryService.Get(2)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.ID, result.ID)
	assert.Equal(t, category.Name, result.Name)
}
