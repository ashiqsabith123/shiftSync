package usecases

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"shiftsync/pkg/domain"
	"shiftsync/pkg/encrypt"
	"shiftsync/pkg/helper/response"
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

func (u *employeeUseCase) FormStatus(ctx context.Context, empID int) string {

	status := u.employeeRepo.FormStatus(ctx, empID)

	switch status {

	case "P":
		return "Pending for verification"
	case "C":
		return "Admin requested for correction"
	case "A":
		return "Welcome to dashboard"

	}
	return ""
}

func (u *employeeUseCase) GetDutySchedules(ctx context.Context, id int) (response.Duty, error) {
	duty, err := u.employeeRepo.GetDutySchedules(ctx, id)

	if err == nil {
		switch duty.Duty_type {
		case "M":
			duty.Duty_type = "Morning Duty"
			duty.Time = "7:00 AM - 3:00 PM"
			return duty, nil
		case "E":
			duty.Duty_type = "Evening duty"
			duty.Time = "3 AM - 10:00 PM"
			return duty, nil
		case "N":
			duty.Duty_type = "Night Duty"
			duty.Time = "10:00 PM - 5:00 AM "
			return duty, nil
		}

	}

	return duty, err
}
