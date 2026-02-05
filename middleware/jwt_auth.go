package middleware

import (
	"log"
	"net/http"
	"runtime"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/naveeharn/hospital-information-service-backend/helper"
	"github.com/naveeharn/hospital-information-service-backend/internal/service"
)

func AuthorizeJWT(jwtService service.JwtService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" || !strings.Contains(authHeader, "Bearer ") {
			result := helper.Result{
				Code: "999",
				Message: "Token Authorization not found or not correct header pattern",
			}
			response := helper.CreateErrorResponse("", nil, result)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		token, err := jwtService.ValidateToken(strings.TrimPrefix(authHeader, "Bearer "))
		if err != nil {
			helper.LoggerErrorPath(runtime.Caller(0))
			log.Println(err.Error())
			result := helper.Result{
				Code: "999",
				Message: "Token is not valid",
			}
			response := helper.CreateErrorResponse("", err.Error(), result)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			helper.LoggerErrorPath(runtime.Caller(0))
			result := helper.Result{
				Code: "999",
				Message: "Token is not valid",
			}
			response := helper.CreateErrorResponse("", nil, result)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		id, ok := claims["id"].(string)
		if !ok {
			result := helper.Result{
				Code: "999",
				Message: "Token can not parse",
			}
			response := helper.CreateErrorResponse("", nil, result)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		ctx.Set("id", id)
	}
}