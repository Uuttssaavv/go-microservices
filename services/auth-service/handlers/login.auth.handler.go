package handlers

import (
	"go-microservices/services/auth-service/models"
	"go-microservices/common/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
		token, err := utilities.SignToken(map[string]interface{}{"id": loginEntity.ID, "email": loginEntity.Email}, 24*60*1)

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