package domain

import "time"

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
	Employee_id          int      `json:"empid" `
	FormID               int      `json:"formid" gorm:"primaryKey;autoIncrement:false"`
	Form                 Employee `gorm:"foreignKey:FormID"`
	Gender               string   `json:"gender" gorm:"type:char(1)"`
	Marital_status       string   `json:"maritalstatus"  gorm:"type:char(1)"`
	Date_of_birth        string   `json:"dateofbirth"`
	P_address            string   `json:"paddress"`
	C_address            string   `json:"caddress"`
	Account_no           string   `json:"accno"`
	Ifsc_code            string   `json:"ifsccode"`
	Name_as_per_passbokk string   `json:"nameaspass"`
	Pan_number           string   `json:"pannumber"`
	Adhaar_no            string   `json:"adhaarnumber"`
	Photo                string   `json:"photo"`
	Status               string   `json:"status" gorm:"type:char(1)"`
	Correction           string   `json:"correction"`
	Designation          string   `json:"designation"`
	Department           string   `json:"department"`
	Approved_by          int      `json:"approved"`
}

type Attendance struct {
	EmployeeID uint   `json:"empid"`
	Punch      Form   `gorm:"foreignKey:EmployeeID"`
	Date       string `json:"date"`
	Punch_in   string `json:"punchin"`
	Punch_out  string `json:"punchout"`
	CreatedAt  time.Time
}

type Leave struct {
	EmployeeID uint   `json:"empid"`
	Leave_type string `json:"leavetype"`
	From       string `json:"fromdate"`
	To         string `json:"todate"`
	Reason     string `json:"reason"`
	Status     string `json:"status" gorm:"type:char(1)"`
}

type Salary struct {
	EmployeeID     uint     `json:"empid" gorm:"primaryKey;autoIncrement:false"`
	Salary         Employee `gorm:"foreignKey:EmployeeID"`
	Grade          string   `json:"grade"`
	Base_salary    int      `json:"basesalary"`
	D_allowance    int      `json:"dallowance"`
	Sp_allowance   int      `json:"spallowance"`
	M_allowance    int      `json:"mallowance"`
	Over_time      int      `json:"overtime"`
	Tax            int      `json:"tax"`
	Provident_fund int      `json:"provident"`
	Gross_salary   int      `json:"grosssalary"`
	Net_salary     int      `json:"netsalary"`
}

type Razorpay struct {
	EmployeeID uint   `json:"empid"`
	Razor      Form   `gorm:"foreignKey:EmployeeID"`
	ContactID  string `json:"contactid"`
}
