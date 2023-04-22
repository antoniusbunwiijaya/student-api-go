package repository

import (
	"antoniusbunwijaya/student-api-go/helper"
	"antoniusbunwijaya/student-api-go/model/domain"
	"context"
	"database/sql"
)

type MajorRepositoryImpl struct {
}

func NewMajorRepository() MajorRepository {
	return &MajorRepositoryImpl{}
}

func (m MajorRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, major domain.Major) domain.Major {
	SQL := "insert into majors(major_name) values(?)"
	result, err := tx.ExecContext(ctx, SQL, major.MajorName)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	major.Id = int(id)
	return major
}

func (m MajorRepositoryImpl) FindByMajorName(ctx context.Context, tx *sql.Tx, majorName string) (domain.Major, bool) {
	SQL := "select id, major_name from majors where major_name = ?"
	rows, err := tx.QueryContext(ctx, SQL, majorName)
	helper.PanicIfError(err)
	defer rows.Close()

	major := domain.Major{}
	if rows.Next() {
		err := rows.Scan(&major.Id, &major.MajorName)
		helper.PanicIfError(err)
		return major, true
	} else {
		return major, false
	}
}
