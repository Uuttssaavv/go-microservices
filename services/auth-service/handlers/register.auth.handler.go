package handlers

import (
	"go-microservices/services/auth-service/models"
	"go-microservices/common/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)


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
	
	switch statusCode {
	case http.StatusCreated:

		token, err := utilities.SignToken(map[string]interface{}{"id": registerEntity.ID, "email": registerEntity.Email}, 24*60*1)

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
