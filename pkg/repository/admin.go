package repository

import (
	"context"
	"errors"
	"fmt"
	"shiftsync/pkg/domain"
	repo "shiftsync/pkg/repository/interfaces"

	"gorm.io/gorm"
)

type adminDatabase struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) repo.AdminRepository {
	return &adminDatabase{DB: DB}
}

func (a *adminDatabase) FindAdmin(ctx context.Context, find domain.Admin) (domain.Admin, error) {

	var adm domain.Admin
	if err := a.DB.Where("id= ? OR email = ? OR phone = ? OR user_name = ?", find.ID, find.Email, find.Phone, find.User_name).First(&adm).Error; err != nil {
		return find, errors.New("admin not found")
	}

	return adm, nil
}

func (a *adminDatabase) SaveAdmin(ctx context.Context, admin domain.Admin) error {
	if err := a.DB.Create(&admin).Error; err != nil {
		return err
	}

	return nil
}

func (a *adminDatabase) GetAllForms(ctx context.Context) ([]domain.Form, error) {

	var forms []domain.Form
	err := a.DB.Raw("SELECT * FROM forms WHERE status='P'").Scan(&forms).Error

	return forms, err
}

func (a *adminDatabase) FindFormByID(ctx context.Context, fID int) error {
	var form domain.Form
	if err := a.DB.Where("employee_id=?", fID).First(&form).Error; err != nil {
		return errors.New("form not found with given id")
	}
	return nil
}

func (a *adminDatabase) ApproveApplication(ctx context.Context, form domain.Form) {

	err := a.DB.Raw("UPDATE forms SET status='A' WHERE employee_id = ?", form.EmployeeID).Scan(&form).Error
	fmt.Println(err)
}

func (a *adminDatabase) FormCorrection(ctx context.Context, form domain.Form) {
	err := a.DB.Raw("UPDATE forms SET correction = ? WHERE employee_id = ?", form.Correction, form.EmployeeID).Scan(&form).Error
	fmt.Println(err)
}
