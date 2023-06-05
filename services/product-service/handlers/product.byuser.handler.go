package handlers

import (
	"go-microservices/common/models"
	"go-microservices/common/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetProductByUser(ctx *gin.Context) {

	var user models.UserEntity
	value, _ := ctx.Get("user")

	err := utilities.Unmarshal(value, &user)

	if err != nil {
		utilities.APIResponse(ctx, "Could not parse the parameter", http.StatusBadRequest, nil)
		return
	}

	productResp, statusCode := h.services.GetProductsByUser(user.ID)

	switch statusCode {

	case http.StatusNotFound:
		utilities.APIResponse(ctx, "Could not find the products for given users", statusCode, nil)
		return

	case http.StatusFound:
		utilities.APIResponse(ctx, "Fetched products", statusCode, productResp)
		return

	default:
		utilities.APIResponse(ctx, "Something in the server went wrong", http.StatusInternalServerError, nil)
	}
}
