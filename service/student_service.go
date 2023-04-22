package service

import (
	"antoniusbunwijaya/student-api-go/model/web"
	"context"
)

type StudentService interface {
	Create(ctx context.Context, request web.StudentCreateRequest) web.StudentResponse
	Update(ctx context.Context, request web.StudentUpdateRequest) web.StudentResponse
	Delete(ctx context.Context, studentId int)
	FindById(ctx context.Context, studentId int) web.StudentDetailResponse
	FindAll(ctx context.Context) []web.StudentResponse
}
