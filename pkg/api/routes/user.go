package routes

import (
	"shiftsync/pkg/api/handler"

	"github.com/gin-gonic/gin"
)

func UserRoutes(api *gin.RouterGroup, employeeHandler *handler.EmployeeHandler) {

	// signup
	signup := api.Group("/signup")
	{
		signup.GET("/", employeeHandler.GetSignUp)
		signup.POST("/", employeeHandler.PostSignup)
	}
}
