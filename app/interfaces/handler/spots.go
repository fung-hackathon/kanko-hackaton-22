package handler

import (
	"kanko-hackaton-22/app/data"

	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Spots(c echo.Context) error {
	data := struct {
		Spots []data.Spot
	}{
		Spots: data.SpotsData,
	}
	err := c.Render(http.StatusOK, "spots.html", data)
	log.Println(err)
	return err
}
