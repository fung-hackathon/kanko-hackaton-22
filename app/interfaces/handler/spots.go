package handler

import (
	"kanko-hackaton-22/app/config"
	"kanko-hackaton-22/app/data"
	"kanko-hackaton-22/app/infra"
	"kanko-hackaton-22/app/logger"

	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ViewHandler struct {
	infra *infra.Firestore
}

func NewViewHandler(infra *infra.Firestore) *ViewHandler {
	return &ViewHandler{infra: infra}
}

func (h *ViewHandler) Spots(c echo.Context) error {
	data := struct {
		Spots []data.Spot
		Host  string
	}{
		Spots: data.SpotsData,
		Host:  config.HOST.File(""),
	}
	err := c.Render(http.StatusOK, "spots.html", data)
	log.Println(err)
	return err
}

type galleryProps struct {
	Spots []gallery
	Zero  bool
}

type gallery struct {
	Name     string
	Url      string
	OpenTime string
	Image    string
	Quiz     quiz
	Progress bool
}

type quiz struct {
	Q       string
	A       string
	Comment string
}

func (h *ViewHandler) Gallery(c echo.Context) error {
	userID := c.QueryParam("userid")

	user, err := h.infra.Get(userID)
	if err != nil {
		logger.Log{
			Message: "cannot get ",
			Cause:   err,
		}.Err()
		c.String(http.StatusInternalServerError, err.Error())
	}

	progress_raw, ok := (user["progress"]).([]interface{})
	if !ok {
		logger.Log{
			Message: "cannot cast to []interface{} from interface{}",
			Cause:   err,
		}.Err()
		c.String(http.StatusInternalServerError, err.Error())
	}
	var progress []gallery
	zeroFlag := true
	for i, v := range progress_raw {
		newProg, ok := v.(bool)
		if !ok {
			logger.Log{
				Message: "cannot cast to bool from interface{}",
				Cause:   err,
			}.Err()
			c.String(http.StatusInternalServerError, err.Error())
		}

		newValue := gallery{
			Name:     data.SpotsData[i].Name,
			Url:      data.SpotsData[i].Url,
			OpenTime: data.SpotsData[i].OpenTime,
			Image:    data.SpotsData[i].Image,
			Quiz: quiz{
				Q:       data.SpotsData[i].Quiz.Q,
				A:       data.SpotsData[i].Quiz.A,
				Comment: data.SpotsData[i].Quiz.Comment,
			},
			Progress: newProg,
		}

		progress = append(progress, newValue)

		if newProg {
			zeroFlag = false
		}
	}

	log.Println(progress)

	data := struct {
		Gallery galleryProps
		Host    string
	}{
		Gallery: galleryProps{
			Spots: progress,
			Zero:  zeroFlag,
		},
		Host: config.HOST.File(""),
	}

	if err := c.Render(http.StatusOK, "gallery.html", data); err != nil {
		logger.Log{
			Message: "cannot compile template",
			Cause:   err,
		}.Err()
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func (h *ViewHandler) Landing(c echo.Context) error {
	data := struct {
		Spots []data.Spot
		Host  string
	}{
		Spots: data.SpotsData,
		Host:  config.HOST.File(""),
	}
	err := c.Render(http.StatusOK, "index.html", data)
	log.Println(err)
	return err
}
