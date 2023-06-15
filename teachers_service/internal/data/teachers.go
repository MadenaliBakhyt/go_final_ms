package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"teachers_service/internal/validator"
	"time"
)

type Teacher struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Subject string `json:"subject"`
	Salary  int64  `json:"salary"`
}

func ValidateTeacher(v *validator.Validator, teacher *Teacher) {
	v.Check(teacher.Name != "", "name", "must be provided")
	v.Check(teacher.Surname != "", "surname", "must be provided")
	v.Check(teacher.Subject != "", "subject", "must be provided")
	v.Check(teacher.Salary != 0, "salary", "must be provided")
}

type TeacherModel struct {
	DB *sql.DB
}

func (m TeacherModel) Insert(teacher *Teacher) error {
	query := `INSERT INTO teachers (name, surname, subject, salary)
				VALUES ($1, $2, $3, $4)
				RETURNING id`
	args := []any{teacher.Name, teacher.Surname, teacher.Subject, teacher.Salary}
	return m.DB.QueryRow(query, args...).Scan(&teacher.ID)
}

func (m TeacherModel) Get(id int64) (*Teacher, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `
		SELECT *
		FROM teachers
		WHERE id = $1`
	var teacher Teacher

	err := m.DB.QueryRow(query, id).Scan(
		&teacher.ID,
		&teacher.Name,
		&teacher.Surname,
		&teacher.Subject,
		&teacher.Salary,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &teacher, nil
}

func (m TeacherModel) Update(teacher *Teacher) error {
	query := `
			UPDATE teachers
			SET name = $1, surname = $2, subject = $3, salary = $4
			WHERE id = $5
			RETURNING id`
	args := []any{
		teacher.Name,
		teacher.Surname,
		teacher.Subject,
		teacher.Salary,
		teacher.ID,
	}
	return m.DB.QueryRow(query, args...).Scan(&teacher.ID)
}

func (m TeacherModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	query := `
				DELETE FROM teachers
				WHERE id = $1`
	result, err := m.DB.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}

func (m TeacherModel) SelectAll() ([]*Teacher, error) {
	query := fmt.Sprintf(`
								SELECT *
								FROM teachers ORDER BY id`)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	teachers := []*Teacher{}
	for rows.Next() {
		var teacher Teacher

		err := rows.Scan(
			&teacher.ID,
			&teacher.Name,
			&teacher.Surname,
			&teacher.Subject,
			&teacher.Salary,
		)
		if err != nil {
			return nil, err
		}
		teachers = append(teachers, &teacher)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return teachers, nil
}
