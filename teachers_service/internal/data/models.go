package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Teachers TeacherModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Teachers: TeacherModel{DB: db},
	}
}
