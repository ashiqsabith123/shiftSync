package auth

import (
	"fmt"
	"shiftsync/pkg/config"
	"time"

	"github.com/golang-jwt/jwt"
)

var expiryTime = time.Now().Add(10 * time.Minute).Unix()
var SECRET_KEY = config.JwtConfig()

func GenerateTokens(id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: expiryTime,
		Id:        fmt.Sprint(id),
	})

	generatedTokens, err := token.SignedString([]byte(SECRET_KEY))

	return generatedTokens, err

}
