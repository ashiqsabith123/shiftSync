package handler

import (
	"errors"
	"fmt"
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

	fmt.Println(value)

	details, ver := auth.ValidateOtpTokens(value)
	fmt.Println(details)
	fmt.Println(ver)

	if ver != nil {
		resp := response.ErrorResponse(500, "unable to find details", err.Error(), nil)
		ctxt.JSON(http.StatusInternalServerError, resp)
		return
	}

	t := verification.ValidateOtp(details.Phone, otp.Code)

	if t != nil {
		resp := response.ErrorResponse(400, "Invalid otp", err.Error(), nil)
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

	ctxt.SetCookie("employee-cookie", token, 20*60, "", "", false, true)
	resp := response.ErrorResponse(200, "succesfuly logged in", "", token)
	ctxt.JSON(200, resp)
}

func (u *EmployeeHandler) GetForm(ctxt *gin.Context) {
	resp := response.SuccessResponse(200, "Fill the form", request.Form{})
	ctxt.JSON(200, resp)
}

func (u *EmployeeHandler) PostForm(ctxt *gin.Context) {

	var tempForm request.Form

	value, ok := ctxt.Get("employeeId")

	if !ok || value == "" {
		resp := response.ErrorResponse(500, "Value not found", "", nil)
		ctxt.JSON(http.StatusInternalServerError, resp)
	}

	//fmt.Println(value, ok)

	if err := ctxt.ShouldBindJSON(&tempForm); err != nil {
		resp := response.ErrorResponse(400, "Invalid input", err.Error(), tempForm)
		ctxt.JSON(400, resp)
		return
	}

	var form domain.Form

	tempid, _ := strconv.Atoi(value.(string))

	form.Employee_id = tempid

	copier.Copy(&form, &tempForm)

	// if err:= u.

	if err := u.employeeUseCase.AddForm(ctxt, form); err != nil {
		resp := response.ErrorResponse(400, "Deatils", err.Error(), tempForm)
		ctxt.JSON(400, resp)
		return
	}

	resp := response.SuccessResponse(200, "Form submitted succesfully pending for verification", nil)
	ctxt.JSON(200, resp)

}
