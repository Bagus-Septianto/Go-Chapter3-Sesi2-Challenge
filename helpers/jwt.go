package helpers

import (
	"errors"
	"strings"

	jwt294 "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = "rahasia"

func GenerateToken(id uint, username string, role string) string {
	claims := jwt294.MapClaims{
		"id": id,
		"username": username,
		"role": role,
	}

	parseToken := jwt294.NewWithClaims(jwt294.SigningMethodHS256, claims)

	signedToken, _ := parseToken.SignedString([]byte(secretKey))

	return signedToken
}

func VerifyToken(c *gin.Context) (interface{}, error) {
	errResponse := errors.New("sign in to proceed") // setting error message
	headerToken := c.Request.Header.Get("Authorization") // ambil header "Authorization"
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer { // tidak ada awalan "Bearer" pada Authorization
		return nil, errResponse
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt294.Parse(stringToken, 
		func(t *jwt294.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt294.SigningMethodHMAC); !ok {
				return nil, errResponse
			}
			return []byte(secretKey), nil
		})

	if _, ok := token.Claims.(jwt294.MapClaims); !ok && !token.Valid {
		return nil, errResponse
	}

	return token.Claims.(jwt294.MapClaims), nil
}