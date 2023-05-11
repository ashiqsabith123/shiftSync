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
	ScheduleDuty(ctx context.Context, duty domain.Attendance) error
	GetLeaveRequests(ctx context.Context) ([]response.LeaveRequests, error)
	ApproveLeaveRequests(ctx context.Context, id int) error
	DeclineLeaveRequests(ctx context.Context, id int) error
}
