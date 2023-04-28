package interfaces

import (
	"context"
	"shiftsync/pkg/domain"
)

type EmployeeUseCase interface {
	SignUp(r context.Context, signup domain.Employee) error
	Login(r context.Context, login domain.Employee) (domain.Employee, error)
	SignUpOtp(r context.Context, find domain.Employee) error
	AddForm(r context.Context, form domain.Form)
}
