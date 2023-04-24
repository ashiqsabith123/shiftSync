package handler

import (
	"fmt"
	"net/http"
	"shiftsync/pkg/domain"
	"shiftsync/pkg/helper/response"
	service "shiftsync/pkg/usecases/interfaces"

	"github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
	employeeUseCase service.EmployeeUseCase
}

func NewEmployeeHandler(userUseCase service.EmployeeUseCase) *EmployeeHandler {
	return &EmployeeHandler{employeeUseCase: userUseCase}
}

func (u *EmployeeHandler) GetSignUp(ctxt *gin.Context) {

	resp := response.SuccessResponse(200, "Welcome to signup page", domain.Employee_Signup{})

	ctxt.JSON(200, resp)
}

func (u *EmployeeHandler) PostSignup(ctxt *gin.Context) {
	var signup domain.Employee_Signup

	fmt.Println(ctxt.Params)

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
