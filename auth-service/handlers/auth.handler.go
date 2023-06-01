package handlers

import (
	"fmt"
	"go-microservices/auth-service/models"
	"go-microservices/auth-service/services"
	"go-microservices/common/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service services.Service
}

func NewAuthHandler(service services.Service) *handler {
	return &handler{service: service}
}

func (handler *handler) LoginHandler(context *gin.Context) {
	//parse the body
	var input models.UserInputModel

	err := context.ShouldBindJSON(&input)

	if err != nil {
		errors := utilities.ParseBindingError(err)
		utilities.ValidatorErrorResponse(context, http.StatusBadRequest, errors)
		return
	}
	// call the service
	loginEntity, statusCode := handler.service.Login(&input)

	switch statusCode {
	case http.StatusOK:
		token, err := utilities.Sign(map[string]interface{}{"id": loginEntity.ID, "email": loginEntity.Email}, 24*60*1)

		if err != nil {
			utilities.APIResponse(context, "Cannt generate token", http.StatusInternalServerError, err)
			return
		}
		var userResponse models.UserResponse

		utilities.Unmarshal(loginEntity, &userResponse)
		userResponse.Token = token

		utilities.APIResponse(context, "Login successful", http.StatusOK, userResponse)

		return
	case http.StatusNotFound:
		utilities.APIResponse(context, "User does not exist", http.StatusNotFound, err)

		return
	case http.StatusUnauthorized:
		utilities.APIResponse(context, "Invalid email or password", http.StatusNotFound, err)
	}

}

func (handler *handler) RegisterHandler(context *gin.Context) {
	//parse the body
	var input models.UserInputModel

	err := context.ShouldBindJSON(&input)

	if err != nil {
		errors := utilities.ParseBindingError(err)
		utilities.ValidatorErrorResponse(context, http.StatusBadRequest, errors)
		return
	}
	// call the service
	registerEntity, statusCode := handler.service.Register(&input)
	fmt.Printf("%+v", registerEntity)
	switch statusCode {
	case http.StatusCreated:

		token, err := utilities.Sign(map[string]interface{}{"id": registerEntity.ID, "email": registerEntity.Email}, 24*60*1)

		if err != nil {
			utilities.APIResponse(context, "Cannt generate token", http.StatusInternalServerError, err)
			return
		}

		var userResponse models.UserResponse

		utilities.Unmarshal(registerEntity, &userResponse)
		userResponse.Token = token

		utilities.APIResponse(context, "Registered the user", http.StatusCreated, userResponse)
		return

	case http.StatusConflict:
		utilities.APIResponse(context, "User already exists", http.StatusConflict, err)
		return

	case http.StatusExpectationFailed:
		utilities.APIResponse(context, "Unable to create user", http.StatusExpectationFailed, err)
		return
	}

}
