package repository

import (
	"backend/models"
	"context"
)

type EmployeeRepo interface {
	Update(ctx context.Context, employee *models.Employee) (*models.Employee, error)
	Delete(ctx context.Context, id int64) (bool, error)
	Insert(ctx context.Context, employee *models.Employee) (int64, error)
	GetById(ctx context.Context, id int64) (*models.Employee, error)
	Fetch(ctx context.Context) ([]*models.Employee, error)
}
