package routes

import (
	"shiftsync/pkg/api/handler"

	"github.com/gin-gonic/gin"
)

func UserRoutes(api *gin.RouterGroup, userHandler *handler.EmployeeHandler) {

	// signup
	signup := api.Group("/signup")
	{
		signup.GET("/", userHandler.GetSignUp)
	}
}
