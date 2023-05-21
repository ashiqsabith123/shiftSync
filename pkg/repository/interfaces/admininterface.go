package interfaces

import (
	"context"
	"shiftsync/pkg/domain"
	"shiftsync/pkg/helper/request"
	"shiftsync/pkg/helper/response"
)

type AdminRepository interface {
	FindAdmin(ctx context.Context, find domain.Admin) (domain.Admin, error)
	SaveAdmin(ctx context.Context, admin domain.Admin) error
	GetAllForms(ctx context.Context) ([]response.Form, error)
	ApproveApplication(ctx context.Context, form domain.Form, id int, empId int)
	FindFormByID(ctx context.Context, fID int) error
	FormCorrection(ctx context.Context, form domain.Form)
	GetAllEmployees(ctx context.Context) ([]response.AllEmployee, error)
	GetAllEmployeesSchedules(ctx context.Context) ([]response.Schedule, error)
	ScheduleDuty(ctx context.Context, duty domain.Attendance) error
	GetLeaveRequests(ctx context.Context) ([]response.LeaveRequests, error)
	ChangeLeaveStatus(ctx context.Context, status request.LeaveStatus) error
	AddSalaryDetails(ctx context.Context, salaryDetails domain.Salary) error
	EditSalaryDetails(ctx context.Context, editDetails domain.Salary) error
}
