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
		application := api.Group("/application")
		{
			application.GET("/", adminHandler.ViewApplications)
			application.POST("/approve", adminHandler.ApproveApplication)
			application.PATCH("/correction", adminHandler.FormCorrection)
		}
	}
}
