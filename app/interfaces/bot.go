package interfaces

import (
	"kanko-hackaton-22/app/interfaces/handler"

	"github.com/labstack/echo/v4"
)

func botRouter(e *echo.Echo, handler *handler.BotHandler) {
	e.POST("/", handler.Bot)
}
