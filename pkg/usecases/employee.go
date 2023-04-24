package usecases

import (
	"context"
	"errors"
	"shiftsync/pkg/domain"
	repo "shiftsync/pkg/repository/interfaces"
	service "shiftsync/pkg/usecases/interfaces"

	"golang.org/x/crypto/bcrypt"
)

type employeeUseCase struct {
	employeeRepo repo.EmployeeRepository
}

func NewEmployeeUseCase(rep repo.EmployeeRepository) service.EmployeeUseCase {
	return &employeeUseCase{employeeRepo: rep}
}

func (u *employeeUseCase) SignUp(cntxt context.Context, signup domain.Employee_Signup) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(signup.Pass_word), 14)
	if err != nil {
		return errors.New("bcrypt failed:" + err.Error())
	}

	signup.Pass_word = string(hash)
	return u.employeeRepo.AddEmployee(cntxt, signup)
}
