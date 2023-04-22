package app

import (
	"antoniusbunwijaya/student-api-go/helper"
	"database/sql"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/student_api_go")
	helper.PanicIfError(err)

	return db
}
