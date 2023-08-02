package interfaces

import (
	"context"
	"shiftsync/pkg/domain"
	"shiftsync/pkg/helper/request"
	"shiftsync/pkg/helper/response"
)

type EmployeeUseCase interface {
	SignUp(r context.Context, signup domain.Employee) error
	Login(r context.Context, login domain.Employee) (domain.Employee, error)
	SignUpOtp(r context.Context, find domain.Employee) error
	AddForm(r context.Context, form domain.Form) error
	FormStatus(ctx context.Context, empID int) (response.FormStatus, error)
	GetDutySchedules(ctx context.Context, id int) (response.Duty, error)
	PunchIn(ctx context.Context, ID int) (string, error)
	VerifyOtpForPunchin(ctx context.Context, id int, otp request.OtpStruct) error
	PunchOut(ctx context.Context, id int) error
	ApplyLeave(ctx context.Context, leave domain.Leave) (string, error)
	GetLeaveStatusHistory(ctx context.Context, id int) ([]response.LeaveHistory, error)
	Attendance(ctx context.Context, id int) ([]response.Attendance, error)
	GetSalaryHistory(ctx context.Context, id int) ([]response.Salaryhistory, error)
	GetSalaryDetails(ctx context.Context, id int) (response.Salarydetails, error)
	GetDataForSalarySlip(ctx context.Context, id int) ([]byte, error)
}
