package controllers

import (
	"fmt"
	"net/http"
	"skyshi-rest-api/models"
	"skyshi-rest-api/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ActivityController interface {
	GetAllActivity(c echo.Context) (err error)
	GetActivity(c echo.Context) (err error)
	CreateActivity(c echo.Context) (err error)
	UpdateActivity(c echo.Context) (err error)
	DeleteActivity(c echo.Context) (err error)
}

type activityController struct {
	activityService services.ActivityService
}

func NewActivityController(_s services.ActivityService) ActivityController {
	return activityController{
		activityService: _s,
	}
}

func (_c activityController) GetAllActivity(c echo.Context) (err error) {
	var req models.Activity

	result, err := _c.activityService.GetAllActivity(req)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  "Success",
		"message": "Success",
		"data":    result,
	})
}

func (_c activityController) CreateActivity(c echo.Context) (err error) {
	var req models.Activity
	if err = c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "bad request",
		})
	}
	if req.Title == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  "Bad Request",
			"message": "title cannot be null",
			"data":    nil,
		})
	}

	result, err := _c.activityService.CreateActivity(req)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"status":  "Success",
		"message": "Success",
		"data":    result,
	})
}

func (_c activityController) GetActivity(c echo.Context) (err error) {
	var req models.Activity

	req.ID, _ = strconv.Atoi(c.Param("id"))
	result, err := _c.activityService.GetActivity(req)

	if err != nil {
		msg := fmt.Sprintf("Activity with ID %d Not Found", req.ID)
		return c.JSON(http.StatusNotFound, echo.Map{
			"status":  "Not Found",
			"message": msg,
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  "Success",
		"message": "Success",
		"data":    result,
	})
}

func (_c activityController) UpdateActivity(c echo.Context) (err error) {
	var req models.Activity
	if err = c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "bad request",
		})
	}

	if req.Title == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  "Bad Request",
			"message": "title cannot be null",
			"data":    nil,
		})
	}

	req.ID, _ = strconv.Atoi(c.Param("id"))
	result, err := _c.activityService.UpdateActivity(req)

	if err != nil {
		msg := fmt.Sprintf("Activity with ID %d Not Found", req.ID)
		return c.JSON(http.StatusNotFound, echo.Map{
			"status":  "Not Found",
			"message": msg,
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  "Success",
		"message": "Success",
		"data":    result,
	})
}

func (_c activityController) DeleteActivity(c echo.Context) (err error) {
	var req models.Activity

	req.ID, _ = strconv.Atoi(c.Param("id"))
	result, err := _c.activityService.DeleteActivity(req)

	if err != nil {
		msg := fmt.Sprintf("Activity with ID %d Not Found", req.ID)
		return c.JSON(http.StatusNotFound, echo.Map{
			"status":  "Not Found",
			"message": msg,
			"data":    result,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  "Success",
		"message": "Success",
		"data":    result,
	})
}
