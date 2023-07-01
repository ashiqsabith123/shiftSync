package handler

import (
	"errors"
	"net/http"
	"shiftsync/pkg/auth"
	"shiftsync/pkg/domain"
	"shiftsync/pkg/helper/request"
	"shiftsync/pkg/helper/response"
	service "shiftsync/pkg/usecases/interfaces"
	"shiftsync/pkg/verification"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type EmployeeHandler struct {
	employeeUseCase service.EmployeeUseCase
}

func NewEmployeeHandler(userUseCase service.EmployeeUseCase) *EmployeeHandler {
	return &EmployeeHandler{employeeUseCase: userUseCase}
}

// Sign Up Page godoc
// @summary Api for get signup page ///
// @id Signup
// @description api for employees to signup
// @tags Employee - Sign up
// @Produce json
// @Router /employee/signup [get]
// @Success 200 {object} request.SignUp{} "Welcome to signup page"
func (u *EmployeeHandler) GetSignUp(ctxt *gin.Context) {

	resp := response.SuccessResponse(200, "Welcome to signup page", request.SignUp{})

	ctxt.JSON(200, resp)
}

// Sign Up Page godoc
// @summary Api for post signup details
// @id Signup
// @description api for employees to signup
// @tags Employee - Sign up
// @Produce json
// @Param input body domain.Employee{} true "Sign up details"
// @Router /employee/signup [post]
// @Success 200 {object} response.Response{} "Otp senEmployeed succesfully"
// @Failure 400 {object} response.Response{} "User already exist"
// @Failure 500 {object} response.Response{} "unable to signup"
func (u *EmployeeHandler) PostSignup(ctxt *gin.Context) {

	var signup domain.Employee

	if err := ctxt.ShouldBindJSON(&signup); err != nil {

		resp := response.ErrorResponse(400, "Invalid input", err.Error(), nil)
		ctxt.JSON(http.StatusBadRequest, resp)
		return

	}

	if err := u.employeeUseCase.SignUpOtp(ctxt, signup); err == nil {
		resp := response.ErrorResponse(400, "User already exist", "", nil)
		ctxt.JSON(http.StatusBadRequest, resp)
		return
	}

	e, b := verification.SendOtp(signup.Phone)

	if b != nil {
		resp := response.ErrorResponse(500, e, b.Error(), nil)
		ctxt.JSON(http.StatusInternalServerError, resp)
		return
	}

	token, err := auth.GenerateTokenForOtp(signup)

	if err != nil {
		resp := response.ErrorResponse(500, "unable to signup", err.Error(), nil)
		ctxt.JSON(http.StatusInternalServerError, resp)
		return
	}

	ctxt.SetCookie("employee", token, 20*60, "", "", false, true)

	resp := response.SuccessResponse(200, "Otp send succesfully", nil)
	ctxt.JSON(200, resp)

}

// Veify otp godoc
// @summary Api for post otp
// @id Verify otp
// @description api for employees to verify otp
// @tags Employee - Sign up
// @Produce json
// @Param input body request.OTPStruct{} true "Otp"
// @Router /employee/signup/verify-otp [post]
// @Success 200 {object} response.Response{} "Successfully Account Created"
// @Failure 400 {object} response.Response{} "Invalid otp"
// @Failure 500 {object} response.Response{} "Unable to find details""
func (u *EmployeeHandler) VerifyOtp(ctxt *gin.Context) {

	var otp request.OTPStruct

	if err := ctxt.ShouldBindJSON(&otp); err != nil {

		resp := response.ErrorResponse(400, "Invalid input", err.Error(), nil)
		ctxt.JSON(http.StatusBadRequest, resp)
		return

	}

	value, err := ctxt.Cookie("employee")
	ctxt.SetCookie("employee", "", -1, "", "", false, true)

	if err != nil {
		resp := response.ErrorResponse(500, "unable to find details", err.Error(), nil)
		ctxt.JSON(http.StatusInternalServerError, resp)
		return
	}

	details, ver := auth.ValidateOtpTokens(value)

	if ver != nil {
		resp := response.ErrorResponse(500, "unable to find details", err.Error(), nil)
		ctxt.JSON(http.StatusInternalServerError, resp)
		return
	}

	t := verification.ValidateOtp(details.Phone, otp.Code)

	if t != nil {
		resp := response.ErrorResponse(400, "Invalid otp", t.Error(), nil)
		ctxt.JSON(http.StatusBadRequest, resp)
		return
	}

	var signup domain.Employee
	copier.Copy(&signup, &details)

	if err := u.employeeUseCase.SignUp(ctxt, signup); err != nil {
		resp := response.ErrorResponse(400, "Invalid", err.Error(), nil)
		ctxt.JSON(http.StatusBadRequest, resp)
		return
	} else {
		resp := response.SuccessResponse(201, "Successfully Account Created", nil)
		ctxt.JSON(201, resp)
	}

}

// Sign In Page godoc
// @summary Api for get signin page
// @id Sign In
// @description api for employees to signin
// @tags Employee - Sign in
// @Produce json
// @Router /employee/signin [get]
// @Success 200 {object} request.LoginStruct{} "Welcome to sign in page"
func (u *EmployeeHandler) GetLogin(ctxt *gin.Context) {
	resp := response.SuccessResponse(200, "Welcome to sign in page", request.LoginStruct{})
	ctxt.JSON(200, resp)
}

// Sign In Page godoc
// @summary Api for post signin details
// @id Sign In
// @description api for employees to signin
// @tags Employee - Sign in
// @Produce json
// @Param input body request.LoginStruct{} true "Sign in details"
// @Router /employee/signin [post]
// @Success 200 {object} request.LoginStruct{} "Succesfuly logged in"
// @Failure 400 {object} response.Response{} "Username and password is mandatory"
// @Failure 500 {object} response.Response{} "Unable to login"
func (u *EmployeeHandler) PostLogin(ctxt *gin.Context) {
	var values request.LoginStruct

	if err := ctxt.ShouldBindJSON(&values); err != nil {
		resp := response.ErrorResponse(400, "Invalid input", err.Error(), request.LoginStruct{})
		ctxt.JSON(400, resp)
		return
	}

	if values.User_name == "" || values.Pass_word == "" {
		err := errors.New("missing credentials")
		resp := response.ErrorResponse(400, "Username and password is mandatory", err.Error(), request.LoginStruct{})
		ctxt.JSON(400, resp)
		return
	}

	var login domain.Employee

	copier.Copy(&login, &values)

	emp, err := u.employeeUseCase.Login(ctxt, login)

	if err != nil {
		resp := response.ErrorResponse(400, "Login failed", err.Error(), nil)
		ctxt.JSON(400, resp)
		return
	}

	token, err := auth.GenerateTokens(emp.ID)

	if err != nil {
		resp := response.ErrorResponse(500, "unable to login", err.Error(), nil)
		ctxt.JSON(http.StatusInternalServerError, resp)
		return
	}

	ctxt.SetCookie("employee-cookie", token, 30*24*60, "/", "", false, true)
	resp := response.SuccessResponse(200, "succesfuly logged in", token)
	ctxt.JSON(200, resp)
}

func (u *EmployeeHandler) Logout(ctx *gin.Context) {
	ctx.SetCookie("employee-cookie", "", -1, "", "", false, true)
	response := response.SuccessResponse(200, "successfully logged out", nil)
	ctx.JSON(http.StatusOK, response)
}

// Form page godoc
// @summary Api for get form
// @id Get Form
// @description api for enter all details of employee
// @tags Employee - Form
// @Produce json
// @Router /employee/form [get]
// @Success 200 {object} request.Form{} "Fill the form"
func (u *EmployeeHandler) GetForm(ctxt *gin.Context) {
	resp := response.SuccessResponse(200, "Fill the form", request.Form{})
	ctxt.JSON(200, resp)
}

// Form page godoc
// @summary Api for post form details
// @id Post form
// @description api for employees to post their form details
// @tags Employee - Form
// @Produce json
// @Router /employee/form [post]
// @Success 200 {object} request.Form{} "Form submitted succesfully pending for verification"
// @Failure 400 {object} response.Response{} "Invalid input"
// @Failure 500 {object} response.Response{} "Employee id not found in cookie"
func (u *EmployeeHandler) PostForm(ctxt *gin.Context) {

	var tempForm request.Form

	empID, ok := ctxt.Get("userId")

	if !ok || empID == "" {
		resp := response.ErrorResponse(500, "Employee id not found cookie", "", nil)
		ctxt.JSON(http.StatusInternalServerError, resp)
	}

	if err := ctxt.ShouldBindJSON(&tempForm); err != nil {
		resp := response.ErrorResponse(400, "Invalid input", err.Error(), tempForm)
		ctxt.JSON(400, resp)
		return
	}

	var form domain.Form

	tempid, _ := strconv.Atoi(empID.(string))

	form.FormID = tempid
	copier.Copy(&form, &tempForm)

	if err := u.employeeUseCase.AddForm(ctxt, form); err != nil {
		resp := response.ErrorResponse(400, "Eroor when adding form", err.Error(), tempForm)
		ctxt.JSON(400, resp)
		return
	}

	resp := response.SuccessResponse(200, "Form submitted succesfully pending for verification", nil)
	ctxt.JSON(200, resp)

}

// Dashboard godoc
// @summary Api for get dashboard
// @id Dashboard
// @description api for get employee dashboard
// @tags Employee - Dashboard
// @Produce json
// @Router /employee/dashboard [get]
// @Success 200 {object} response.Response{} "Succesfully get dashboard"
// @Failure 400 {object} response.Response{} "Eror while getting dashboard"
// @Failure 500 {object} response.Response{} "Employee id not found in cookie"
func (u *EmployeeHandler) GetDashboard(ctx *gin.Context) {

	tempID, ok := ctx.Get("userId")

	if !ok || tempID == "" {
		resp := response.ErrorResponse(500, "Employee id not found in cookie", "", nil)
		ctx.JSON(http.StatusInternalServerError, resp)
	}

	empId, _ := strconv.Atoi(tempID.(string))

	status, err := u.employeeUseCase.FormStatus(ctx, empId)

	if err != nil {
		resp := response.ErrorResponse(400, "Error", err.Error(), "")
		ctx.JSON(400, resp)
		return
	}

	resp := response.SuccessResponse(200, status.Status, status.Correction)
	ctx.JSON(200, resp)

}

// Punching godoc
// @id Get duty
// @summary Api for employees to get duty
// @description Api for employees to get duty schedule
// @tags Employee - Punching
// @Produce json
// @Router /employee/duty [get]
// @Success 200 {object} response.Duty{} "duty schedules"
// @Failure 404 {object} response.Response{} "failed to get duty schedules"
// @Failure 500 {object} response.Response{} "employee id not found"
func (e *EmployeeHandler) GetDuty(c *gin.Context) {

	tempid, ok := c.Get("userId")
	id, _ := strconv.Atoi(tempid.(string))

	if !ok {
		resp := response.ErrorResponse(http.StatusInternalServerError, "employee id not found", "", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	duty, err := e.employeeUseCase.GetDutySchedules(c, id)

	if err != nil {
		resp := response.ErrorResponse(404, "failed to get duty schedules", err.Error(), nil)
		c.JSON(404, resp)
		return
	}

	resp := response.SuccessResponse(200, "duty schedules", duty)
	c.JSON(200, resp)

}

// Punching godoc
// @id Punchin
// @summary Api for employees to punchin
// @description Api for employees to punch in it send an otp to the verified phone number
// @tags Employee - Punching
// @Produce json
// @Router /employee/duty/punchin [get]
// @Success 200 {object} response.Response{} "Otp send to your verified phone number"
// @Failure 404 {object} response.Response{} "failed to get duty schedules"
// @Failure 500 {object} response.Response{} "employee id not found"
// @Failure 400 {object} response.Response{} "error while punching"
func (e *EmployeeHandler) PunchIn(c *gin.Context) {

	tempid, ok := c.Get("userId")
	id, _ := strconv.Atoi(tempid.(string))

	if !ok {
		resp := response.ErrorResponse(http.StatusInternalServerError, "employee id not found", "", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	_, err := e.employeeUseCase.GetDutySchedules(c, id)

	if err != nil {
		resp := response.ErrorResponse(404, "failed to get duty schedules", err.Error(), nil)
		c.JSON(404, resp)
		return
	}

	status, err := e.employeeUseCase.PunchIn(c, id)

	if err != nil {
		resp := response.ErrorResponse(400, status, err.Error(), nil)
		c.JSON(400, resp)
		return
	}

	resp := response.SuccessResponse(200, "Otp send to your verified phone number", nil)
	c.JSON(200, resp)

}

// Punching godoc
// @id Verify otp punchin
// @summary Api for verify otp for punchin
// @description Api for employees to enter the otp to punchin
// @tags Employee - Punching
// @Produce json
// @Router /employee/duty/punchin [post]
// @Success 200 {object} response.Response{} "Punched succesfully"
// @Failure 404 {object} response.Response{} "failed to get duty schedules"
// @Failure 500 {object} response.Response{} "employee id not found"
// @Failure 400 {object} response.Response{} "error while punching"
func (e *EmployeeHandler) VerifyOtpPunchin(c *gin.Context) {

	var otp request.OTPStruct
	if err := c.ShouldBindJSON(&otp); err != nil {
		resp := response.ErrorResponse(400, "Invalid input", err.Error(), otp)
		c.JSON(400, resp)
		return
	}

	tempid, ok := c.Get("userId")
	id, _ := strconv.Atoi(tempid.(string))

	if !ok {
		resp := response.ErrorResponse(http.StatusInternalServerError, "employee id not found", "", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	if err := e.employeeUseCase.VerifyOtpForPunchin(c, id, otp); err != nil {
		resp := response.ErrorResponse(http.StatusInternalServerError, "failed", err.Error(), nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp := response.SuccessResponse(200, "Punched succesfully", nil)
	c.JSON(200, resp)

}

// Punching godoc
// @id Punch Out
// @summary Api for punch out
// @description Api for employees to punch out
// @tags Employee - Punching
// @Produce json
// @Router /employee/duty/punchout [get]
// @Success 200 {object} response.Response{} "Punchout succesfully"
// @Failure 404 {object} response.Response{} "failed to get duty schedules"
// @Failure 500 {object} response.Response{} "employee id not found"
// @Failure 400 {object} response.Response{} "error in punch out"
func (e *EmployeeHandler) PunchOut(c *gin.Context) {

	tempid, ok := c.Get("userId")
	id, _ := strconv.Atoi(tempid.(string))

	if !ok {
		resp := response.ErrorResponse(http.StatusInternalServerError, "employee id not found", "", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	punchOutErr := e.employeeUseCase.PunchOut(c, id)

	if punchOutErr != nil {
		resp := response.ErrorResponse(400, "error in punch out", punchOutErr.Error(), nil)
		c.JSON(400, resp)
		return
	}

	resp := response.SuccessResponse(200, "Punchout succesfully", nil)
	c.JSON(200, resp)

}

// Apply Leave godoc
// @id Apply leave
// @summary Api for employees to apply leave
// @description Api for employees to apply leave
// @tags Employee - Leave
// @Produce json
// @Param input body request.Leave{} true "Leave request details"
// @Router /employee/leave/apply [post]
// @Success 200 {object} response.Response{} "successfully applied for leave"
// @Failure 400 {object} response.Response{} "invalid input"
// @Failure 500 {object} response.Response{} "employee id not found"
// @Failure 422 {object} response.Response{} "error while applying leave"
func (e *EmployeeHandler) ApplyLeave(c *gin.Context) {
	var reqLeave request.Leave

	if err := c.ShouldBindJSON(&reqLeave); err != nil {
		resp := response.ErrorResponse(400, "Invalid input", err.Error(), reqLeave)
		c.JSON(400, resp)
		return
	}

	tempid, ok := c.Get("userId")

	if !ok {
		resp := response.ErrorResponse(http.StatusInternalServerError, "employee id not found", "", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	id, _ := strconv.Atoi(tempid.(string))

	var leave domain.Leave

	leave.EmployeeID = uint(id)

	copier.Copy(&leave, &reqLeave)

	responce, leaveErr := e.employeeUseCase.ApplyLeave(c, leave)

	if leaveErr != nil {
		resp := response.ErrorResponse(422, "error while applying leave", leaveErr.Error(), nil)
		c.JSON(422, resp)
		return
	}

	resp := response.SuccessResponse(200, "succesfully applied leave", responce)
	c.JSON(200, resp)

}

// Leave status/history godoc
// @id Leave status
// @Summary Api for check leave status and history
// @description Api for employees to get leaave status/history
// @tags Employee - Leave
// @Produce json
// @Router /employee/leave/status  [get]
// @Success 200 {object} []response.LeaveHistory{} "successfully fetched leave status/history"
// @Failure 404 {object} response.Response{} "no leave history found"
// @Failure 500 {object} response.Response{} "employee id not found"
func (e *EmployeeHandler) LeaveStatus(c *gin.Context) {
	tempid, ok := c.Get("userId")
	id, _ := strconv.Atoi(tempid.(string))

	if !ok {
		resp := response.ErrorResponse(http.StatusInternalServerError, "employee id not found", "", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	leaveHistory, err := e.employeeUseCase.GetLeaveStatusHistory(c, id)

	if err != nil || len(leaveHistory) == 0 {
		resp := response.ErrorResponse(404, "no leave history found", "", nil)
		c.JSON(404, resp)
		return
	}

	c.JSON(200, gin.H{
		"status":               200,
		"Leave Status/History": leaveHistory,
	})

}

// Employee attendance godoc
// @summary Api for fetch attendance details of an employee
// @id Attendance
// @description api for get employees attendances
// @tags Employee - Attendance
// @Produce json
// @Router /employee/attendance [get]
// @Success 200 {object} []response.Attendance{} "successfully fetched attendance"
// @Failure 404 {object} response.Response{} "failed to get attendance"
// @Failure 500 {object} response.Response{} "employee id not found"
func (e *EmployeeHandler) Attendance(c *gin.Context) {
	tempid, ok := c.Get("userId")
	id, _ := strconv.Atoi(tempid.(string))

	if !ok {
		resp := response.ErrorResponse(http.StatusInternalServerError, "employee id not found", "", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	attendance, err := e.employeeUseCase.Attendance(c, id)

	if err != nil {
		resp := response.ErrorResponse(404, "failed to get attendance", err.Error(), nil)
		c.JSON(404, resp)
		return
	}

	if err != nil || len(attendance) == 0 {
		resp := response.ErrorResponse(404, "no attendace history found", "", nil)
		c.JSON(404, resp)
		return
	}

	c.JSON(200, gin.H{
		"status":     200,
		"Attendance": attendance,
	})

}

// Employee Salary godoc
// @summary Api for fetch salary history of an employee
// @id Salary History
// @description api for get employees salary history
// @tags Employee - Salary
// @Produce json
// @Router /employee/salary/history [get]
// @Success 200 {object} []response.Salaryhistory "successfully fetched salary history"
// @Failure 404 {object} response.Response{} "failed to get attendance"
// @Failure 500 {object} response.Response{} "employee id not found"
func (e *EmployeeHandler) TransactionHistory(c *gin.Context) {
	tempid, ok := c.Get("userId")
	id, _ := strconv.Atoi(tempid.(string))

	if !ok {
		resp := response.ErrorResponse(http.StatusInternalServerError, "employee id not found", "", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	history, err := e.employeeUseCase.GetSalaryHistory(c, id)

	if err != nil {
		resp := response.ErrorResponse(404, "failed to get salary history", err.Error(), nil)
		c.JSON(404, resp)
		return
	}

	if err != nil || len(history) == 0 {
		resp := response.ErrorResponse(404, "no salary history found", "", nil)
		c.JSON(404, resp)
		return
	}

	c.JSON(200, gin.H{
		"status":         200,
		"Salary History": history,
	})

}

// Employee Salary godoc
// @summary Api for fetch salary details of an employee
// @id Salary Details
// @description api for get employees salary details
// @tags Employee - Salary
// @Produce json
// @Router /employee/salary/details [get]
// @Success 200 {object} response.Salarydetails "successfully fetched salary details"
// @Failure 404 {object} response.Response{} "ailed to get salary details"
// @Failure 500 {object} response.Response{} "employee id not found"
func (u *EmployeeHandler) SalaryDetails(c *gin.Context) {
	tempid, ok := c.Get("userId")
	id, _ := strconv.Atoi(tempid.(string))

	if !ok {
		resp := response.ErrorResponse(http.StatusInternalServerError, "employee id not found", "", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	details, err := u.employeeUseCase.GetSalaryDetails(c, id)

	if err != nil {
		resp := response.ErrorResponse(404, "failed to get salary details", err.Error(), nil)
		c.JSON(404, resp)
		return
	}

	if err != nil || details.Gross_salary == 0 {
		resp := response.ErrorResponse(404, "no attendace history found", "", nil)
		c.JSON(404, resp)
		return
	}

	c.JSON(200, gin.H{
		"status":         200,
		"Salary Details": details,
	})

}

// Employee Salary godoc
// @summary Api for get salary slip of ana employee
// @id Salary Slip
// @description api for get employees salary salary slip
// @tags Employee - Salary
// @Produce json
// @Produces - application/pdf
// @Router /employee/salary/download-slip [get]
// @Failure 404 {object} response.Response{} "ailed to get salary details"
// @Failure 500 {object} response.Response{} "employee id not found"
func (u *EmployeeHandler) SalarySlip(c *gin.Context) {
	tempid, ok := c.Get("userId")
	id, _ := strconv.Atoi(tempid.(string))

	if !ok {
		resp := response.ErrorResponse(http.StatusInternalServerError, "employee id not found", "", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	salarySlip, err := u.employeeUseCase.GetDataForSalarySlip(c, id)
	if err != nil {
		resp := response.ErrorResponse(http.StatusInternalServerError, "error while creating pdf", err.Error(), nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename=file.pdf")

	// Send the PDF data as the response
	c.Data(http.StatusOK, "application/pdf", salarySlip)
}
