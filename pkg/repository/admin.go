package repository

import (
	"context"
	"errors"
	"fmt"
	"shiftsync/pkg/domain"
	"shiftsync/pkg/helper/request"
	"shiftsync/pkg/helper/response"
	repo "shiftsync/pkg/repository/interfaces"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type adminDatabase struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB, repo repo.EmployeeRepository) repo.AdminRepository {

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
	err := a.DB.Raw("select employees.id, employees.first_name, employees.last_name, employees.email, employees.phone, forms.gender, forms.marital_status, forms.date_of_birth, forms.p_address, forms.c_address, forms.account_no, forms.ifsc_code, forms.name_as_per_passbokk, forms.pan_number, forms.adhaar_no, forms.designation,forms.department, forms.photo from employees inner join forms on employees.id = forms.form_id where forms.status='P'").Scan(&forms).Error

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

func (a *adminDatabase) ApproveApplication(ctx context.Context, form domain.Form, id int, empId int) {

	err := a.DB.Raw("UPDATE forms SET status='A', employee_id = ?, approved_by =? WHERE form_id= ?", id, empId, form.FormID).Scan(&form).Error
	fmt.Println(err)
}

func (a *adminDatabase) FormCorrection(ctx context.Context, form domain.Form) {
	err := a.DB.Raw("UPDATE forms SET correction = ? WHERE form_id = ?", form.Correction, form.FormID).Scan(&form).Error
	fmt.Println(err)
}

func (a *adminDatabase) GetAllEmployees(ctx context.Context) ([]response.AllEmployee, error) {
	var emp []response.AllEmployee
	if err := a.DB.Raw("select forms.form_id as id, forms.employee_id as empid, employees.first_name || ' ' || employees.last_name as name, employees.email, employees.phone, forms.gender,forms.date_of_birth, forms.department,forms.designation from employees inner join forms on employees.id = forms.form_id   where forms.status='A' ").Scan(&emp).Error; err != nil {
		return emp, err
	}

	return emp, nil
}

func (a *adminDatabase) GetAllEmployeesSchedules(ctx context.Context) ([]response.Schedule, error) {
	var emp []response.Schedule
	if err := a.DB.Raw("SELECT forms.form_id AS id, employees.first_name || ' ' || employees.last_name AS name, employees.email, employees.phone, forms.designation, duties.status FROM forms INNER JOIN employees ON employees.id = forms.form_id LEFT OUTER JOIN duties ON forms.form_id = duties.employee_id WHERE duties.employee_id IS NULL OR duties.status = 'C';").Scan(&emp).Error; err != nil {
		return emp, err
	}

	return emp, nil
}

func (a *adminDatabase) ScheduleDuty(ctx context.Context, duty domain.Duty) error {

	duty.Status = "S"

	if err := a.DB.Create(&duty).Error; err != nil {
		return err
	}

	return nil
}

func (a *adminDatabase) GetLeaveRequests(ctx context.Context) ([]response.LeaveRequests, error) {
	var leaveRequests []response.LeaveRequests
	if err := a.DB.Raw("SELECT employees.first_name || ' ' || employees.last_name AS name, forms.form_id as id, leaves.leave_type, leaves.from, leaves.to, leaves.reason FROM forms JOIN employees ON employees.id = forms.form_id JOIN leaves ON employees.id = leaves.employee_id WHERE leaves.status = 'R';").Scan(&leaveRequests).Error; err != nil {
		return leaveRequests, err
	}

	return leaveRequests, nil
}

func (a *adminDatabase) ChangeLeaveStatus(ctx context.Context, status request.LeaveStatus) error {
	if err := a.DB.Raw("update leaves set status = ? where employee_id = ?", status.Status, status.Id).Scan(&domain.Leave{}).Error; err != nil {
		return err
	}

	return nil
}

func (a *adminDatabase) AddSalaryDetails(ctx context.Context, salaryDetails domain.Salary) error {
	if err := a.DB.Create(&salaryDetails).Error; err != nil {
		return err
	}

	return nil
}

func (a *adminDatabase) EditSalaryDetails(ctx context.Context, editDetails domain.Salary) error {
	var salaryDetails domain.Salary
	if err := a.DB.Where("employee_id = ?", editDetails.EmployeeID).First(&salaryDetails).Error; err != nil {
		return err
	}

	copier.Copy(&salaryDetails, &editDetails)

	a.DB.Save(&salaryDetails)
	return nil
}

func (e *adminDatabase) CalculateTotalWorkingHours(ctx context.Context, id int, month string) (float32, error) {
	var hour float32
	if err := e.DB.Raw("SELECT CAST(SUM(EXTRACT(epoch FROM (TO_TIMESTAMP(punch_out, 'HH24:MI:SS') - TO_TIMESTAMP(punch_in, 'HH24:MI:SS')))) AS FLOAT) / 3600 AS hours FROM attendances WHERE employee_id = ? AND date_trunc('month', TO_DATE(date, 'YYYY-MM-DD')) = date_trunc('month', TO_DATE(?, 'YYYY-MM-DD'));", id, month).Scan(&hour).Error; err != nil {
		return 0, err
	}

	return hour, nil
}

func (e *adminDatabase) GetGradeOfTheEmployee(ctx context.Context, id int) (string, error) {
	var grade string
	if err := e.DB.Raw("SELECT grade FROM salaries 	WHERE employee_id = ?", id).Scan(&grade).Error; err != nil {
		return "", err
	}

	return grade, nil
}

func (e *adminDatabase) AddAllAlowances(ctx context.Context, id int) (float32, error) {
	var allowance float32
	if err := e.DB.Raw("SELECT (d_allowance + sp_allowance + m_allowance) AS allowance FROM salaries WHERE employee_id = ?", id).Scan(&allowance).Error; err != nil {
		return 0, err
	}

	return allowance, nil
}

func (a *adminDatabase) CaculateDeductions(ctx context.Context, id int) (float32, error) {
	var deductions float32
	if err := a.DB.Raw("SELECT (tax + provident_fund ) AS deductions FROM salaries WHERE employee_id = ?", id).Scan(&deductions).Error; err != nil {
		return 0, err
	}

	return deductions, nil
}

func (a *adminDatabase) UpdateFullSalary(ctx context.Context, id int, gross, net float32) error {
	if err := a.DB.Exec("UPDATE salaries SET gross_salary = ? ,net_salary =? WHERE employee_id = ?", gross, net, id).Error; err != nil {
		return err
	}
	return nil
}

func (a *adminDatabase) FindEmployeeById(ctx context.Context, id int) response.EmployeeDetails {
	var details response.EmployeeDetails
	if err := a.DB.Raw("SELECT first_name || ' ' || last_name AS name, phone, email FROM employees WHERE id = ?", id).Scan(&details).Error; err != nil {
		fmt.Println(err)
	}

	return details
}
