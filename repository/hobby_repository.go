package repository

import (
	"antoniusbunwijaya/student-api-go/model/domain"
	"context"
	"database/sql"
)

type HobbyRepository interface {
	Save(ctx context.Context, tx *sql.Tx, hobby domain.Hobby) domain.Hobby
	FindByHobbyName(ctx context.Context, tx *sql.Tx, hobbyName string) (domain.Hobby, error)
	GetHobbiesByStudentId(ctx context.Context, tx *sql.Tx, studentId int) []domain.Hobby
	CreateStudentHobby(ctx context.Context, tx *sql.Tx, studentId int, hobbyId int) int
	DeleteHobbiesByStudentId(ctx context.Context, tx *sql.Tx, studentId int)
}
