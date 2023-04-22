package service

import (
	"antoniusbunwijaya/student-api-go/exception"
	"antoniusbunwijaya/student-api-go/helper"
	"antoniusbunwijaya/student-api-go/model/domain"
	"antoniusbunwijaya/student-api-go/model/web"
	"antoniusbunwijaya/student-api-go/repository"
	"context"
	"database/sql"
	"fmt"
	"github.com/go-playground/validator/v10"
	"time"
)

type StudentServiceImpl struct {
	StudentRepository repository.StudentRepository
	DB                *sql.DB
	Validate          *validator.Validate
	HobbyService      HobbyService
	MajorService      MajorService
}

func NewStudentService(
	studentRepository repository.StudentRepository,
	DB *sql.DB,
	validate *validator.Validate,
	hobbyService HobbyService,
	majorService MajorService) StudentService {
	return &StudentServiceImpl{
		StudentRepository: studentRepository,
		DB:                DB,
		Validate:          validate,
		HobbyService:      hobbyService,
		MajorService:      majorService,
	}
}

func (s StudentServiceImpl) Create(ctx context.Context, request web.StudentCreateRequest) web.StudentResponse {
	err := s.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	major, isExistMajor := s.MajorService.FindByMajorName(ctx, request.Major)
	//fmt.Println(isExistMajor)
	if isExistMajor == false {
		major = s.MajorService.Create(ctx, request.Major)
	}

	student := domain.Student{
		Name:      request.Name,
		Age:       request.Age,
		Gender:    request.Gender,
		CreatedAt: time.DateTime,
		Major:     major,
	}

	student = s.StudentRepository.Save(ctx, tx, student)
	studentId := student.Id

	hobbies := s.HobbyService.Creates(ctx, request.Hobbies)
	fmt.Println(hobbies, studentId)
	for i := 0; i < len(hobbies); i++ {
		// todo bug stuck in sending request...
		//s.HobbyService.CreateStudentHobby(ctx, studentId, hobbies[i].Id)
	}
	return helper.ToStudentResponse(student)
}

func (s StudentServiceImpl) Update(ctx context.Context, request web.StudentUpdateRequest) web.StudentResponse {
	err := s.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	student, err := s.StudentRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	student.Name = request.Name
	student.Age = request.Age
	student.Gender = request.Gender
	//student.CreatedAt = request.CreatedAt
	//student.Major.MajorName = request.Major
	student = s.StudentRepository.Update(ctx, tx, student)

	// todo update student_hobbies and major...
	return helper.ToStudentResponse(student)
}

func (s StudentServiceImpl) Delete(ctx context.Context, studentId int) {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	student, err := s.StudentRepository.FindById(ctx, tx, studentId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	s.HobbyService.DeleteHobbiesByStudentId(ctx, studentId)
	s.StudentRepository.Delete(ctx, tx, student)
}

func (s StudentServiceImpl) FindById(ctx context.Context, studentId int) web.StudentDetailResponse {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	fmt.Println("student_service_impl1 findById")
	student, err := s.StudentRepository.FindById(ctx, tx, studentId)
	fmt.Println("student_service_impl2 findById")
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	hobbies := s.HobbyService.GetHobbiesByStudentId(ctx, studentId)
	fmt.Println("hobbies", hobbies)
	return helper.ToStudentDetailResponse(student, hobbies)
}

func (s StudentServiceImpl) FindAll(ctx context.Context) []web.StudentResponse {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	students := s.StudentRepository.FindAll(ctx, tx)

	return helper.ToStudentResponses(students)
}
