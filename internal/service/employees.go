package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/nicitapa/firstProgect/internal/errs"
	"github.com/nicitapa/firstProgect/internal/models"
	"time"
)

var (
	defaultTTL = time.Minute * 5
)

func (s *Service) GetAllEmployees() (employees []models.Employees, err error) {
	ctx := context.Background()
	employees, err = s.repository.GetAllEmployees(ctx)
	if err != nil {
		return nil, err
	}

	return employees, nil
}

func (s *Service) GetEmployeesByID(id int) (employees models.Employees, err error) {
	ctx := context.Background()

	err = s.cache.Get(ctx, fmt.Sprintf("employees_%d", id), &employees)
	if err == nil {

		return employees, nil
	}

	employees, err = s.repository.GetEmployeesByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Employees{}, errs.ErrEmployeesNotfound
		}
		return models.Employees{}, err
	}

	if err = s.cache.Set(ctx, fmt.Sprintf("employees_%d", employees.ID), employees, defaultTTL); err != nil {
		fmt.Printf("error during cache set: %v\n", err.Error())
	}

	return employees, nil
}

func (s *Service) CreateEmployees(employees models.Employees) (err error) {
	ctx := context.Background()
	if len(employees.Name) < 4 {
		return errs.ErrInvalidEmployeesName
	}

	err = s.repository.CreateEmployees(ctx, employees)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateEmployeesByID(employees models.Employees) (err error) {
	ctx := context.Background()
	_, err = s.repository.GetEmployeesByID(ctx, employees.ID)
	if err != nil {
		if errors.Is(err, errs.ErrNotfound) {
			return errs.ErrEmployeesNotfound
		}
		return err
	}
	err = s.repository.UpdateEmployeesByID(ctx, employees)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteEmployeesByID(id int) (err error) {
	ctx := context.Background()
	_, err = s.repository.GetEmployeesByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errs.ErrEmployeesNotfound
		}
		return err
	}

	err = s.repository.DeleteEmployeesByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
