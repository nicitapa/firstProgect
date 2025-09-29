package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/nicitapa/firstProgect/internal/errs"
	"github.com/nicitapa/firstProgect/internal/models"
	"github.com/rs/zerolog"
	"net/http"
	"os"
	"strconv"
)

// GetAllEmployees
// @Summary Получение пользователей
// @Description Получение списка всех пользователей
// @Tags Employees
// @Produce json
// @Success 200 {array} models.Employees
// @Failure 500 {object} CommonError
// @Router /Employees [get]
func (ctrl *Controller) GetAllEmployees(c *gin.Context) {
	logger := zerolog.New(os.Stdout).With().Str("func_name", "controller.GetAllEmployees").Logger()
	userID := c.GetInt(userIDCtx)
	if userID == 0 {
		c.JSON(http.StatusBadRequest, CommonError{Error: "invalid userID in context"})
		return
	}

	logger.Debug().Int("user_id", userID).Msg("GetUser")

	products, err := ctrl.service.GetAllEmployees()
	if err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, products)
}

// GetEmployeesByID
// @Summary Получить пользователя по ID
// @Description Получение пользователя по ID
// @Tags Employees
// @Produce json
// @Param id path int true "id Пользлвател"
// @Success 200 {object} models.Employees
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /Employees/{id} [get]
func (ctrl *Controller) GetEmployeesByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		ctrl.handleError(c, errs.ErrInvalidEmployeesID)
		return
	}

	employees, err := ctrl.service.GetEmployeesByID(id)
	if err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, employees)
}

type CreateEmployeesRequest struct {
	Name  string `json:"name" `
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// CreateEmployees
// @Summary Добавление нового пользователя
// @Description Добавление нового пользователя
// @Tags Employees
// @Consume json
// @Produce json
// @Param request_body body CreateEmployeesRequest true "информация о новом пользователе"
// @Success 201 {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 422 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /Employees [post]
func (ctrl *Controller) CreateEmployees(c *gin.Context) {
	var employees models.Employees
	if err := c.ShouldBindJSON(&employees); err != nil {
		ctrl.handleError(c, errors.Join(errs.ErrInvalidRequestBody, err))
		return
	}
	if employees.Name == "" || employees.Email == "" || employees.Age < 0 {
		ctrl.handleError(c, errs.ErrInvalidFieldValue)
		return
	}

	if err := ctrl.service.CreateEmployees(employees); err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, CommonResponse{Message: "Employees created successfully"})
}

// UpdateEmployeesByID
// @Summary Обновить пользователя по ID
// @Description Обновление пользователя по ID
// @Tags Employees
// @Consume json
// @Produce json
// @Param id path int true "id пользователя"
// @Param request_body body CreateEmployeesRequest true "информация о пользователе"
// @Success 200 {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 422 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /Employees/{id} [put]
func (ctrl *Controller) UpdateEmployeesByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		ctrl.handleError(c, errs.ErrInvalidEmployeesID)
		return
	}

	var employees models.Employees
	if err = c.ShouldBindJSON(&employees); err != nil {
		ctrl.handleError(c, err)
		return
	}

	employees.ID = id

	if err = ctrl.service.UpdateEmployeesByID(employees); err != nil {
		ctrl.handleError(c, errors.Join(errs.ErrInvalidRequestBody, err))
		return
	}

	c.JSON(http.StatusOK, CommonResponse{Message: "Employees updated successfully"})
}

// DeleteEmployeesByID
// @Summary Удалить пользователя по ID
// @Description Удаление пользователя по ID
// @Tags Employees
// @Produce json
// @Param id path int true "id пользователя"
// @Success 200 {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /Employees/{id} [delete]
func (ctrl *Controller) DeleteEmployeesByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		ctrl.handleError(c, errs.ErrInvalidEmployeesID)
		return
	}

	if err = ctrl.service.DeleteEmployeesByID(id); err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, CommonResponse{Message: "Employees deleted successfully"})
}
