package repository

import (
	"context"
	"errors"
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
