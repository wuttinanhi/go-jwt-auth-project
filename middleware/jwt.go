package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	jwtservice "github.com/wuttinanhi/go-jwt-auth-project/jwt-service"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get authorization header
		authHeader := c.GetHeader("Authorization")

		// remove "Bearer: " from authHeader
		token := authHeader[8:]

		// validate token
		if !jwtservice.GetJWTService().ValidateToken(token) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// get user id from token
		authJwt := &jwtservice.AuthJWT{}

		// parse token
		jwtservice.GetJWTService().ParseToken(token, authJwt)

		// set user id to context
		c.Set("userId", authJwt.UserId)

		// proceed to next middleware
		c.Next()

		defer func() {
			if err := recover(); err != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
			}
		}()
	}
}
