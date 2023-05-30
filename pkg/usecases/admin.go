package usecases

import (
	"context"
	"errors"
	"shiftsync/pkg/domain"
	"shiftsync/pkg/encrypt"
	"shiftsync/pkg/helper"
	"shiftsync/pkg/helper/request"
	"shiftsync/pkg/helper/response"
	repo "shiftsync/pkg/repository/interfaces"
	service "shiftsync/pkg/usecases/interfaces"

	"golang.org/x/crypto/bcrypt"
)

type AdminUseCase struct {
	adminRepo repo.AdminRepository
}

func NewAdminUseCase(adRep repo.AdminRepository) service.AdminUseCase {
	return &AdminUseCase{adminRepo: adRep}
}

func (a *AdminUseCase) SignUp(ctx context.Context, admin domain.Admin) error {
	_, err := a.adminRepo.FindAdmin(ctx, admin)
	if err == nil {
		return errors.New("admin already exist")
	}

	hash, hasherr := bcrypt.GenerateFromPassword([]byte(admin.Pass_word), 14)

	if hasherr != nil {
		return errors.New("hashing failed" + hasherr.Error())
	}

	admin.Pass_word = string(hash)

	if err := a.adminRepo.SaveAdmin(ctx, admin); err != nil {
		return errors.New("unable to add admin " + err.Error())
	}

	return nil
}

func (a *AdminUseCase) SignIn(ctx context.Context, details domain.Admin) (domain.Admin, error) {
	admin, err := a.adminRepo.FindAdmin(ctx, details)
	if err != nil {
		return details, errors.New("invalid credentials " + err.Error())
	}

	if berr := bcrypt.CompareHashAndPassword([]byte(admin.Pass_word), []byte(details.Pass_word)); berr != nil {
		return details, errors.New("incorrect password")
	}

	return admin, nil
}

func (a *AdminUseCase) Applications(ctx context.Context) ([]response.Form, error) {
	forms, err := a.adminRepo.GetAllForms(ctx)

	for i := 0; i < len(forms); i++ {

		forms[i].Account_no = string(encrypt.Decrypt(helper.Decode(forms[i].Account_no)))
		forms[i].Adhaar_no = string(encrypt.Decrypt(helper.Decode(forms[i].Adhaar_no)))
		forms[i].Pan_number = string(encrypt.Decrypt(helper.Decode(forms[i].Pan_number)))
		forms[i].Name_as_per_passbokk = string(encrypt.Decrypt(helper.Decode(forms[i].Name_as_per_passbokk)))
		forms[i].Ifsc_code = string(encrypt.Decrypt(helper.Decode(forms[i].Ifsc_code)))
	}

	if err != nil {
		return []response.Form{}, errors.New("no forms found")
	}

	return forms, nil

}

func (a *AdminUseCase) ApproveApplication(ctx context.Context, form domain.Form, admID int) error {
	if err := a.adminRepo.FindFormByID(ctx, form.FormID); err != nil {
		return err
	}

	empid := helper.CreateId()

	a.adminRepo.ApproveApplication(ctx, form, empid, admID)

	return nil
}

func (a *AdminUseCase) FormCorrection(ctx context.Context, form domain.Form) error {
	if err := a.adminRepo.FindFormByID(ctx, form.Employee_id); err != nil {
		return err
	}

	a.adminRepo.FormCorrection(ctx, form)

	return nil
}

func (a *AdminUseCase) GetAllEmployeesSchedules(ctx context.Context) ([]response.Schedule, error) {
	data, err := a.adminRepo.GetAllEmployeesSchedules(ctx)
	return data, err
}

func (a *AdminUseCase) ScheduleDuty(ctx context.Context, duty domain.Duty) error {

	if err := a.adminRepo.ScheduleDuty(ctx, duty); err != nil {
		return err
	}

	return nil
}

func (a *AdminUseCase) GetLeaveRequests(ctx context.Context) ([]response.LeaveRequests, error) {
	leaveRquests, err := a.adminRepo.GetLeaveRequests(ctx)

	if err != nil {
		return leaveRquests, err
	}

	return leaveRquests, nil
}

func (a *AdminUseCase) ApproveLeaveRequests(ctx context.Context, id int) error {
	var status request.LeaveStatus

	status.Id = id
	status.Status = "A"

	if err := a.adminRepo.ChangeLeaveStatus(ctx, status); err != nil {
		return err
	}

	return nil
}

func (a *AdminUseCase) DeclineLeaveRequests(ctx context.Context, id int) error {
	var status request.LeaveStatus

	status.Id = id
	status.Status = "D"

	if err := a.adminRepo.ChangeLeaveStatus(ctx, status); err != nil {
		return err
	}

	return nil
}

func (a *AdminUseCase) AddSalaryDetails(ctx context.Context, salaryDetails domain.Salary) error {

	if err := a.adminRepo.AddSalaryDetails(ctx, salaryDetails); err != nil {
		return err
	}

	return nil
}

func (a *AdminUseCase) EditSalaryDetails(ctx context.Context, editDetails domain.Salary) error {
	if err := a.adminRepo.EditSalaryDetails(ctx, editDetails); err != nil {
		return err
	}

	return nil
}

func (a *AdminUseCase) FindEmployeeById(ctx context.Context, id int) response.EmployeeDetails {
	details := a.adminRepo.FindEmployeeById(ctx, id)
	return details
}

func (a *AdminUseCase) FetchAccountDetailsById(ctx context.Context, id int) response.AccountDetails {
	details := a.adminRepo.FetchAccountDetailsById(ctx, id)

	details.Account_no = string(encrypt.Decrypt(helper.Decode(details.Account_no)))
	details.Ifsc_code = string(encrypt.Decrypt(helper.Decode(details.Ifsc_code)))

	return details
}
