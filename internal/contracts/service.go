package contracts

import (
	"context"
	"github.com/nicitapa/firstProgect/internal/models"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type ServiceI interface {
	GetAllEmployees() (employees []models.Employees, err error)
	GetEmployeesByID(id int) (employees models.Employees, err error)
	CreateEmployees(employees models.Employees) (err error)
	UpdateEmployeesByID(employees models.Employees) (err error)
	DeleteEmployeesByID(id int) (err error)

	CreateUser(ctx context.Context, user models.User) (err error)
	Authenticate(ctx context.Context, user models.User) (int, models.Role, error)
}
