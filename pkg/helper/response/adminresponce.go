package response

type Form struct {
	ID                   int    `json:"id"`
	First_name           string `json:"firstname"`
	Last_name            string `json:"lastname"`
	Email                string `json:"email"`
	Gender               string `json:"gender" gorm:"type:char(1)"`
	Marital_status       string `json:"maritalstatus"  gorm:"type:char(1)"`
	Phone                int64  `json:"phone"`
	Date_of_birth        string `json:"dateofbirth"`
	P_address            string `json:"paddress"`
	C_address            string `json:"caddress"`
	Account_no           string `json:"accno"`
	Ifsc_code            string `json:"ifsccode"`
	Name_as_per_passbokk string `json:"nameaspass"`
	Pan_number           string `json:"pannumber"`
	Designation          string `json:"designation"`
	Department           string `json:"department"`
	Adhaar_no            string `json:"adhaarnumber"`
	Photo                string `json:"photo"`
}

type AllEmployee struct {
	Id            int    `json:"id"`
	Empid         int    `json:"empid"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Phone         int64  `json:"phone"`
	Gender        string `json:"gender"`
	Date_of_birth string `json:"dateofbirth"`
	Department    string `json:"department"`
	Designation   string `json:"designation"`
}

type Schedule struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       int64  `json:"phone"`
	Designation string `json:"designation"`
}

type LeaveRequests struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	From       string `json:"fromdate"`
	To         string `json:"todate"`
	Leave_type string `json:"leavetype"`
	Reason     string `json:"reason"`
}
