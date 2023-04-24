package handler

import (
	service "shiftsync/pkg/usecases/interfaces"

	"github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
	employeeUseCase service.EmployeeUseCase
}

func NewEmployeeHandler(userUseCase service.EmployeeUseCase) *EmployeeHandler {
	return &EmployeeHandler{employeeUseCase: userUseCase}
}

func (u EmployeeHandler) GetSignUp(ctxt *gin.Context) {

	ctxt.JSON(200, gin.H{
		"Message:": "Welcome to signup page",
	})
}
