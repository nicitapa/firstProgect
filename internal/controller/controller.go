package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/nicitapa/firstProgect/internal/contracts"
	"github.com/nicitapa/firstProgect/internal/errs"
	"net/http"
)

type Controller struct {
	router  *gin.Engine
	service contracts.ServiceI
}

func NewController(service contracts.ServiceI) *Controller {
	return &Controller{
		service: service,
		router:  gin.Default(),
	}
}

func (ctrl *Controller) handleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, errs.ErrEmployeesNotfound) ||
		errors.Is(err, errs.ErrUserNotFound) ||
		errors.Is(err, errs.ErrNotfound):
		c.JSON(http.StatusNotFound, CommonError{Error: err.Error()})
	case errors.Is(err, errs.ErrInvalidEmployeesID) || errors.Is(err, errs.ErrInvalidRequestBody):
		c.JSON(http.StatusBadRequest, CommonError{Error: err.Error()})
	case errors.Is(err, errs.ErrIncorrectUsernameOrPassword) || errors.Is(err, errs.ErrInvalidToken):
		c.JSON(http.StatusUnauthorized, CommonError{Error: err.Error()})
	case errors.Is(err, errs.ErrInvalidFieldValue) ||
		errors.Is(err, errs.ErrInvalidEmployeesName) ||
		errors.Is(err, errs.ErrUsernameAlreadyExists):
		c.JSON(http.StatusUnprocessableEntity, CommonError{Error: err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, CommonError{Error: err.Error()})
	}
}
