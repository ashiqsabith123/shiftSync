package usecases

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"shiftsync/pkg/domain"
	"shiftsync/pkg/encrypt"
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

func (u *employeeUseCase) SignUp(cntxt context.Context, signup domain.Employee) error {

	hash, err := bcrypt.GenerateFromPassword([]byte(signup.Pass_word), 14)
	if err != nil {
		return errors.New("bcrypt failed:" + err.Error())
	}

	signup.Pass_word = string(hash)
	return u.employeeRepo.AddEmployee(cntxt, signup)
}

func (u *employeeUseCase) Login(r context.Context, login domain.Employee) (domain.Employee, error) {
	employee, err := u.employeeRepo.FindEmployee(r, login)
	fmt.Println(employee)
	fmt.Println(err)

	if err != nil {
		//fmt.Println("Hello")
		return login, errors.New(err.Error())
	}

	if err := bcrypt.CompareHashAndPassword([]byte(employee.Pass_word), []byte(login.Pass_word)); err != nil {

		return login, errors.New("incorrect password")
	}

	return employee, nil
}

func (u *employeeUseCase) SignUpOtp(r context.Context, find domain.Employee) error {
	_, err := u.employeeRepo.FindEmployee(r, find)

	return err
}

func (u *employeeUseCase) AddForm(r context.Context, form domain.Form) error {

	if err := u.employeeRepo.CheckFormDetails(r, form); err != nil {
		return err
	}

	form.Account_no = base64.StdEncoding.EncodeToString(encrypt.Encrypt([]byte(form.Account_no)))
	fmt.Println([]byte(base64.StdEncoding.EncodeToString(encrypt.Encrypt([]byte(form.Pan_number)))))
	form.Pan_number = base64.StdEncoding.EncodeToString(encrypt.Encrypt([]byte(form.Pan_number)))
	form.Adhaar_no = base64.StdEncoding.EncodeToString(encrypt.Encrypt([]byte(form.Adhaar_no)))
	form.Ifsc_code = base64.StdEncoding.EncodeToString(encrypt.Encrypt([]byte(form.Ifsc_code)))
	form.Name_as_per_passbokk = base64.StdEncoding.EncodeToString(encrypt.Encrypt([]byte(form.Name_as_per_passbokk)))
	form.Status = "P"

	if err := u.employeeRepo.AddForm(r, form); err != nil {
		return err
	}

	return nil
}
