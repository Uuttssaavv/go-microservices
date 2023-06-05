package middlewares

import (
	"go-microservices/common/utilities"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	//
	return gin.HandlerFunc(func(ctx *gin.Context) {
		//
		authorization := ctx.GetHeader("Authorization")
		if authorization == "" {
			// return authorization required response
			utilities.APIResponse(ctx, "No authorization provided in the header", http.StatusBadRequest, nil)
			defer ctx.AbortWithStatus(http.StatusBadRequest)
		}

		// remove the Bearer from the authorization
		jwtToken := strings.SplitAfter(authorization, "Bearer ")[1]

		// verify the token
		token, err := utilities.VerifyToken(jwtToken)

		if err != nil {
			utilities.APIResponse(ctx, "Token invalid or expired", http.StatusUnauthorized, nil)
			defer ctx.AbortWithStatus(http.StatusUnauthorized)
		} else {
			//  set the token claims to the context
			ctx.Set("user", token.Claims)
			// and continue the request
			ctx.Next()
		}

	})
}
