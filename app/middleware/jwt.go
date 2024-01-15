package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"my-blog/app/dto"
	"net/http"
	"time"
)

var secret = []byte("my-blog")

type MyCustomClaims struct {
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
	jwt.RegisteredClaims
}

func GenerateJWT(data *dto.LoginParamsData) (string, error) {
	claims := MyCustomClaims{
		data.Phone, data.Password, data.Password,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			Issuer:    "test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(secret)
	return ss, err
}

func ParseToken(token string) (*MyCustomClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*MyCustomClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		//code = e.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = 1
		} else {
			claims, err := ParseToken(token)
			if err != nil {
				code = 1
			} else if time.Now().Unix() > claims.ExpiresAt.Unix() {
				code = 1
			}
		}

		if code == 1 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  "1",
				"data": nil,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
