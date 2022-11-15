package controllers

import (
	"fmt"
	"net/http"
	"skyshi-rest-api/models"
	"skyshi-rest-api/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TodoController interface {
	GetAllTodo(c echo.Context) (err error)
	GetTodo(c echo.Context) (err error)
	CreateTodo(c echo.Context) (err error)
	UpdateTodo(c echo.Context) (err error)
	DeleteTodo(c echo.Context) (err error)
}

type todoController struct {
	todoService services.TodoService
}

func NewTodoController(_s services.TodoService) TodoController {
	return todoController{
		todoService: _s,
	}
}

func (_c todoController) GetAllTodo(c echo.Context) (err error) {
	var req models.Todo

	req.ActivityGroupID, _ = strconv.Atoi(c.QueryParam("activity_group_id"))
	result, err := _c.todoService.GetAllTodo(req)

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

func (_c todoController) CreateTodo(c echo.Context) (err error) {
	var req models.Todo

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

	if req.ActivityGroupID == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  "Bad Request",
			"message": "activity_group_id cannot be null",
			"data":    nil,
		})
	}

	if req.Priority == "" {
		req.Priority = "very-high"
	}

	req.IsActive = true

	result, err := _c.todoService.CreateTodo(req)

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

func (_c todoController) GetTodo(c echo.Context) (err error) {
	var req models.Todo

	req.ID, _ = strconv.Atoi(c.Param("id"))
	result, err := _c.todoService.GetTodo(req)

	if err != nil {
		msg := fmt.Sprintf("Todo with ID %d Not Found", req.ID)
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

func (_c todoController) UpdateTodo(c echo.Context) (err error) {
	var req models.Todo

	if err = c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "bad request",
		})
	}

	req.ID, _ = strconv.Atoi(c.Param("id"))
	result, err := _c.todoService.UpdateTodo(req)

	if err != nil {
		msg := fmt.Sprintf("Todo with ID %d Not Found", req.ID)
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

func (_c todoController) DeleteTodo(c echo.Context) (err error) {
	var req models.Todo

	req.ID, _ = strconv.Atoi(c.Param("id"))
	result, err := _c.todoService.DeleteTodo(req)

	if err != nil {
		msg := fmt.Sprintf("Todo with ID %d Not Found", req.ID)
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
