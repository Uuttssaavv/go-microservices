package tests

import (
	"bytes"
	entities "go-microservices/common/models"
	"go-microservices/services/auth-service/handlers"
	mocks "go-microservices/services/auth-service/tests/mock"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLoginHandler_Success(t *testing.T) {
	//
	gin.SetMode(gin.TestMode)

	testContext, _ := gin.CreateTestContext(httptest.NewRecorder())

	//  create the mock service
	mockService := mocks.NewService(t)

	// initialize the handler
	handler := handlers.NewAuthHandler(mockService)

	//  stub the Login function
	mockService.On("Login", mock.Anything).Return(&entities.UserEntity{
		ID:    1,
		Email: "test@example.com",
	}, http.StatusOK)

	testContext.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBufferString(
		`{"email":"test@example.com","password":"password"}`))

	// call the login handler
	handler.LoginHandler(testContext)

	// assert
	assert.Equal(t, http.StatusOK, testContext.Writer.Status())

	mockService.AssertExpectations(t)
}

func TestLoginHandler_NotFound(t *testing.T) {
	//
	gin.SetMode(gin.TestMode)

	testContext, _ := gin.CreateTestContext(httptest.NewRecorder())

	//  create the mock service
	mockService := mocks.NewService(t)

	// initialize the handler
	handler := handlers.NewAuthHandler(mockService)

	//  stub the Login function
	mockService.On("Login", mock.Anything).Return(nil, http.StatusNotFound)

	testContext.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBufferString(
		`{"email":"test@example.com","password":"password"}`))

	// call the login handler
	handler.LoginHandler(testContext)

	// assert
	assert.Equal(t, http.StatusNotFound, testContext.Writer.Status())

	mockService.AssertExpectations(t)
}

func TestLoginHandler_Unauthorized(t *testing.T) {
	//
	gin.SetMode(gin.TestMode)

	testContext, _ := gin.CreateTestContext(httptest.NewRecorder())

	//  create the mock service
	mockService := mocks.NewService(t)

	// initialize the handler
	handler := handlers.NewAuthHandler(mockService)

	//  stub the Login function
	mockService.On("Login", mock.Anything).Return(nil, http.StatusUnauthorized)

	testContext.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBufferString(
		`{"email":"test@example.com","password":"password"}`))

	// call the login handler
	handler.LoginHandler(testContext)

	// assert
	assert.Equal(t, http.StatusNotFound, testContext.Writer.Status())

	mockService.AssertExpectations(t)
}
