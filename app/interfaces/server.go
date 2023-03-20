package interfaces

import (
	"fmt"
	"kanko-hackaton-22/app/command"
	"kanko-hackaton-22/app/config"
	"kanko-hackaton-22/app/infra"
	"kanko-hackaton-22/app/interfaces/handler"
	"kanko-hackaton-22/app/logger"
	"log"
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
	// DI
	infra, err := infra.Initialize()
	if err != nil {
		log.Fatalln(err)
	}
	command := command.New(infra)
	botHandler := handler.NewBotHandler(command)
	viewHandler := handler.NewViewHandler(infra)

	// Routing
	s.Router.Use(logger.EchoLogger())

	s.Router.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	botRouter(s.Router, botHandler)
	viewRouter(s.Router, viewHandler)

	s.Router.Start(fmt.Sprintf(":%s", config.PORT))
}
