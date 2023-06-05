package handlers

import (
	"go-microservices/common/utilities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetProductDetail(ctx *gin.Context) {
	productId, err := strconv.ParseInt(ctx.Param("productId"), 0, 0)

	if err != nil {
		utilities.APIResponse(ctx, "Could not parse the parameter", http.StatusBadRequest, nil)
		return
	}

	productResp, statusCode := h.services.GetProductById(uint(productId))

	switch statusCode {

	case http.StatusNotFound:
		utilities.APIResponse(ctx, "Could not find the product with the given id", statusCode, nil)
		return

	case http.StatusFound:
		utilities.APIResponse(ctx, "Fetched product", statusCode, productResp)
		return

	default:
		utilities.APIResponse(ctx, "Something in the server went wrong", http.StatusInternalServerError, nil)
	}
}
