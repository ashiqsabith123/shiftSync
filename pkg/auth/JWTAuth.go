package auth

import (
	"errors"
	"fmt"
	"shiftsync/pkg/config"
	"shiftsync/pkg/domain"
	"shiftsync/pkg/helper/request"
	"time"

	"github.com/golang-jwt/jwt"
)

var expiryTime = time.Now().Add(10 * time.Minute).Unix()

func GenerateTokens(id uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: expiryTime,
		Id:        fmt.Sprint(id),
	})

	generatedTokens, err := token.SignedString([]byte(config.JwtConfig()))

	return generatedTokens, err

}

func GenerateTokenForOtp(val domain.Employee) (string, error) {

	claims := request.OtpCookieStruct{
		Full_name: val.Full_name,
		Email:     val.Email,
		Phone:     val.Phone,
		User_name: val.User_name,
		Pass_word: val.Pass_word,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiryTime,
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.JwtConfig()))

	return token, err

}

func ValidateTokens(signedtoken string) (jwt.StandardClaims, error) {
	// token, err := jwt.ParseWithClaims(
	// 	signedtoken, jwt.
	// 		StandardClaims{},
	// 	func(token *jwt.Token) (interface{}, error) {

	// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	// 		}

	// 		return []byte(config.JwtConfig()), nil
	// 	},
	// )

	// if err != nil || !token.Valid {
	// 	return jwt.StandardClaims{}, errors.New("not valid token")
	// }

	// // then parse the token to claims
	// claims, ok := token.Claims.(*jwt.StandardClaims)
	// if !ok {
	// 	return jwt.StandardClaims{}, errors.New("can't parse the claims")
	// }

	token, err := jwt.ParseWithClaims(
		signedtoken, &jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(config.JwtConfig()), nil
		})

	if err != nil || !token.Valid {
		return jwt.StandardClaims{}, errors.New("not valid tok")
	}

	if err != nil {

		return jwt.StandardClaims{}, err
	}

	claim, _ := token.Claims.(*jwt.StandardClaims)

	return *claim, nil
}

func ValidateOtpTokens(signedtoken string) (request.OtpCookieStruct, error) {
	token, err := jwt.ParseWithClaims(
		signedtoken, &request.OtpCookieStruct{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.JwtConfig()), nil
		})

	if err != nil {

		return request.OtpCookieStruct{}, err
	}

	claim, _ := token.Claims.(*request.OtpCookieStruct)

	return *claim, nil
}