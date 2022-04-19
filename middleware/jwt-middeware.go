package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/widodragon/goshop/service"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := ctx.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]
		if tokenString == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		}
		token, err := service.NewJWTService().ValidateToken(tokenString)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[Name]: ", claims["name"])
			log.Println("Claims[Admin]: ", claims["admin"])
			log.Println("Claims[Issuer]: ", claims["iss"])
			log.Println("Claims[IssuedAt]: ", claims["iat"])
			log.Println("Claims[ExpiresAt]: ", claims["exp"])

		} else {
			log.Println(err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		}
		ctx.Next()

	}
}
