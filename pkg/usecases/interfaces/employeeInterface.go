package interfaces

import (
	"context"
	"shiftsync/pkg/domain"
	"shiftsync/pkg/helper/response"
)

type EmployeeUseCase interface {
	SignUp(r context.Context, signup domain.Employee) error
	Login(r context.Context, login domain.Employee) (domain.Employee, error)
	SignUpOtp(r context.Context, find domain.Employee) error
	AddForm(r context.Context, form domain.Form) error
	FormStatus(ctx context.Context, empID int) string
	GetDutySchedules(ctx context.Context, id int) (response.Duty, error)
}
