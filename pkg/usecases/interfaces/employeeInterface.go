package interfaces

import (
	"context"
	"shiftsync/pkg/domain"
)

type EmployeeUseCase interface {
	SignUp(r context.Context, signup domain.Employee) error
	Login(r context.Context, login domain.Employee) (domain.Employee, error)
}
