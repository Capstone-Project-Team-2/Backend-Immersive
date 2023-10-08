package handler

import (
	"capstone-tickets/apps/middlewares"
	"capstone-tickets/features/events"
	"capstone-tickets/helpers"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-playground/form"
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

	var filename string
	file, header, errFile := c.Request().FormFile("banner_picture")
	if errFile != nil {
		if strings.Contains(errFile.Error(), "no such file") {
			fmt.Println(helpers.DefaultFile)
			filename = helpers.DefaultFile
		} else {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400+" errFile "+errFile.Error(), nil))
		}
	}

	var eventReq EventRequest
	dec := form.NewDecoder()
	values, errForm := c.FormParams()
	if errForm != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400+" errForm "+errForm.Error(), nil))
	}

	errDec := dec.Decode(&eventReq, values)
	if errDec != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400+" errDec "+errDec.Error(), nil))
	}

	if filename == "" {
		filename = strings.ReplaceAll(header.Filename, " ", "_")
	}
	var eventCore = EventRequestToCore(eventReq)
	eventCore.PartnerID = id
	eventCore.BannerPicture = filename
	fmt.Println("handler event core:", eventCore)
	err := handler.eventService.Add(eventCore, file)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.Error500, nil))
	}
	return c.JSON(http.StatusCreated, helpers.WebResponse(http.StatusCreated, "operation success", nil))
}

func (handler *EventHandler) GetAll(c echo.Context) error {
	var pageParam, itemParam, searchParam string
	pageParam = c.QueryParam("page")
	itemParam = c.QueryParam("item")
	searchParam = c.QueryParam("search")

	result, next, err := handler.eventService.GetAll(pageParam, itemParam, searchParam)
	if err != nil {
		if strings.Contains(err.Error(), "no row affected") {
			return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, helpers.Error404, nil))
		}
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.Error500, nil))
	}
	var eventResp = ListEventCoreToResponse(result)
	return c.JSON(http.StatusOK, helpers.FindAllWebResponse(http.StatusOK, "operation success", eventResp, next))
}

func (handler *EventHandler) Get(c echo.Context) error {
	event_id := c.Param("event_id")
	result, err := handler.eventService.Get(event_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.Error500, nil))
	}
	var eventResp = EventCoreToResponse(result)
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "operation success", eventResp))
}

func (handler *EventHandler) Update(c echo.Context) error {
	idParam := c.Param("event_id")
	parter_id, role := middlewares.ExtractToken(c)
	if role != "Partner" {
		return c.JSON(http.StatusUnauthorized, helpers.WebResponse(http.StatusUnauthorized, helpers.Error401+" operation can be done by a partner", nil))
	}

	var filename string
	file, header, errFile := c.Request().FormFile("banner_picture")
	if errFile != nil {
		if strings.Contains(errFile.Error(), "no such file") {
			fmt.Println(helpers.DefaultFile)
			filename = helpers.DefaultFile
		} else {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400+" errFile "+errFile.Error(), nil))
		}
	}

	var eventReq EventRequest
	dec := form.NewDecoder()
	values, errForm := c.FormParams()
	if errForm != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400+" errForm "+errForm.Error(), nil))
	}

	errDec := dec.Decode(&eventReq, values)
	if errDec != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400+" errDec "+errDec.Error(), nil))
	}

	if filename == "" {
		filename = strings.ReplaceAll(header.Filename, " ", "_")
	}

	var eventCore = EventRequestToCore(eventReq)
	eventCore.BannerPicture = filename
	err := handler.eventService.Update(idParam, parter_id, eventCore, file)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.Error500+" "+err.Error(), nil))
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "operation success", nil))
}

func (handler *EventHandler) Validation(c echo.Context) error {
	_, role := middlewares.ExtractToken(c)
	if role != "Admin" && role != "Superadmin" {
		return c.JSON(http.StatusUnauthorized, helpers.WebResponse(http.StatusUnauthorized, helpers.Error401, nil))
	}
	event_id := c.Param("event_id")
	var valid ValidationRequest
	errBind := c.Bind(&valid)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.Error400+""+errBind.Error(), nil))
	}
	err := handler.eventService.Validate(event_id, valid.ValidationStatus)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.Error500, nil))
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "operation success", nil))
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
