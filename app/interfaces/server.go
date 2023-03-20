package interfaces

import (
	"fmt"
	"kanko-hackaton-22/app/config"
	"kanko-hackaton-22/app/logger"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Router *echo.Echo
}

func NewServer() *Server {
	return &Server{echo.New()}
}

func (s *Server) Serve() {
	s.Router.Use(logger.EchoLogger())

	s.Router.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	s.Router.Start(fmt.Sprintf(":%s", config.PORT))
}
