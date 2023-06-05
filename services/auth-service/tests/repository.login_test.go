package tests

import (
	entities "go-microservices/common/models"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func (m *MockRepository) Login(entity *entities.UserEntity) (*entities.UserEntity, int) {
	if entity.Email == "nonexistent@example.com" || entity.Phone == "123456789" {
		return nil, http.StatusNotFound
	}

	if entity.Password == "wrongpassword" {
		return nil, http.StatusUnauthorized
	}

	user := &entities.UserEntity{
		ID:       1,
		Email:    entity.Email,
		Phone:    entity.Phone,
		Password: entity.Password,
	}
	return user, http.StatusOK
}
func TestLogin(t *testing.T) {

	mockRepo := &MockRepository{}

	entity := &entities.UserEntity{
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
	expectedStatus := http.StatusOK

	user, status := mockRepo.Login(entity)

	assert.Equal(t, expectedUser, user)

	assert.Equal(t, expectedStatus, status)

}
func TestLogin_UserNotFound(t *testing.T) {

	mockRepo := &MockRepository{}

	entity := &entities.UserEntity{
		Email: "nonexistent@example.com",
		Phone: "123456789",
	}

	expectedStatus := http.StatusNotFound

	user, status := mockRepo.Login(entity)

	assert.Nil(t, user)
	assert.Equal(t, expectedStatus, status)

}

func TestLogin_IncorrectPassword(t *testing.T) {
	mockRepo := &MockRepository{}

	entity := &entities.UserEntity{
		Email:    "existing@example.com",
		Phone:    "0987654321",
		Password: "wrongpassword",
	}

	expectedStatus := http.StatusUnauthorized

	user, status := mockRepo.Login(entity)

	assert.Nil(t, user)
	assert.Equal(t, status, expectedStatus)
}
