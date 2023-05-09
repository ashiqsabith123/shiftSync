package interfaces

import (
	"context"
	"shiftsync/pkg/domain"
	"shiftsync/pkg/helper/response"
)

type EmployeeRepository interface {
	AddEmployee(cntxt context.Context, signup domain.Employee) error
	FindEmployee(cntxt context.Context, find domain.Employee) (domain.Employee, error)
	AddForm(cntxt context.Context, form domain.Form) error
	CheckFormDetails(cntxt context.Context, form domain.Form) error
	FormStatus(ctx context.Context, empID int) string
	GetDutySchedules(ctx context.Context, id int) (response.Duty, error)
	PunchIn(ctx context.Context, punchin domain.Attendance) error
	PunchOut(ctx context.Context, punchout domain.Attendance) error
	ApplyLeave(ctx context.Context, leave domain.Leave) error
}
