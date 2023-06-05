package handlers

import (
	enitities "go-microservices/common/models"
	"go-microservices/common/utilities"
	"go-microservices/services/product-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateProduct(ctx *gin.Context) {

	var user enitities.UserEntity

	jwtClaims, _ := ctx.Get("user")

	errors := utilities.Unmarshal(jwtClaims, &user)

	if errors != nil {
		utilities.APIResponse(ctx, "Cannot parse claims", http.StatusBadRequest, nil)
		return
	}

	var input models.ProductInput

	bindingError := ctx.ShouldBind(&input)

	if bindingError != nil {
		errorResponse := utilities.ParseBindingError(bindingError)

		utilities.ValidatorErrorResponse(ctx, http.StatusBadRequest, errorResponse)
		return

	}
	input.UserID = user.ID

	createResp, statusCode := h.services.CreateProduct(input)

	switch statusCode {
	case http.StatusCreated:
		utilities.APIResponse(ctx, "Created post successfuly", statusCode, createResp)
		return
	case http.StatusNotAcceptable:
		utilities.APIResponse(ctx, "User not created", statusCode, nil)
		return
	default:
		utilities.ValidatorErrorResponse(ctx, statusCode, "Unknown error occured")
	}
}
