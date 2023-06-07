package tests

import (
	entities "go-microservices/common/models"

	mocks "go-microservices/services/auth-service/tests/mock"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {

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
	//  stub the login
	mockRepo.On("Login", &entity).Return(expectedUser, http.StatusOK)

	expectedStatus := http.StatusOK

	user, status := mockRepo.Login(&entity)

	assert.Equal(t, expectedUser, user)

	assert.Equal(t, expectedStatus, status)

}

func TestLogin_Unauthorized(t *testing.T) {
	
	t.Log("Test when incorrect password")
	mockRepo := mocks.NewRepository(t)

	entity := entities.UserEntity{
		Email:    "test@example.com",
		Phone:    "987654321",
		Password: "password",
	}

	//  stub the login
	mockRepo.On("Login", &entity).Return(nil, http.StatusUnauthorized)

	expectedStatus := http.StatusUnauthorized

	user, status := mockRepo.Login(&entity)

	assert.Nil(t, user)

	assert.Equal(t, expectedStatus, status)

}
func TestLogin_UserDoesnotExist(t *testing.T) {
	
	t.Log("Test when user does not exist")
	mockRepo := mocks.NewRepository(t)

	entity := entities.UserEntity{
		Email:    "test@example.com",
		Phone:    "987654321",
		Password: "password",
	}

	//  stub the login
	mockRepo.On("Login", &entity).Return(nil, http.StatusNotFound)

	expectedStatus := http.StatusNotFound

	user, status := mockRepo.Login(&entity)

	assert.Nil(t, user)

	assert.Equal(t, expectedStatus, status)

}
