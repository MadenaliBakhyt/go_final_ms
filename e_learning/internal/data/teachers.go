package data

import (
	"database/sql"
	"e-learning/internal/validator"
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
