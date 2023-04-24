package interfaces

import (
	"context"
	"shiftsync/pkg/domain"
)

type EmployeeRepository interface {
	AddEmployee(cntxt context.Context, signup domain.Employee_Signup) error
}
