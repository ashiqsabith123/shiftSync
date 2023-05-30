package routes

import (
	"shiftsync/pkg/api/handler"
	"shiftsync/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func EmployeeRoutes(api *gin.RouterGroup, employeeHandler *handler.EmployeeHandler) {

	// signup
	signup := api.Group("/signup")
	{
		signup.GET("/", employeeHandler.GetSignUp)
		signup.POST("/", employeeHandler.PostSignup)
		signup.POST("/verify-otp", employeeHandler.VerifyOtp)
	}

	signin := api.Group("/signin")
	{
		signin.GET("/", employeeHandler.GetLogin)
		signin.POST("/", employeeHandler.PostLogin)

	}

	api.Use(middleware.AuthenticateUser)
	{

		api.GET("/dashboard", employeeHandler.GetDashboard)
		api.GET("/attendance", employeeHandler.Attendance)
		api.GET("/logout", employeeHandler.Logout)

		form := api.Group("/form")
		{
			form.GET("/", employeeHandler.GetForm)
			form.POST("/", employeeHandler.PostForm)
		}

		duty := api.Group("/duty")
		{
			duty.GET("/", employeeHandler.GetDuty)
			duty.GET("/punchin", employeeHandler.PunchIn)
			duty.POST("/punchin", employeeHandler.VerifyOtpPunchin)
			duty.GET("/punchout", employeeHandler.PunchOut)
		}

		leave := api.Group("/leave")
		{
			leave.POST("/apply", employeeHandler.ApplyLeave)
			leave.GET("/status", employeeHandler.LeaveStatus)
		}

		salary := api.Group("/salary")
		{
			salary.GET("/history", employeeHandler.TransactionHistory)
			salary.GET("/details", employeeHandler.SalaryDetails)
			salary.GET("/download-slip", employeeHandler.SalarySlip)
		}

	}
}
