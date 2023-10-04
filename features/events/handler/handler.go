package handler

import (
	"capstone-tickets/apps/middlewares"
	"capstone-tickets/features/events"
	"capstone-tickets/helpers"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/monoculum/formam/v3"
)

type EventHandler struct {
	eventService events.EventServiceInterface
}

func New(service events.EventServiceInterface) *EventHandler {
	return &EventHandler{
		eventService: service,
	}
}

func (handler *EventHandler) Add(c echo.Context) error {
	id, role := middlewares.ExtractToken(c)
	if role != "Partner" {
		return c.JSON(http.StatusUnauthorized, helpers.WebResponse(http.StatusUnauthorized, helpers.Error401, nil))
	}
	var eventReq EventRequest

	// errBind := c.Bind(&eventReq)
	// if errBind != nil {
	// 	return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400+" "+errBind.Error(), nil))
	// }

	form, errForm := c.FormParams()
	if errForm != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400+" "+errForm.Error(), nil))
	}

	dec := formam.NewDecoder(&formam.DecoderOptions{TagName: "formam"})
	errDec := dec.Decode(form, &eventReq)
	if errDec != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400+" "+errDec.Error(), nil))
	}

	var filename string
	file, header, errFile := c.Request().FormFile("banner_picture")
	if errFile != nil {
		if strings.Contains(errFile.Error(), "no such file") {
			filename = helpers.DefaultFile
		}
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400+" "+errFile.Error(), nil))
	}
	if filename == "" {
		filename = strings.ReplaceAll(header.Filename, " ", "_")
	}
	var eventCore = EventRequestToCore(eventReq)
	eventCore.PartnerID = id
	eventCore.BannerPicture = filename

	err := handler.eventService.Add(eventCore, file)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.Error500, nil))
	}
	return c.JSON(http.StatusCreated, helpers.WebResponse(http.StatusCreated, "operation success", nil))
}

func (handler *EventHandler) GetAll(c echo.Context) error {
	userId, role := middlewares.ExtractToken(c)
	validation := c.Param("validation")
	execution := c.Param("execution")
	result, err := handler.eventService.GetAll(userId, role, validation, execution)
	if err != nil {
		if strings.Contains(err.Error(), "no row affected") {
			return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, helpers.Error404, nil))
		}
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.Error500, nil))
	}
	var eventResp = ListEventCoreToResponse(result)
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "operation success", eventResp))
}

func (handler *EventHandler) Test(c echo.Context) error {
	var eventReq EventRequest
	form, errForm := c.FormParams()
	if errForm != nil {
		return c.JSON(http.StatusBadRequest, errForm)
	}
	log.Println(form)
	dec := formam.NewDecoder(&formam.DecoderOptions{TagName: "formam"})
	err := dec.Decode(form, &eventReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	var eventCore = EventRequestToCore(eventReq)
	log.Println(eventCore)
	return c.JSON(http.StatusOK, eventCore)
}
