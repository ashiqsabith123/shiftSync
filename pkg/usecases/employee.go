package usecases

import (
	"context"
	service "shiftsync/pkg/usecases/interfaces"
)

type employeeUseCase struct {
}

func NewEmployeeUseCase() service.EmployeeUseCase {
	return &employeeUseCase{}
}

func (u *employeeUseCase) SignUp(context.Context) {

}
