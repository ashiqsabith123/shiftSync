package repository

import (
	"context"
	"encoding/base64"
	"errors"
	"shiftsync/pkg/domain"
	"shiftsync/pkg/encrypt"
	"shiftsync/pkg/helper/response"
	repo "shiftsync/pkg/repository/interfaces"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type employeeDatabase struct {
	DB *gorm.DB
}

func NewEmployeeRepository(DB *gorm.DB) repo.EmployeeRepository {
	return &employeeDatabase{DB: DB}
}

func (e *employeeDatabase) AddEmployee(cntxt context.Context, signup domain.Employee) error {
	err := e.DB.Create(&signup).Error
	return err
}

func (e *employeeDatabase) FindEmployee(cntxt context.Context, find domain.Employee) (domain.Employee, error) {
	var emp domain.Employee

	if err := e.DB.Where("id= ? OR email = ? OR phone = ? OR user_name = ?", find.ID, find.Email, find.Phone, find.User_name).First(&emp).Error; err != nil {

		return find, errors.New("no user found")
	}

	return emp, nil
}

func (e *employeeDatabase) CheckFormDetails(cntxt context.Context, form domain.Form) error {
	if err := e.DB.Where("form_id = ? OR account_no = ? OR pan_number = ? OR adhaar_no = ?", form.FormID, base64.StdEncoding.EncodeToString(encrypt.Encrypt([]byte(form.Account_no))), base64.StdEncoding.EncodeToString(encrypt.Encrypt([]byte(form.Pan_number))), base64.StdEncoding.EncodeToString(encrypt.Encrypt([]byte(form.Adhaar_no)))).First(&domain.Form{}).Error; err != nil {

		return nil
	}

	return errors.New("details alredy found")
}

func (e *employeeDatabase) AddForm(cntxt context.Context, form domain.Form) error {

	if err := e.DB.Create(&form).Error; err != nil {
		return err
	}

	return nil
}

func (e *employeeDatabase) FormStatus(ctx context.Context, empID int) string {
	var status string
	if err := e.DB.Raw("select status from forms where form_id =? ", empID).Scan(&status).Error; err != nil {
		return "Error"
	}

	return status
}

func (e *employeeDatabase) GetDutySchedules(ctx context.Context, id int) (response.Duty, error) {
	var duty domain.Attendance
	var schedule response.Duty

	if err := e.DB.Where("employee_id = ? AND status='S'", id).First(&duty).Error; err != nil {
		return schedule, err
	}

	copier.Copy(&schedule, &duty)

	return schedule, nil
}
