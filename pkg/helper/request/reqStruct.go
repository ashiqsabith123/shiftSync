package request

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type LoginStruct struct {
	User_name string `json:"username"`
	Pass_word string `json:"password"`
}

type SignUp struct {
	Full_name string `json:"fullname"`
	Email     string `json:"email"`
	Phone     int64  `json:"phone"`
	User_name string `json:"username"`
	Pass_word string `json:"password"`
}

type Form struct {
	First_name           string    `json:"firstname"`
	Last_name            string    `json:"lastname"`
	Email                string    `json:"email"`
	Gender               string    `json:"gender" gorm:"type:char(1)"`
	Marital_status       string    `json:"maritalstatus"  gorm:"type:char(1)"`
	Phone                int64     `json:"phone"`
	Date_of_birth        time.Time `json:"dateofbirth"`
	P_address            string    `json:"paddress"`
	C_address            string    `json:"caddress"`
	Account_no           string    `json:"accno"`
	Ifsc_code            string    `json:"ifsccode"`
	Name_as_per_passbokk string    `json:"nameinpass"`
	Pan_number           string    `json:"pannumber"`
	Adhaar_no            int64     `json:"adhaarno"`
	Photo                string    `json:"photo"`
}

type OTPStruct struct {
	Code string `json:"otp"`
}

type OtpCookieStruct struct {
	Full_name string `json:"fullname"`
	Email     string `json:"email"`
	Phone     int64  `json:"phone"`
	User_name string `json:"username"`
	Pass_word string `json:"password"`
	jwt.StandardClaims
}
