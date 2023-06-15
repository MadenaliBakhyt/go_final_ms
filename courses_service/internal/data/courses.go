package data

import (
	"context"
	"courses_service/internal/validator"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Course struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Credits int64  `json:"credits"`
	Price   int64  `json:"price"`
}

func ValidateCourse(v *validator.Validator, course *Course) {
	v.Check(course.Name != "", "name", "must be provided")
	v.Check(course.Credits != 0, "credits", "must be provided")
	v.Check(course.Price != 0, "price", "must be provided")
}

type CourseModel struct {
	DB *sql.DB
}

func (m CourseModel) Insert(course *Course) error {
	query := `INSERT INTO courses (name, credits, price)
				VALUES ($1, $2, $3)
				RETURNING id`
	args := []any{course.Name, course.Credits, course.Price}
	return m.DB.QueryRow(query, args...).Scan(&course.ID)
}

func (m CourseModel) Get(id int64) (*Course, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `
		SELECT *
		FROM courses
		WHERE id = $1`

	var course Course

	err := m.DB.QueryRow(query, id).Scan(
		&course.ID,
		&course.Name,
		&course.Credits,
		&course.Price,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &course, nil
}

func (m CourseModel) Update(course *Course) error {
	query := `
			UPDATE courses
			SET name = $1, credits = $2, price = $4
			WHERE id = $4
			RETURNING id`
	args := []any{
		course.Name,
		course.Credits,
		course.Price,
		course.ID,
	}
	return m.DB.QueryRow(query, args...).Scan(&course.ID)
}

func (m CourseModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	query := `
				DELETE FROM courses
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

func (m CourseModel) SelectAll() ([]*Course, error) {
	query := fmt.Sprintf(`
								SELECT *
								FROM courses ORDER BY id`)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	courses := []*Course{}
	for rows.Next() {
		var course Course

		err := rows.Scan(
			&course.ID,
			&course.Name,
			&course.Credits,
			&course.Price,
		)
		if err != nil {
			return nil, err
		}
		courses = append(courses, &course)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return courses, nil
}
