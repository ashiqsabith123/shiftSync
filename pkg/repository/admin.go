package repository

import (
	"context"
	"errors"
	"fmt"
	"shiftsync/pkg/domain"
	"shiftsync/pkg/helper/response"
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

func (a *adminDatabase) GetAllForms(ctx context.Context) ([]response.Form, error) {

	var forms []response.Form
	err := a.DB.Raw("select employees.id, employees.first_name, employees.last_name, employees.email, employees.phone, forms.gender, forms.marital_status, forms.date_of_birth, forms.p_address, forms.c_address, forms.account_no, forms.ifsc_code, forms.name_as_per_passbokk, forms.pan_number, forms.adhaar_no, forms.photo from employees inner join forms on employees.id = forms.form_id where forms.status='P'").Scan(&forms).Error

	return forms, err
}

func (a *adminDatabase) FindFormByID(ctx context.Context, fID int) error {
	var form domain.Form
	fmt.Println(fID)
	if err := a.DB.Where("form_id=?", fID).First(&form).Error; err != nil {
		return errors.New("form not found with given id")
	}
	return nil
}

func (a *adminDatabase) ApproveApplication(ctx context.Context, form domain.Form, id int) {

	err := a.DB.Raw("UPDATE forms SET status='A', employee_id = ? WHERE form_id= ?", id, form.FormID).Scan(&form).Error
	fmt.Println(err)
}

func (a *adminDatabase) FormCorrection(ctx context.Context, form domain.Form) {
	err := a.DB.Raw("UPDATE forms SET correction = ? WHERE form_id = ?", form.Correction, form.FormID).Scan(&form).Error
	fmt.Println(err)
}
