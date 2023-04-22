package helper

import (
	"antoniusbunwijaya/student-api-go/model/domain"
	"antoniusbunwijaya/student-api-go/model/web"
)

func ToStudentResponse(student domain.Student) web.StudentResponse {
	return web.StudentResponse{
		Id:        student.Id,
		Name:      student.Name,
		Age:       student.Age,
		Gender:    student.Gender,
		CreatedAt: student.CreatedAt,
		Major:     student.Major,
	}
}

func ToStudentDetailResponse(student domain.Student, hobbies []domain.Hobby) web.StudentDetailResponse {
	return web.StudentDetailResponse{
		Id:        student.Id,
		Name:      student.Name,
		Age:       student.Age,
		Gender:    student.Gender,
		CreatedAt: student.CreatedAt,
		Major:     student.Major,
		Hobbies:   hobbies,
	}
}

func ToStudentResponses(students []domain.Student) []web.StudentResponse {
	var studentResponses []web.StudentResponse
	//var hobbiesResponses []domain.Hobby
	for _, student := range students {
		studentResponses = append(studentResponses, ToStudentResponse(student))
	}
	return studentResponses
}
