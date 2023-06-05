package handlers

import (
	entities "go-microservices/common/models"
	"go-microservices/common/utilities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) DeleteProduct(ctx *gin.Context) {

	productId, err := strconv.ParseInt(ctx.Param("productId"), 0, 0)

	var user entities.UserEntity

	value, _ := ctx.Get("user")

	utilities.Unmarshal(value, &user)

	if err != nil {
		utilities.APIResponse(ctx, "Could not parse the parameter", http.StatusBadRequest, nil)

		return
	}

	statusCode := h.services.DeleteProduct(uint(productId), user.ID)

	switch statusCode {
	case http.StatusOK:

		utilities.APIResponse(ctx, "Deleted product.", statusCode, nil)
		return
	case http.StatusNotFound:

		utilities.APIResponse(ctx, "Product with the id doesnot exist", statusCode, nil)
		return

	case http.StatusUnauthorized:

		utilities.APIResponse(ctx, "You are not authorized to delete the product", statusCode, nil)
		return

	case http.StatusLocked:
		utilities.APIResponse(ctx, "Could not delete the product.", statusCode, nil)
		return
		
	default:
		utilities.APIResponse(ctx, "Could not delete the product.", statusCode, nil)
		return
	}

}
