package handler

import (
	"net/http"
	"shiftsync/pkg/auth"
	"shiftsync/pkg/domain"
	"shiftsync/pkg/helper/request"
	"shiftsync/pkg/helper/response"
	service "shiftsync/pkg/usecases/interfaces"

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

	ctxt.SetCookie("admin-cookie", token, 20*60, "", "", false, true)
	resp := response.ErrorResponse(200, "succesfuly logged in", "", token)
	ctxt.JSON(200, resp)

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

	if err != nil {
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
