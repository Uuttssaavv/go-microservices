package handlers

import (
	"fmt"
	"go-microservices/common/utilities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetProductDetail(ctx *gin.Context) {
	productId, err := strconv.ParseInt(ctx.Param("poductId"), 0, 0)

	if err != nil {
		utilities.APIResponse(ctx, "Could not parse the parameter", http.StatusBadRequest, nil)
	}

	productResp, statusCode := h.services.GetProductById(uint(productId))

	fmt.Printf("%+v", productResp)
	println(statusCode)
}
