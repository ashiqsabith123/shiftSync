package routes

import (
	"shiftsync/pkg/api/handler"
	"shiftsync/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(api *gin.RouterGroup, adminHandler *handler.AdminHandler) {

	signup := api.Group("/signup")
	{
		signup.GET("/", adminHandler.GetSignUp)
		signup.POST("/", adminHandler.PostSignup)
	}

	signin := api.Group("/signin")
	{
		signin.GET("/", adminHandler.GetSignin)
		signin.POST("/", adminHandler.PostSignin)
	}

	api.Use(middleware.AuthenticateAdmin)
	{
		api.GET("/logout", adminHandler.Logout)
		application := api.Group("/application")
		{
			application.GET("/", adminHandler.ViewApplications)
			application.POST("/approve", adminHandler.ApproveApplication)
			application.PATCH("/correction", adminHandler.FormCorrection)
		}

		duty := api.Group("/schedule")
		{
			duty.GET("/", adminHandler.GetScheduleDuty)
			duty.POST("/", adminHandler.ScheduleDuty)

		}

		leave := api.Group("/leave")
		{
			leave.GET("/request", adminHandler.GetAllLeaveRequets)
			leave.PATCH("/approve", adminHandler.ApproveLeaveRequests)
			leave.PATCH("/decline", adminHandler.DeclineLeaveRequests)

		}

		salary := api.Group("salary")
		{
			salary.POST("/add-details", adminHandler.AddSalaryDetails)
			salary.PATCH("/edit-details", adminHandler.EditSalaryDetails)
			salary.GET("/transactions", adminHandler.GetAllTransactions)

		}

	}
}
