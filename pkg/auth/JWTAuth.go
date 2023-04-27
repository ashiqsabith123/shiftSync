package auth

import (
	"fmt"
	"shiftsync/pkg/config"
	"shiftsync/pkg/domain"
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

type Values struct {
	Full_name string `json:"fullname"`
	Email     string `json:"email"`
	Phone     int64  `json:"phone"`
	User_name string `json:"username"`
	Pass_word string `json:"password"`
	jwt.StandardClaims
}

func GenerateTokenForOtp(val domain.Employee) (string, error) {

	claims := Values{
		Full_name: val.Full_name,
		Email:     val.Email,
		Phone:     val.Phone,
		User_name: val.User_name,
		Pass_word: val.Pass_word,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiryTime,
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))

	return token, err

}

func ValidatOtpTokens(signedtoken string) (Values, error) {
	token, err := jwt.ParseWithClaims(
		signedtoken, &Values{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		})

	if err != nil {

		return Values{}, err
	}

	claim, _ := token.Claims.(*Values)

	return *claim, nil
}
