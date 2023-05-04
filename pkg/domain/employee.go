package domain

import (
	"time"
)

type Employee struct {
	ID         uint   `json:"id" gorm:"primaryKey;unique"`
	First_name string `json:"firstname"`
	Last_name  string `json:"lastname"`
	Email      string `json:"email"`
	Phone      int64  `json:"phone"`
	User_name  string `json:"username"`
	Pass_word  string `json:"password"`
}

type Form struct {
	Employee_id          int       `json:"empid" `
	FormID               int       `json:"formid" gorm:"primaryKey;autoIncrement:false"`
	Form                 Employee  `gorm:"foreignKey:FormID"`
	Gender               string    `json:"gender" gorm:"type:char(1)"`
	Marital_status       string    `json:"maritalstatus"  gorm:"type:char(1)"`
	Date_of_birth        time.Time `json:"dateofbirth"`
	P_address            string    `json:"paddress"`
	C_address            string    `json:"caddress"`
	Account_no           string    `json:"accno"`
	Ifsc_code            string    `json:"ifsccode"`
	Name_as_per_passbokk string    `json:"nameaspass"`
	Pan_number           string    `json:"pannumber"`
	Adhaar_no            string    `json:"adhaarnumber"`
	Photo                string    `json:"photo"`
	Status               string    `json:"status" gorm:"type:char(1)"`
	Correction           string    `json:"correction"`
	Approved_by          string    `json:"approved"`
}
