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
		form := api.Group("/form")
		{
			form.GET("/", employeeHandler.GetForm)
			form.POST("/", employeeHandler.PostForm)
		}
	}
}
