package service

import (
	"antoniusbunwijaya/student-api-go/model/domain"
	"context"
	"database/sql"
)

type HobbyService interface {
	Creates(ctx context.Context, tx *sql.Tx, hobbyName []string) []domain.Hobby
	FindByHobbyName(ctx context.Context, hobbyName string) (domain.Hobby, bool)
	GetHobbiesByStudentId(ctx context.Context, studentId int) []domain.Hobby
	CreateStudentHobby(ctx context.Context, tx *sql.Tx, studentId int, hobbyId int)
	DeleteHobbiesByStudentId(ctx context.Context, tx *sql.Tx, studentId int)
}
