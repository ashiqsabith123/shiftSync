package repository

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
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

	cntxt.Done()
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
	var duty domain.Duty
	var schedule response.Duty

	if err := e.DB.Where("employee_id = ? AND status='S'", id).First(&duty).Error; err != nil {
		return schedule, err
	}

	copier.Copy(&schedule, &duty)

	return schedule, nil
}

func (e *employeeDatabase) PunchIn(ctx context.Context, punchin domain.Attendance) error {

	if err := e.DB.Create(&punchin).Error; err != nil {
		return err
	}

	if err := e.DB.Exec("UPDATE duties SET status = 'W' WHERE employee_id = ?", punchin.EmployeeID).Error; err != nil {
		return err
	}
	return nil
}

func (e *employeeDatabase) PunchOut(ctx context.Context, punchout domain.Attendance) error {
	//var temp domain.Attendance

	if err := e.DB.Exec("UPDATE attendances SET punch_out = ? WHERE employee_id = ? AND created_at = (SELECT created_at FROM attendances WHERE employee_id = ? ORDER BY created_at DESC LIMIT 1)", punchout.Punch_out, punchout.EmployeeID, punchout.EmployeeID).Error; err != nil {
		return err
	}

	if err := e.DB.Exec("UPDATE duties SET status = 'C' WHERE employee_id = ?", punchout.EmployeeID).Error; err != nil {
		return err
	}

	return nil
}

func (e *employeeDatabase) ApplyLeave(ctx context.Context, leave domain.Leave) error {
	if err := e.DB.Create(&leave).Error; err != nil {
		return err
	}

	return nil
}

func (e *employeeDatabase) LeaveStatusHistory(ctx context.Context, id int) ([]response.LeaveHistory, error) {
	var history []response.LeaveHistory
	if err := e.DB.Raw("SELECT leave_type, leaves.to, leaves.from, status FROM leaves WHERE employee_id = ?", id).Scan(&history).Error; err != nil {

		return []response.LeaveHistory{}, err
	}

	fmt.Println("his", history)
	return history, nil
}

func (e *employeeDatabase) Attendance(ctx context.Context, id int) ([]response.Attendance, error) {
	var attendance []response.Attendance
	if err := e.DB.Raw("select date, punch_in , punch_out, duty_type from attendances where status = 'C' and employee_id = ?", id).Scan(&attendance).Error; err != nil {
		return []response.Attendance{}, err
	}

	return attendance, nil
}
