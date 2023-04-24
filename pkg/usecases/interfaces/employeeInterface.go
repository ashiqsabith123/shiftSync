package interfaces

import (
	"context"
	"shiftsync/pkg/domain"
)

type EmployeeUseCase interface {
	SignUp(r context.Context, signup domain.Employee_Signup) error
}
