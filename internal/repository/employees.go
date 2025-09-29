package repository

import (
	"context"
	"github.com/nicitapa/firstProgect/internal/models"
)

func (r *Repository) GetAllEmployees(ctx context.Context) (employees []models.Employees, err error) {

	if err = r.db.SelectContext(ctx, &employees, `
		SELECT id, name, email, age, 
		FROM employees
		ORDER BY id`); err != nil {
		return nil, r.translateError(err)
	}

	return employees, nil

}

func (r *Repository) GetEmployeesByID(ctx context.Context, id int) (employees models.Employees, err error) {
	if err = r.db.GetContext(ctx, &employees, `
		SELECT id, name, email, age
		FROM employees
		WHERE id = $1`, id); err != nil {
		return models.Employees{}, r.translateError(err)
	}

	return employees, nil
}

func (r *Repository) CreateEmployees(ctx context.Context, employees models.Employees) (err error) {
	_, err = r.db.ExecContext(ctx, `INSERT INTO employees (name, email, age)
					VALUES ($1, $2, $3)`,
		employees.Name,
		employees.Email,
		employees.Age)
	if err != nil {
		return r.translateError(err)
	}

	return nil
}

func (r *Repository) UpdateEmployeesByID(ctx context.Context, employees models.Employees) (err error) {
	_, err = r.db.ExecContext(ctx, `
		UPDATE employees SET name = $1, 
		                    email = $2, 
		                    age = $3,
		                    		                WHERE id = $4`,
		employees.Name,
		employees.Email,
		employees.Age,
		employees.ID)
	if err != nil {
		return r.translateError(err)
	}

	return nil
}

func (r *Repository) DeleteEmployeesByID(ctx context.Context, id int) (err error) {
	_, err = r.db.ExecContext(ctx, `DELETE FROM employees WHERE id = $1`, id)
	if err != nil {
		return r.translateError(err)
	}

	return nil
}
