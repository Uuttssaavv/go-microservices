package tests

import (
	entities "go-microservices/common/models"

	mocks "go-microservices/services/auth-service/tests/mock"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	t.Log("Test successful user login")
	mockRepo := mocks.NewRepository(t)

	entity := entities.UserEntity{
		Email:    "test@example.com",
		Phone:    "987654321",
		Password: "password",
	}

	expectedUser := &entities.UserEntity{
		ID:       1,
		Email:    "test@example.com",
		Phone:    "987654321",
		Password: "password",
	}
	
	//  stub the Register call
	mockRepo.On("Register", &entity).Return(expectedUser, http.StatusOK)

	expectedStatus := http.StatusOK

	user, status := mockRepo.Register(&entity)

	assert.Equal(t, expectedUser, user)

	assert.Equal(t, expectedStatus, status)

}

func TestRegister_AlreadyExists(t *testing.T) {
	
	t.Log("Test when incorrect password")
	mockRepo := mocks.NewRepository(t)

	entity := entities.UserEntity{
		Email:    "test@example.com",
		Phone:    "987654321",
		Password: "password",
	}

	//  stub the Register call
	mockRepo.On("Register", &entity).Return(nil, http.StatusConflict)

	expectedStatus := http.StatusConflict

	user, status := mockRepo.Register(&entity)

	assert.Nil(t, user)

	assert.Equal(t, expectedStatus, status)

}
func TestLogin_UnableToCreate(t *testing.T) {
	
	t.Log("Test when user does not exist")
	mockRepo := mocks.NewRepository(t)

	entity := entities.UserEntity{
		Email:    "test@example.com",
		Phone:    "987654321",
		Password: "password",
	}

	//  stub the Register call
	mockRepo.On("Register", &entity).Return(nil, http.StatusExpectationFailed)

	expectedStatus := http.StatusExpectationFailed

	user, status := mockRepo.Register(&entity)

	assert.Nil(t, user)

	assert.Equal(t, expectedStatus, status)

}
