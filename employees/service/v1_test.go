package service_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"algogrit.com/httpex/employees/repository"
	"algogrit.com/httpex/employees/service"
	"algogrit.com/httpex/entities"
)

func TestIndex(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockRepo := repository.NewMockEmployeeRepository(ctrl)

	sut := service.NewV1(mockRepo)

	expectedEmployees := []entities.Employee{
		{1, "Gaurav", "LnD", 1001},
	}

	mockRepo.EXPECT().FindAll().Return(expectedEmployees, nil)
	mockRepo.EXPECT().FindAll().Return(nil, errors.New("some error"))

	employees, err := sut.Index()

	assert.Nil(t, err)

	assert.NotEmpty(t, employees)
	assert.Equal(t, expectedEmployees, employees)

	employees, err = sut.Index()

	assert.NotNil(t, err)

	assert.Empty(t, employees)
}
