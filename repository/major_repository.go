package repository

import (
	"antoniusbunwijaya/student-api-go/model/domain"
	"context"
	"database/sql"
)

type MajorRepository interface {
	Save(ctx context.Context, tx *sql.Tx, major domain.Major) domain.Major
	FindByMajorName(ctx context.Context, tx *sql.Tx, majorName string) (domain.Major, bool)
}
