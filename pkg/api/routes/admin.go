package routes

import (
	"shiftsync/pkg/api/handler"

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

	forms := api.Group("/forms")
	{
		forms.GET("/", adminHandler.ViewApplications)
	}
}
