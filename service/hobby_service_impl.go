package service

import (
	"antoniusbunwijaya/student-api-go/helper"
	"antoniusbunwijaya/student-api-go/model/domain"
	"antoniusbunwijaya/student-api-go/repository"
	"context"
	"database/sql"
)

type HobbyServiceImpl struct {
	HobbyRepository repository.HobbyRepository
	DB              *sql.DB
}

func NewHobbyService(hobbyRepository repository.HobbyRepository, DB *sql.DB) HobbyService {
	return &HobbyServiceImpl{HobbyRepository: hobbyRepository, DB: DB}
}

func (h HobbyServiceImpl) Creates(ctx context.Context, tx *sql.Tx, hobbyName []string) []domain.Hobby {
	var hobbiesArrays []domain.Hobby
	for i := 0; i < len(hobbyName); i++ {
		hobbyAvailable, isExist := h.FindByHobbyName(ctx, hobbyName[i])
		//fmt.Println(hobbyName[i], isExist)
		if isExist == false {
			hobby := domain.Hobby{
				HobbyName: hobbyName[i],
			}
			hobby = h.HobbyRepository.Save(ctx, tx, hobby)
			hobbiesArrays = append(hobbiesArrays, hobby)
		} else {
			hobbiesArrays = append(hobbiesArrays, hobbyAvailable)
		}
	}

	// success
	return hobbiesArrays
}

func (h HobbyServiceImpl) FindByHobbyName(ctx context.Context, hobbyName string) (domain.Hobby, bool) {
	tx, err := h.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	hobby, err := h.HobbyRepository.FindByHobbyName(ctx, tx, hobbyName)

	if err != nil {
		return hobby, false
	}
	return hobby, true
}

func (h HobbyServiceImpl) GetHobbiesByStudentId(ctx context.Context, studentId int) []domain.Hobby {
	tx, err := h.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	hobbies := h.HobbyRepository.GetHobbiesByStudentId(ctx, tx, studentId)

	return hobbies
}

func (h HobbyServiceImpl) CreateStudentHobby(ctx context.Context, tx *sql.Tx, studentId int, hobbyId int) {
	h.HobbyRepository.CreateStudentHobby(ctx, tx, studentId, hobbyId)
}

func (h HobbyServiceImpl) DeleteHobbiesByStudentId(ctx context.Context, tx *sql.Tx, studentId int) {
	h.HobbyRepository.DeleteHobbiesByStudentId(ctx, tx, studentId)
}
