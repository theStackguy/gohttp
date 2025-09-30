package employee

import (
	"backend/models"
	"backend/repository"
	"context"
	"database/sql"
	"gorm.io/gorm"
)

type sqlEmployeeRepo struct {
	Conn *sql.DB
}

var (
	employeeid  int64
	firstname   sql.NullString
	lastname    sql.NullString
	designation sql.NullString
	country     sql.NullString
	salary      float64
)

func NewSqlEmployeeRepo(Conn *sql.DB) repository.EmployeeRepo {
	return &sqlEmployeeRepo{
		Conn: Conn,
	}
}

func (m *sqlEmployeeRepo) Fetch(ctx context.Context) ([]*models.Employee, error) {
	query := `SELECT CAST(id AS NVARCHAR) AS ID, first_name, last_name , designation, salary, country FROM anandhu2.employee_new;`
	return m.fetch(ctx, query)
}

func (m *sqlEmployeeRepo) fetch(ctx context.Context, query string, args ...any) ([]*models.Employee, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	payload := make([]*models.Employee, 0)

	for rows.Next() {

		if err := rows.Scan(
			&employeeid,
			&firstname,
			&lastname,
			&designation,
			&salary,
			&country,
		); err != nil {
			return nil, err
		}

		singleEmployee := &models.Employee{
			Model: gorm.Model{ID: uint(employeeid)},
			FirstName:   firstname.String,
			LastName:    lastname.String,
			Designation: designation.String,
			Salary:      salary,
			Country:     country.String,
		}
		payload = append(payload, singleEmployee)
	}

	return payload, nil
}

func (m *sqlEmployeeRepo) Delete(ctx context.Context, id int64) (bool, error) {
	query := `DELETE FROM anandhu2.employee_new WHERE id = @EmployeeId`
	_, err := m.Conn.ExecContext(ctx, query, sql.Named("EmployeeId", id))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (m *sqlEmployeeRepo) Insert(ctx context.Context, employee *models.Employee) (int64, error) {
	query := `INSERT INTO anandhu2.employee_new ( first_name, last_name,  designation, salary, country) 
	          OUTPUT INSERTED.id
              VALUES ( @FirstName, @LastName, @Designation, CAST(@Salary AS INT), @Country)`

	err := m.Conn.QueryRowContext(
		ctx,
		query,
		sql.Named("FirstName", employee.FirstName),
		sql.Named("LastName", employee.LastName),
		sql.Named("Designation", employee.Designation),
		sql.Named("Salary", employee.Salary),
		sql.Named("Country", employee.Country)).Scan(&employeeid)
	if err != nil {
		return 0, err
	}
	return employeeid, nil

}

func (m *sqlEmployeeRepo) GetById(ctx context.Context, id int64) (*models.Employee, error) {
	query := `SELECT CAST(id AS NVARCHAR) AS ID, first_name, last_name , designation, salary, country FROM anandhu2.employee_new WHERE id = @EmployeeId;`

	err := m.Conn.QueryRowContext(ctx, query, sql.Named("EmployeeId", id)).Scan(&employeeid, &firstname, &lastname, &designation, &salary, &country)
	if err != nil {
		return nil, err
	}
	return &models.Employee{
		Model: gorm.Model{ID: uint(employeeid)},
		FirstName:   firstname.String,
		LastName:    lastname.String,
		Designation: designation.String,
		Salary:      salary,
		Country:     country.String,
	}, nil
}

func (m *sqlEmployeeRepo) Update(ctx context.Context, employee *models.Employee) (*models.Employee, error) {
	query := `UPDATE anandhu2.employee_new SET first_name=@FirstName, last_name=@LastName, designation=@Designation, country=@Country, salary= @Salary WHERE id= @EmpID;`
	_,err := m.Conn.ExecContext(
		ctx,
		query,
		sql.Named("FirstName", employee.FirstName),
		sql.Named("LastName", employee.LastName),
		sql.Named("Designation", employee.Designation),
		sql.Named("Country", employee.Country),
		sql.Named("salary", employee.Salary),
		sql.Named("EmpID", employee.ID))
	if err != nil {
		return nil, err
	}

	return employee, nil

}
