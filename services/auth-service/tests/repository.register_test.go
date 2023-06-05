package tests

import (
	entities "go-microservices/common/models"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)


func (m *MockRepository) Register(entity *entities.UserEntity) (*entities.UserEntity, int) {

	if entity.Email == "existing@example.com" || entity.Phone == "123456789" {
		return nil, http.StatusConflict
	}

	user := &entities.UserEntity{
		ID:    1,
		Email: entity.Email,
		Phone: entity.Phone,
	}
	return user, http.StatusCreated
}

func TestRegister(t *testing.T) {
	t.Log("Test successful user register")
	mockRepo := &MockRepository{}

	entity := &entities.UserEntity{
		Email: "test@example.com",
		Phone: "987654321",
	}

	expectedUser := &entities.UserEntity{
		ID:    1,
		Email: "test@example.com",
		Phone: "987654321",
	}
	expectedStatus := http.StatusCreated

	user, status := mockRepo.Register(entity)

	assert.Equal(t, expectedUser, user)
	assert.Equal(t, expectedStatus, status)

}

func TestRegister_Fail(t *testing.T) {
	t.Log("user register fail")
	mockRepo := &MockRepository{}

	entity := &entities.UserEntity{
		Email: "existing@example.com",
		Phone: "123456789",
	}

	
	expectedStatus := http.StatusConflict

	user, status := mockRepo.Register(entity)

	assert.Nil(t, user)
	assert.Equal(t,status,expectedStatus)
}

