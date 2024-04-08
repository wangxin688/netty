package middleware

import (
	"net/http"
	"strings"

	"netty/core"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// AuthMiddleware checks the presence and validity of the JWT token
// in the request's Authorization header and extracts the user ID
// and username from the claims
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader(core.AuthorizationString)

		// If no token is provided, return Unauthorized
		if tokenString == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Unauthorized, no token provided",
				"data":    nil,
			})
			return
		}

		p := strings.Split(tokenString, " ")
		if len(p) != 2 || p[0] != core.AuthorizationBearer {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Unauthorized, invalid token",
				"data":    nil,
			})
			return
		}

		token, err := jwt.ParseWithClaims(p[1], &core.Claims{}, func(token *jwt.Token) (interface{}, error) {
			if core.Settings.SecretKey == "" {
				panic("No secret key provided")
			}
			return []byte(core.Settings.SecretKey), nil
		})

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Unauthorized, invalid token",
				"data":    nil,
			})
			return
		}

		if claims, ok := token.Claims.(*core.Claims); ok && token.Valid {
			if claims.Refreshed {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"code":    http.StatusUnauthorized,
					"message": "Unauthorized, refresh token is not valid for authentication",
					"data":    nil,
				})
				return
			}
			ctx.Set("userID", claims.UserId)
			ctx.Set("username", claims.Username)
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Unauthorized, invalid token",
				"data":    nil,
			})
			return
		}
		ctx.Next()
	}
}
