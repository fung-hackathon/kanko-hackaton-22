package handler

import (
	"kanko-hackaton-22/app/data"

	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ViewHandler struct{}

func NewViewHandler() *ViewHandler {
	return &ViewHandler{}
}

func (h *ViewHandler) Spots(c echo.Context) error {
	data := struct {
		Spots []data.Spot
	}{
		Spots: data.SpotsData,
	}
	err := c.Render(http.StatusOK, "spots.html", data)
	log.Println(err)
	return err
}

func (h *ViewHandler) Gallery(c echo.Context) error {
	data := struct {
		Spots []data.Spot
	}{
		Spots: data.SpotsData,
	}
	err := c.Render(http.StatusOK, "gallery.html", data)
	log.Println(err)
	return err
}
