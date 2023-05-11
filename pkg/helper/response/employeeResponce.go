package response

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
