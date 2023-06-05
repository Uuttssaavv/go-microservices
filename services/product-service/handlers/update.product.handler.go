package handlers

import (
	enitities "go-microservices/common/models"
	"go-microservices/common/utilities"
	"go-microservices/services/product-service/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) UpdateProduct(ctx *gin.Context) {

	productId, err := strconv.ParseInt(ctx.Param("productId"), 0, 0)

	if err != nil {
		utilities.APIResponse(ctx, "Bad parameter in the request", http.StatusBadRequest, nil)
		return
	}
	// get user set by the middleware
	var user enitities.UserEntity

	jwtClaims, _ := ctx.Get("user")

	errors := utilities.Unmarshal(jwtClaims, &user)

	if errors != nil {
		utilities.APIResponse(ctx, "Cannot parse claims", http.StatusBadRequest, nil)
		return
	}

	var input models.UpdateProductInput
	//  bind json to product struct

	err = ctx.ShouldBind(&input)

	if err != nil {
		bindingErrors := utilities.ParseBindingError(err)

		utilities.ValidatorErrorResponse(ctx, http.StatusBadRequest, bindingErrors)
		return
	}

	input.ID = uint(productId)
	input.UserID = user.ID

	productResp, statusCode := h.services.UpdateProduct(input)

	switch statusCode {
	case http.StatusOK:
		utilities.APIResponse(ctx, "Updated the product", statusCode, productResp)
		return

	case http.StatusForbidden:
		utilities.APIResponse(ctx, "Could not update the product", statusCode, nil)
		return

	case http.StatusNotFound:
		utilities.APIResponse(ctx, "The product does not exist for the user", statusCode, nil)
		return
		
	default:
		utilities.APIResponse(ctx, "Something went wrong. Please try again", http.StatusInternalServerError, nil)
	}

}
