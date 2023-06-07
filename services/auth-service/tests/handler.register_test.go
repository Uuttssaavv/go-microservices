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

func TestRegisterHandler_BindingError(t *testing.T) {
	//
	gin.SetMode(gin.TestMode)

	testContext, _ := gin.CreateTestContext(httptest.NewRecorder())

	//  create the mock service
	mockService := mocks.NewService(t)

	// initialize the handler
	handler := handlers.NewAuthHandler(mockService)

	testContext.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBufferString(
		`invalid-json`))

	// call the register handler
	handler.RegisterHandler(testContext)

	// assert
	assert.Equal(t, http.StatusBadRequest, testContext.Writer.Status())

	mockService.AssertExpectations(t)
}

func TestRegisterHandler_Success(t *testing.T) {
	//
	gin.SetMode(gin.TestMode)

	testContext, _ := gin.CreateTestContext(httptest.NewRecorder())

	//  create the mock service
	mockService := mocks.NewService(t)

	// initialize the handler
	handler := handlers.NewAuthHandler(mockService)

	//  stub the Register function
	mockService.On("Register", mock.Anything).Return(&entities.UserEntity{
		ID:    1,
		Email: "test@example.com",
	}, http.StatusOK)

	testContext.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBufferString(
		`{"email":"test@example.com","password":"password"}`))

	// call the register handler
	handler.RegisterHandler(testContext)

	// assert
	assert.Equal(t, http.StatusOK, testContext.Writer.Status())

	mockService.AssertExpectations(t)
}

func TestRegisterHandler_AlreadyExist(t *testing.T) {
	//
	gin.SetMode(gin.TestMode)

	testContext, _ := gin.CreateTestContext(httptest.NewRecorder())

	//  create the mock service
	mockService := mocks.NewService(t)

	// initialize the handler
	handler := handlers.NewAuthHandler(mockService)

	//  stub the Register function
	mockService.On("Register", mock.Anything).Return(nil, http.StatusConflict)

	testContext.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBufferString(
		`{"email":"test@example.com","password":"password"}`))

	// call the register handler
	handler.RegisterHandler(testContext)

	// assert
	assert.Equal(t, http.StatusConflict, testContext.Writer.Status())

	mockService.AssertExpectations(t)
}

func TestRegisterHandler_Failure(t *testing.T) {
	//
	gin.SetMode(gin.TestMode)

	testContext, _ := gin.CreateTestContext(httptest.NewRecorder())

	//  create the mock service
	mockService := mocks.NewService(t)

	// initialize the handler
	handler := handlers.NewAuthHandler(mockService)

	//  stub the Register function
	mockService.On("Register", mock.Anything).Return(nil, http.StatusExpectationFailed)

	testContext.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBufferString(
		`{"email":"test@example.com","password":"password"}`))

	// call the register handler
	handler.RegisterHandler(testContext)

	// assert
	assert.Equal(t, http.StatusExpectationFailed, testContext.Writer.Status())

	mockService.AssertExpectations(t)
}
