package repository

import (
	"context"
	"errors"
	"shiftsync/pkg/domain"
	repo "shiftsync/pkg/repository/interfaces"

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

func (e *employeeDatabase) AddForm(cntxt context.Context, form domain.Form) error {
	return nil
}

func (e *employeeDatabase) CheckFormDetails(cntxt context.Context, form domain.Form) error {
	if err := e.DB.Where("email = ? OR account_no = ? OR pan_number = ? OR adhaar_no =?", form.Email, form.Account_no, form.Pan_number, form.Adhaar_no).First(&domain.Form{}); err != nil {
		return nil
	}

	return errors.New("details alredy found")
}
