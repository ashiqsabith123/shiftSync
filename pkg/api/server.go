package http

import (
	"shiftsync/pkg/api/handler"
	"shiftsync/pkg/api/routes"

	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewHTTPServer(employeeHandler *handler.EmployeeHandler) *ServerHTTP {

	// creating an instance of gin engine
	server := gin.New()

	// logger middleware
	server.Use(gin.Logger())

	routes.UserRoutes(server.Group("/employee"), employeeHandler)

	return &ServerHTTP{engine: server}
}

func (s *ServerHTTP) Start() {

	s.engine.Run(":8000")
}
