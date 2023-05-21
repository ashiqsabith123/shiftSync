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

// -------------------Sign Up-----------------------------//
func (u *EmployeeHandler) GetSignUp(ctxt *gin.Context) {

	resp := response.SuccessResponse(200, "Welcome to signup page", request.SignUp{})

	ctxt.JSON(200, resp)
}

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

//---------------------------------------------------------//

// -------------------Sign In-----------------------------//

func (u *EmployeeHandler) GetLogin(ctxt *gin.Context) {
	resp := response.SuccessResponse(200, "Welcome to login page", request.LoginStruct{})
	ctxt.JSON(200, resp)
}

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

	ctxt.SetCookie("employee-cookie", token, 10*60, "/", "", false, true)
	resp := response.ErrorResponse(200, "succesfuly logged in", "", token)
	ctxt.JSON(200, resp)
}

func (u *EmployeeHandler) GetForm(ctxt *gin.Context) {
	resp := response.SuccessResponse(200, "Fill the form", request.Form{})
	ctxt.JSON(200, resp)
}

func (u *EmployeeHandler) PostForm(ctxt *gin.Context) {

	var tempForm request.Form

	empID, ok := ctxt.Get("userId")

	if !ok || empID == "" {
		resp := response.ErrorResponse(500, "Value not found", "", nil)
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
		resp := response.ErrorResponse(400, "Deatils", err.Error(), tempForm)
		ctxt.JSON(400, resp)
		return
	}

	resp := response.SuccessResponse(200, "Form submitted succesfully pending for verification", nil)
	ctxt.JSON(200, resp)

}

func (u *EmployeeHandler) GetDashboard(ctx *gin.Context) {

	tempID, ok := ctx.Get("userId")

	if !ok || tempID == "" {
		resp := response.ErrorResponse(500, "Value not found", "", nil)
		ctx.JSON(http.StatusInternalServerError, resp)
	}

	empId, _ := strconv.Atoi(tempID.(string))

	status := u.employeeUseCase.FormStatus(ctx, empId)

	resp := response.SuccessResponse(200, status, nil)
	ctx.JSON(200, resp)

}

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

// ApplyLeave godoc
// @summary ApplyLeave
// @description api for employees to apply leave
// @tags leave
// @Produce json
// @Param input body request.Leave{} true "input field"
// @Router /leave/apply [post]
// @Success 200 {object} response.Response{} "successfully applied for leave"
// @Failure 400 {object} response.Response{} "invalid input"

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

	if err := e.employeeUseCase.ApplyLeave(c, leave); err != nil {
		resp := response.ErrorResponse(400, "error while applying leave", err.Error(), nil)
		c.JSON(400, resp)
		return
	}

	resp := response.SuccessResponse(200, "succesfully applied leave", nil)
	c.JSON(200, resp)

}

func (e *EmployeeHandler) LeaveStatus(c *gin.Context) {
	tempid, ok := c.Get("userId")
	id, _ := strconv.Atoi(tempid.(string))

	if !ok {
		resp := response.ErrorResponse(http.StatusInternalServerError, "employee id not found", "", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	status, err := e.employeeUseCase.GetLeaveStatusHistory(c, id)

	if err != nil {
		resp := response.ErrorResponse(404, "failed to get leave history", err.Error(), nil)
		c.JSON(404, resp)
		return
	}

	c.JSON(200, gin.H{
		"status":    200,
		"employees": status,
	})

}

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

	c.JSON(200, gin.H{
		"status":    200,
		"employees": attendance,
	})

}
