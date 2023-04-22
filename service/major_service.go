package service

import (
	"antoniusbunwijaya/student-api-go/model/domain"
	"context"
)

type MajorService interface {
	Create(ctx context.Context, majorName string) domain.Major
	FindByMajorName(ctx context.Context, majorName string) (domain.Major, bool)
}
