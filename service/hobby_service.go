package service

import (
	"antoniusbunwijaya/student-api-go/model/domain"
	"context"
)

type HobbyService interface {
	Creates(ctx context.Context, hobbyName []string) []domain.Hobby
	FindByHobbyName(ctx context.Context, hobbyName string) (domain.Hobby, bool)
	GetHobbiesByStudentId(ctx context.Context, studentId int) []domain.Hobby
	CreateStudentHobby(ctx context.Context, studentId int, hobbyId int)
	DeleteHobbiesByStudentId(ctx context.Context, studentId int)
}
