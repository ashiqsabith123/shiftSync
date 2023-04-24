package interfaces

import "context"

type EmployeeUseCase interface {
	SignUp(r context.Context)
}
