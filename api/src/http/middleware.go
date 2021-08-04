package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JSONMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.ContentType() != gin.MIMEJSON {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid content-type"})
			return
		}
	}
}

type UserLogged struct {
	Id   int32
	Role string
}

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {

		os.Setenv("ACCESS_SECRET", "jdnfksdmfksd")

		bearer := c.GetHeader("Authorization")
		if bearer == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
			return
		}
		strArr := strings.Split(bearer, " ")

		if len(strArr) < 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
			return
		}
		claims := jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(strArr[1], claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("ACCESS_SECRET")), nil
		})

		if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid || err != nil {

			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		userLogged := &UserLogged{
			Id:   int32(claims["user_id"].(float64)),
			Role: claims["role"].(string),
		}

		c.Set("userLogged", userLogged)
	}
}

func IsManager() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, ok := c.Get("role")
		if ok {
			if role != "manager" {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": http.StatusText(http.StatusForbidden)})
				return
			}
		}
	}
}
