package response

import "time"

type Duty struct {
	Duty_type string `json:"type"`
	Time      string `json:"time"`
}

type LeaveHistory struct {
	Leave_type string `json:"leavetype"`
	From       string `json:"fromdate"`
	To         string `json:"todate"`
	Status     string `json:"status"`
}

type Attendance struct {
	Date        string `json:"date"`
	Punch_in    string `json:"timein"`
	Punch_out   string `json:"timeout"`
	Duty_type   string `json:"dutytype"`
	Total_hours int    `json:"totalhour"`
	Over_time   int    `json:"overtime"`
}

type LeaveAppiled struct {
	From string
	To   string
}

type LeaveCount struct {
	Id   int
	Date int
}

type SalaryHistory struct {
	Refrence_id  string    `json:"rerence_id"`
	Date         time.Time `json:"date"`
	Time         time.Time `json:"time"`
	Allowance    int       `json:"allowance"`
	Deductions   int       `json:"deductions"`
	Net_salary   int       `json:"net_salary"`
	Gross_salary int       `json:"gross_salary"`
}

type Salaryhistory struct {
	Refrence_id  string `json:"rerence_id"`
	Date         string `json:"date"`
	Time         string `json:"time"`
	Allowance    int    `json:"allowance"`
	Deductions   int    `json:"deductions"`
	Net_salary   int    `json:"net_salary"`
	Gross_salary int    `json:"gross_salary"`
}

type Salarydetails struct {
	Grade          string `json:"grade"`
	Base_salary    int    `json:"basesalary"`
	Bonus          int    `json:"bonus"`
	Leave_pay      int    `json:"leave_pay"`
	D_allowance    int    `json:"dallowance"`
	Sp_allowance   int    `json:"spallowance"`
	M_allowance    int    `json:"mallowance"`
	Over_time      int    `json:"overtime"`
	Tax            int    `json:"tax"`
	Provident_fund int    `json:"provident"`
	Gross_salary   int    `json:"grosssalary"`
	Net_salary     int    `json:"netsalary"`
}

type SalarySlip struct {
	Employee_id    string
	Name           string
	Designation    string
	Account_no     string
	Grade          string
	Duties         string
	Leave_count    string
	Base_salary    string
	D_allowance    string
	Sp_allowance   string
	M_allowance    string
	Leave_pay      string
	Over_time      string
	Provident_fund string
	Tax            string
	Gross_salary   string
	Net_salary     string
	Deductions     string
}

type FormStatus struct {
	Status     string
	Correction string
}
