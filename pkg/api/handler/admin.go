package handler

import (
	"fmt"
	"net/http"
	"shiftsync/pkg/auth"
	"shiftsync/pkg/domain"
	"shiftsync/pkg/helper/request"
	"shiftsync/pkg/helper/response"
	service "shiftsync/pkg/usecases/interfaces"
	"shiftsync/razorpay"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type AdminHandler struct {
	adminusecase service.AdminUseCase
}

func NewAdminHandler(usecase service.AdminUseCase) *AdminHandler {
	return &AdminHandler{adminusecase: usecase}
}

func (a *AdminHandler) GetSignin(ctxt *gin.Context) {
	resp := response.SuccessResponse(200, "welcome to signin", request.LoginStruct{})
	ctxt.JSON(200, resp)
}

func (a *AdminHandler) PostSignin(ctxt *gin.Context) {
	var adm request.LoginStruct
	if err := ctxt.ShouldBindJSON(&adm); err != nil {
		resp := response.ErrorResponse(400, "Invalid input", err.Error(), request.LoginStruct{})
		ctxt.JSON(400, resp)
		return
	}

	var login domain.Admin

	copier.Copy(&login, &adm)

	admin, err := a.adminusecase.SignIn(ctxt, login)

	if err != nil {
		resp := response.ErrorResponse(400, "Failed to  loginn", err.Error(), request.LoginStruct{})
		ctxt.JSON(400, resp)
		return
	}

	token, terr := auth.GenerateTokens(uint(admin.ID))

	if terr != nil {
		resp := response.ErrorResponse(500, "unable to login", err.Error(), nil)
		ctxt.JSON(http.StatusInternalServerError, resp)
		return
	}

	ctxt.SetCookie("admin-cookie", token, 10*60, "", "", false, true)
	resp := response.ErrorResponse(200, "succesfuly logged in", "", token)
	ctxt.JSON(200, resp)

}

func (u *AdminHandler) Logout(ctx *gin.Context) {
	ctx.SetCookie("admin-cookie", "", -1, "", "", false, true)
	response := response.SuccessResponse(200, "successfully logged out", nil)
	ctx.JSON(http.StatusOK, response)
}

func (a *AdminHandler) GetSignUp(ctxt *gin.Context) {
	resp := response.SuccessResponse(200, "Welcome to signup page", request.SignUp{})
	ctxt.JSON(200, resp)
}

func (a *AdminHandler) PostSignup(ctxt *gin.Context) {

	var rep request.SignUp

	if err := ctxt.ShouldBindJSON(&rep); err != nil {
		resp := response.ErrorResponse(400, "invalid input", err.Error(), rep)
		ctxt.JSON(http.StatusBadRequest, resp)
		return
	}

	var admin domain.Admin

	copier.Copy(&admin, &rep)

	if err := a.adminusecase.SignUp(ctxt, admin); err != nil {
		resp := response.ErrorResponse(400, "invalid", err.Error(), nil)
		ctxt.JSON(400, resp)
		return
	}

	resp := response.SuccessResponse(201, "Suucesfully account created", rep)
	ctxt.JSON(201, resp)

}

func (a *AdminHandler) ViewApplications(ctx *gin.Context) {
	forms, err := a.adminusecase.Applications(ctx)

	if err != nil || len(forms) == 0 {
		ctx.JSON(204, gin.H{
			"status":  204,
			"message": "no new forms found",
		})
		return
	}

	var resform []response.Form
	copier.Copy(&resform, &forms)

	ctx.JSON(200, gin.H{
		"status": 200,
		"forms":  resform,
	})

}

func (a *AdminHandler) ApproveApplication(ctx *gin.Context) {
	var res request.FormApprove
	if err := ctx.ShouldBindJSON(&res); err != nil {
		resp := response.ErrorResponse(400, "invalid input", err.Error(), res)
		ctx.JSON(400, resp)
		return
	}

	var form domain.Form
	copier.Copy(&form, &res)
	ctxempId, _ := ctx.Get("userId")

	empId, _ := strconv.Atoi(ctxempId.(string))

	details := a.adminusecase.FindEmployeeById(ctx, form.FormID)

	contactId, razroErr := razorpay.CreateContact(details)

	if razroErr != nil {
		resp := response.ErrorResponse(400, "error", razroErr.Error(), res)
		ctx.JSON(400, resp)
		return
	}

	accDetails := a.adminusecase.FetchAccountDetailsById(ctx, form.FormID)

	accDetails.Id = form.FormID
	accDetails.Name = details.Name
	accDetails.Contact_id = contactId
	fmt.Println(accDetails)
	if err := razorpay.CreateFundAccount(ctx, accDetails); err != nil {
		resp := response.ErrorResponse(400, "error", err.Error(), res)
		ctx.JSON(400, resp)
		return
	}

	if err := a.adminusecase.ApproveApplication(ctx, form, empId); err != nil {
		resp := response.ErrorResponse(400, "error", err.Error(), res)
		ctx.JSON(400, resp)
		return
	}

	resp := response.SuccessResponse(200, "approved succesfully", nil)
	ctx.JSON(200, resp)
}

func (a *AdminHandler) FormCorrection(ctx *gin.Context) {
	var res request.FormCorrection
	if err := ctx.ShouldBindJSON(&res); err != nil {
		resp := response.ErrorResponse(400, "invalid input", err.Error(), res)
		ctx.JSON(400, resp)
		return
	}

	var form domain.Form

	copier.Copy(&form, &res)

	if err := a.adminusecase.FormCorrection(ctx, form); err != nil {
		resp := response.ErrorResponse(400, "error", err.Error(), res)
		ctx.JSON(400, resp)
		return
	}

	resp := response.SuccessResponse(200, "succesfull send for correction", nil)
	ctx.JSON(200, resp)
}

func (a *AdminHandler) GetScheduleDuty(c *gin.Context) {

	tempData, err := a.adminusecase.GetAllEmployeesSchedules(c)

	if err != nil || len(tempData) == 0 {
		c.JSON(204, gin.H{
			"status":  204,
			"message": "no employees found",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":    200,
		"employees": tempData,
	})

}

func (a *AdminHandler) ScheduleDuty(c *gin.Context) {
	var req request.DutySchedule

	if err := c.ShouldBindJSON(&req); err != nil {
		resp := response.ErrorResponse(400, "invalid input", err.Error(), req)
		c.JSON(400, resp)
		return
	}

	var duty domain.Duty

	copier.Copy(&duty, &req)

	if err := a.adminusecase.ScheduleDuty(c, duty); err != nil {
		resp := response.ErrorResponse(http.StatusInternalServerError, "error", err.Error(), nil)
		c.JSON(500, resp)
		return
	}

	resp := response.SuccessResponse(200, "dutyscheduled succesfully", duty.EmployeeID)
	c.JSON(200, resp)

}

func (a *AdminHandler) GetAllLeaveRequets(c *gin.Context) {
	tempData, err := a.adminusecase.GetLeaveRequests(c)

	if err != nil || len(tempData) == 0 {
		c.JSON(204, gin.H{
			"status":  204,
			"message": "no leave requests found",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":         200,
		"Leave Requests": tempData,
	})
}

func (a *AdminHandler) ApproveLeaveRequests(c *gin.Context) {

	var res request.FormApprove
	if err := c.ShouldBindJSON(&res); err != nil {
		resp := response.ErrorResponse(400, "invalid input", err.Error(), res)
		c.JSON(400, resp)
		return
	}

	if err := a.adminusecase.ApproveLeaveRequests(c, res.FormID); err != nil {
		resp := response.ErrorResponse(400, "failed to approve request", err.Error(), res)
		c.JSON(400, resp)
		return
	}

	resp := response.SuccessResponse(200, "request approved succesfully", "")
	c.JSON(200, resp)

}

func (a *AdminHandler) DeclineLeaveRequests(c *gin.Context) {

	var res request.FormApprove
	if err := c.ShouldBindJSON(&res); err != nil {
		resp := response.ErrorResponse(400, "invalid input", err.Error(), res)
		c.JSON(400, resp)
		return
	}

	if err := a.adminusecase.DeclineLeaveRequests(c, res.FormID); err != nil {
		resp := response.ErrorResponse(400, "failed to decline request", err.Error(), res)
		c.JSON(400, resp)
		return
	}

	resp := response.SuccessResponse(200, "request declined succesfully", "")
	c.JSON(200, resp)

}

func (a *AdminHandler) AddSalaryDetails(c *gin.Context) {
	var salarydetails response.SalaryDetails

	if err := c.ShouldBindJSON(&salarydetails); err != nil {
		resp := response.ErrorResponse(400, "invalid input", err.Error(), salarydetails)
		c.JSON(400, resp)
		return
	}

	var salaryDetails domain.Salary

	copier.Copy(&salaryDetails, &salarydetails)

	if err := a.adminusecase.AddSalaryDetails(c, salaryDetails); err != nil {
		resp := response.ErrorResponse(400, "failed to add salary details", err.Error(), "")
		c.JSON(400, resp)
		return
	}

	resp := response.SuccessResponse(200, "salary details added succesfully", "")
	c.JSON(200, resp)

}

func (a *AdminHandler) EditSalaryDetails(c *gin.Context) {
	var salarydetails response.SalaryDetails

	if err := c.ShouldBindJSON(&salarydetails); err != nil {
		resp := response.ErrorResponse(400, "invalid input", err.Error(), salarydetails)
		c.JSON(400, resp)
		return
	}

	var salaryDetails domain.Salary

	copier.Copy(&salaryDetails, &salarydetails)

	if err := a.adminusecase.EditSalaryDetails(c, salaryDetails); err != nil {
		resp := response.ErrorResponse(400, "failed to edit salary details", err.Error(), "")
		c.JSON(400, resp)
		return
	}

	resp := response.SuccessResponse(200, "salary details edited succesfully", "")
	c.JSON(200, resp)

}

func (a *AdminHandler) GetAllTransactions(c *gin.Context) {
	tempData, err := a.adminusecase.GetAllTransactions(c)

	if err != nil {
		c.JSON(204, gin.H{
			"status":  204,
			"message": "no details",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":       200,
		"Transactions": tempData,
	})
}

func (a *AdminHandler) SalaryNotAdded(c *gin.Context) {
	tempData, err := a.adminusecase.GetEmployeeSalaryNotAdded(c)

	if err != nil {
		c.JSON(204, gin.H{
			"status":  204,
			"message": "no employees to add salaray details",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":    200,
		"Employees": tempData,
	})
}

func (a *AdminHandler) GetAllEmployees(c *gin.Context) {
	tempData, err := a.adminusecase.GetAllEmployees(c)

	if err != nil {
		c.JSON(204, gin.H{
			"status":  204,
			"message": "no employees found",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":    200,
		"Employees": tempData,
	})
}
