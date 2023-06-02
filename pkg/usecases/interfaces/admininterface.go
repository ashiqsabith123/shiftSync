package interfaces

import (
	"context"
	"shiftsync/pkg/domain"
	"shiftsync/pkg/helper/response"
)

type AdminUseCase interface {
	SignUp(ctx context.Context, admin domain.Admin) error
	SignIn(ctx context.Context, details domain.Admin) (domain.Admin, error)
	Applications(ctx context.Context) ([]response.Form, error)
	ApproveApplication(ctx context.Context, form domain.Form, empId int) error
	FormCorrection(ctx context.Context, form domain.Form) error
	GetAllEmployeesSchedules(ctx context.Context) ([]response.Schedule, error)
	ScheduleDuty(ctx context.Context, duty domain.Duty) error
	GetLeaveRequests(ctx context.Context) ([]response.LeaveRequests, error)
	ApproveLeaveRequests(ctx context.Context, id int) error
	DeclineLeaveRequests(ctx context.Context, id int) error
	AddSalaryDetails(ctx context.Context, salaryDetails domain.Salary) error
	EditSalaryDetails(ctx context.Context, editDetails domain.Salary) error
	FindEmployeeById(ctx context.Context, id int) response.EmployeeDetails
	FetchAccountDetailsById(ctx context.Context, id int) response.AccountDetails
	GetAllTransactions(ctx context.Context) ([]response.AllTransactions, error)
}
