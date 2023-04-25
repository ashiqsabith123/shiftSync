package handler

import (
	"errors"
	"net/http"
	"shiftsync/pkg/auth"
	"shiftsync/pkg/domain"
	"shiftsync/pkg/helper/request"
	"shiftsync/pkg/helper/response"
	service "shiftsync/pkg/usecases/interfaces"

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

	resp := response.SuccessResponse(200, "Welcome to signup page", domain.Employee{})

	ctxt.JSON(200, resp)
}

func (u *EmployeeHandler) PostSignup(ctxt *gin.Context) {

	var signup domain.Employee

	if err := ctxt.ShouldBindJSON(&signup); err != nil {

		resp := response.ErrorResponse(400, "Invalid input", err.Error(), nil)
		ctxt.JSON(http.StatusBadRequest, resp)
		return

	} else {
		if err := u.employeeUseCase.SignUp(ctxt, signup); err != nil {
			resp := response.ErrorResponse(400, "Invalid input", err.Error(), nil)
			ctxt.JSON(http.StatusBadRequest, resp)
			return
		} else {
			resp := response.SuccessResponse(200, "Successfully Account Created", nil)
			ctxt.JSON(200, resp)
		}
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

	token, err := auth.GenerateTokens(emp.Signup_id)

	if err != nil {
		resp := response.ErrorResponse(500, "unable to login", err.Error(), nil)
		ctxt.JSON(http.StatusInternalServerError, resp)
		return
	}

	ctxt.SetCookie("user-cookie", token, 20*60, "", "", false, true)
	resp := response.ErrorResponse(200, "succesfuly logged in", "", token)
	ctxt.JSON(200, resp)
}
