package repository

import (
	"context"
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

func (e *employeeDatabase) AddEmployee(cntxt context.Context, signup domain.Employee_Signup) error {
	err := e.DB.Create(&signup).Error
	return err
}
