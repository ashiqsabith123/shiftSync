package usecases

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"shiftsync/pdf"
	"shiftsync/pkg/domain"
	"shiftsync/pkg/encrypt"
	"shiftsync/pkg/helper"
	"shiftsync/pkg/helper/request"
	"shiftsync/pkg/helper/response"
	repo "shiftsync/pkg/repository/interfaces"
	service "shiftsync/pkg/usecases/interfaces"
	"shiftsync/pkg/verification"
	"strconv"
	"time"

	"github.com/jinzhu/copier"
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

func (u *employeeUseCase) AddForm(ctx context.Context, form domain.Form) error {

	form.Account_no = base64.StdEncoding.EncodeToString(encrypt.Encrypt([]byte(form.Account_no)))
	form.Pan_number = base64.StdEncoding.EncodeToString(encrypt.Encrypt([]byte(form.Pan_number)))
	form.Adhaar_no = base64.StdEncoding.EncodeToString(encrypt.Encrypt([]byte(form.Adhaar_no)))
	form.Ifsc_code = base64.StdEncoding.EncodeToString(encrypt.Encrypt([]byte(form.Ifsc_code)))
	form.Name_as_per_passbokk = base64.StdEncoding.EncodeToString(encrypt.Encrypt([]byte(form.Name_as_per_passbokk)))
	form.Status = "P"

	details, ok := u.employeeRepo.CheckFormDetails(ctx, form)

	if ok && details.Status == "A" || details.Status == "P" {
		return errors.New("details already found")
	} else if ok && details.Status == "C" {
		if err := u.employeeRepo.FormCorrection(ctx, form); err != nil {
			return err
		}
	} else {
		if err := u.employeeRepo.AddForm(ctx, form); err != nil {
			return err
		}
	}

	return nil
}

func (u *employeeUseCase) FormStatus(ctx context.Context, empID int) (response.FormStatus, error) {

	status, err := u.employeeRepo.FormStatus(ctx, empID)
	if err != nil {
		return status, err
	}

	switch status.Status {

	case "P":
		status.Status = "Pending for verification"
	case "C":
		status.Status = "Admin requested for correction"
	case "A":
		status.Status = "Welcome to dashboard"

	}
	return status, err
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

func (u *employeeUseCase) PunchIn(ctx context.Context, ID int) (string, error) {
	var find domain.Employee

	find.ID = uint(ID)

	details, _ := u.employeeRepo.FindEmployee(ctx, find)

	status, err := verification.SendOtp(details.Phone)

	return status, err

}

func (u *employeeUseCase) VerifyOtpForPunchin(ctx context.Context, id int, otp request.OTPStruct) error {

	var find domain.Employee

	find.ID = uint(id)

	details, _ := u.employeeRepo.FindEmployee(ctx, find)

	if err := verification.ValidateOtp(details.Phone, otp.Code); err != nil {
		return err
	}

	var punch domain.Attendance

	punch.EmployeeID = details.ID
	punch.Date = time.Now().Format("2006-01-02")
	punch.Punch_in = time.Now().Format("15:04:05")

	if err := u.employeeRepo.PunchIn(ctx, punch); err != nil {
		return err
	}

	return nil

}

func (e *employeeUseCase) PunchOut(ctx context.Context, id int) error {

	var duty domain.Attendance

	duty.EmployeeID = uint(id)
	duty.Punch_out = time.Now().Format("15:04:05")
	if err := e.employeeRepo.PunchOut(ctx, duty); err != nil {
		return err
	}

	return nil
}

func (e *employeeUseCase) ApplyLeave(ctx context.Context, leave domain.Leave) (string, error) {

	var count response.LeaveCount

	count.Date = time.Now().Year()
	count.Id = int(leave.EmployeeID)

	checkCount, countErr := e.employeeRepo.GetCountOfLeaveTaken(ctx, count)
	if countErr != nil {
		return "", countErr
	}

	if checkCount >= 100 {
		return "", errors.New("no leaves available")
	}

	date1, _ := time.Parse("02-01-2006", leave.From)
	date2, _ := time.Parse("02-01-2006", leave.To)

	duration := date2.Sub(date1)
	days := int(duration.Hours() / 24)

	if days > (100 - checkCount) {
		return "", errors.New("you to date exceeds the limit of leaves")
	}

	checkApplied, checkErr := e.employeeRepo.CheckLeaveApplied(ctx, leave)
	if checkErr != nil {
		return "", checkErr
	}

	if checkApplied.From != "" {
		oldDate1, _ := time.Parse("02-01-2006", checkApplied.From)
		oldDate2, _ := time.Parse("02-01-2006", checkApplied.To)

		newDate1, _ := time.Parse("02-01-2006", leave.From)
		newDate2, _ := time.Parse("02-01-2006", leave.To)

		fmt.Println("old", oldDate1.Unix())
		fmt.Println("new", newDate1.Unix())

		if (newDate1.Unix() >= oldDate1.Unix() && newDate1.Unix() <= oldDate2.Unix()) ||
			(newDate2.Unix() >= oldDate1.Unix() && newDate2.Unix() <= oldDate2.Unix()) {
			return "", errors.New("you already applied leave on these dates")
		}
	}

	if checkCount > 50 {
		leave.Mode = "U"
	} else {
		leave.Mode = "P"
	}

	leave.Status = "R"

	if err := e.employeeRepo.ApplyLeave(ctx, leave); err != nil {
		return "", err
	}
	str := strconv.Itoa(100 - checkCount)

	return "available leaves:" + str, nil
}

func (e *employeeUseCase) GetLeaveStatusHistory(ctx context.Context, id int) ([]response.LeaveHistory, error) {
	status, err := e.employeeRepo.LeaveStatusHistory(ctx, id)

	if err != nil {
		return status, err
	}

	for i, _ := range status {
		switch status[i].Status {
		case "A":
			status[i].Status = "Approved"
		case "P":
			status[i].Status = "Pending"
		case "D":
			status[i].Status = "Approved"
		}
	}

	fmt.Println("use", status)

	return status, nil

}

func (e *employeeUseCase) Attendance(ctx context.Context, id int) ([]response.Attendance, error) {
	attnedance, err := e.employeeRepo.Attendance(ctx, id)

	if err != nil {
		return []response.Attendance{}, err
	}

	for i := range attnedance {
		t1, _ := time.Parse("15:04:05", attnedance[i].Punch_in)
		t2, _ := time.Parse("15:04:05", attnedance[i].Punch_out)

		duration := t1.Sub(t2)
		fmt.Println(duration)
		hours := int(duration.Hours())

		attnedance[i].Total_hours = hours

		if hours > 8 {
			attnedance[i].Over_time = hours - 8
		} else {
			attnedance[i].Over_time = 0
		}

		switch attnedance[i].Duty_type {
		case "M":
			attnedance[i].Duty_type = "Morning Duty"

		case "E":
			attnedance[i].Duty_type = "Evening duty"

		case "N":
			attnedance[i].Duty_type = "Night Duty"

		}
	}
	return attnedance, nil
}

func (e *employeeUseCase) GetSalaryHistory(ctx context.Context, id int) ([]response.Salaryhistory, error) {
	var history []response.Salaryhistory

	History, err := e.employeeRepo.GetSalaryHistory(ctx, id)

	if err != nil {
		return history, err
	}

	copier.Copy(&history, &History)

	fmt.Println(history)

	for i := range History {

		history[i].Date = History[i].Date.Format("02-01-2006")

		history[i].Time = History[i].Date.Format("15:04:05")

	}

	return history, nil
}

func (e *employeeUseCase) GetSalaryDetails(ctx context.Context, id int) (response.Salarydetails, error) {
	details, err := e.employeeRepo.GetSalaryDetails(ctx, id)

	if err != nil {
		return details, err
	}

	return details, nil
}

func (e *employeeUseCase) GetDataForSalarySlip(ctx context.Context, id int) ([]byte, error) {

	data, getDataErr := e.employeeRepo.GetDataForSalarySlip(ctx, id)
	if getDataErr != nil {
		return nil, getDataErr
	}

	if data.Base_salary == "" {
		return nil, errors.New("you salary details not found contact admin")
	}

	data.Account_no = string(encrypt.Decrypt(helper.Decode(data.Account_no)))

	pdfError := pdf.CreatePdf(data)
	if pdfError != nil {
		return nil, pdfError
	}

	pdfData, pdfPathErr := ioutil.ReadFile("pdf/generated/salary_slip" + data.Employee_id + ".pdf")
	if pdfPathErr != nil {
		return nil, pdfPathErr
	}

	return pdfData, nil

}
