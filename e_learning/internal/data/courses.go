package data

import (
	"database/sql"
	"e-learning/internal/validator"
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
