package interfaces

import (
	"context"
	"shiftsync/pkg/domain"
)

type EmployeeRepository interface {
	AddEmployee(cntxt context.Context, signup domain.Employee) error
	FindEmployee(cntxt context.Context, find domain.Employee) (domain.Employee, error)
	AddForm(cntxt context.Context, form domain.Form) error
}
