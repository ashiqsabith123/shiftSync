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
	First_name string `json:"firstname"`
	Last_name  string `json:"lastname"`
	Email      string `json:"email"`
	Phone      int64  `json:"phone"`
	User_name  string `json:"username"`
	Pass_word  string `json:"password"`
}

type Form struct {
	Gender               string `json:"gender" gorm:"type:char(1)"`
	Marital_status       string `json:"maritalstatus"  gorm:"type:char(1)"`
	Date_of_birth        string `json:"dateofbirth"`
	P_address            string `json:"paddress"`
	C_address            string `json:"caddress"`
	Account_no           string `json:"accno"`
	Ifsc_code            string `json:"ifsccode"`
	Name_as_per_passbokk string `json:"nameinpass"`
	Pan_number           string `json:"pannumber"`
	Adhaar_no            string `json:"adhaarno"`
	Designation          string `json:"designation"`
	Photo                string `json:"photo"`
}

type OtpStruct struct {
	Code string `json:"otp"`
}

type OtpCookieStruct struct {
	First_name string `json:"firstname"`
	Last_name  string `json:"lastname"`
	Email      string `json:"email"`
	Phone      int64  `json:"phone"`
	User_name  string `json:"username"`
	Pass_word  string `json:"password"`
	jwt.StandardClaims
}

type FormApprove struct {
	FormID int `json:"id"`
}

type FormCorrection struct {
	FormID     int    `json:"empid"`
	Correction string `json:"correction"`
}

type DutySchedule struct {
	EmployeeID int    `json:"empid"`
	Duty_type  string `json:"dutytype"`
}

type PunchIn struct {
	Id      int
	Date    string
	Time_in time.Time
}

type Leave struct {
	Leave_type string `json:"leavetype"`
	From       string `json:"fromdate"`
	To         string `json:"todate"`
	Reason     string `json:"reason"`
}

type LeaveStatus struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}
