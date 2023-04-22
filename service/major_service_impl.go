package service

import (
	"antoniusbunwijaya/student-api-go/helper"
	"antoniusbunwijaya/student-api-go/model/domain"
	"antoniusbunwijaya/student-api-go/repository"
	"context"
	"database/sql"
)

type MajorServiceImpl struct {
	MajorRepository repository.MajorRepository
	DB              *sql.DB
}

func NewMajorService(majorRepository repository.MajorRepository, DB *sql.DB) MajorService {
	return &MajorServiceImpl{MajorRepository: majorRepository, DB: DB}
}

func (m MajorServiceImpl) Create(ctx context.Context, majorName string) domain.Major {
	tx, err := m.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	major := domain.Major{
		MajorName: majorName,
	}

	major = m.MajorRepository.Save(ctx, tx, major)
	// success
	return major
}

func (m MajorServiceImpl) FindByMajorName(ctx context.Context, majorName string) (domain.Major, bool) {
	tx, err := m.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	major, isExist := m.MajorRepository.FindByMajorName(ctx, tx, majorName)

	return major, isExist
}
