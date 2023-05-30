package http

import (
	"shiftsync/pkg/api/handler"
	"shiftsync/pkg/api/routes"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewHTTPServer(employeeHandler *handler.EmployeeHandler, adminHandler *handler.AdminHandler) *ServerHTTP {

	// creating an instance of gin engine
	server := gin.New()

	// logger middleware
	server.Use(gin.Logger())

	//server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	url := ginSwagger.URL("localhost:8081/cmd/api/docs/swagger.json") // The url pointing to API definition
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url))

	routes.EmployeeRoutes(server.Group("/employee"), employeeHandler)
	routes.AdminRoutes(server.Group("/admin"), adminHandler)

	return &ServerHTTP{engine: server}
}

func (s *ServerHTTP) Start() {

	s.engine.Run(":8081")
}
