package middleware

import (
	"errors"
	"home-work-z-api/model"
	"home-work-z-api/util"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	Account string `json:"account"`
	jwt.StandardClaims
}

func GenerateToken(account string) (string, error) {
	claims := Claims{
		account,
		jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,
			ExpiresAt: time.Now().Unix() + 7200,
			Issuer:    "Alan",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secretKey"))
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")

		if authHeader == model.Empty_string {
			util.Failure(c, "string", http.StatusUnauthorized, "Authorization is null in Header")
			c.Abort()
			return
		}

		claims, err := parseToken(authHeader)

		if err != nil {
			util.Failure(c, "string", http.StatusBadRequest, err.Error())
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}

func parseToken(tokenstring string) (*Claims, error) {

	token, err := jwt.ParseWithClaims(tokenstring, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("secretKey"), nil
	})

	if err != nil {
		if v, ok := err.(*jwt.ValidationError); ok {
			if v.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("That's not even a token")
			} else if v.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("Token is expired")
			} else if v.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("Token not active yet")
			} else {
				return nil, errors.New("Couldn't handle this token:")
			}
		}
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
