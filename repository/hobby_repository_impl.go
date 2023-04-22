package repository

import (
	"antoniusbunwijaya/student-api-go/helper"
	"antoniusbunwijaya/student-api-go/model/domain"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type HobbyRepositoryImpl struct {
}

func NewHobbyRepository() HobbyRepository {
	return &HobbyRepositoryImpl{}
}

func (h HobbyRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, hobby domain.Hobby) domain.Hobby {
	SQL := "insert into hobbies(hobby_name) values(?)"
	result, err := tx.ExecContext(ctx, SQL, hobby.HobbyName)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	hobby.Id = int(id)
	return hobby
}

func (h HobbyRepositoryImpl) FindByHobbyName(ctx context.Context, tx *sql.Tx, hobbyName string) (domain.Hobby, error) {
	SQL := "select id, hobby_name from hobbies where hobby_name = ?"
	rows, err := tx.QueryContext(ctx, SQL, hobbyName)
	helper.PanicIfError(err)
	defer rows.Close()

	hobby := domain.Hobby{}
	if rows.Next() {
		err := rows.Scan(&hobby.Id, &hobby.HobbyName)
		helper.PanicIfError(err)
		return hobby, nil
	} else {
		return hobby, errors.New("hobby is not found")
	}
}

func (h HobbyRepositoryImpl) GetHobbiesByStudentId(ctx context.Context, tx *sql.Tx, studentId int) []domain.Hobby {
	SQL := "select hobbies.id,hobbies.hobby_name from student_hobbies join hobbies on hobbies.id = student_hobbies.hobby_id join students on students.id = student_hobbies.student_id where student_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, studentId)
	helper.PanicIfError(err)
	defer rows.Close()
	var hobbies []domain.Hobby
	for rows.Next() {
		hobby := domain.Hobby{}
		err := rows.Scan(&hobby.Id, &hobby.HobbyName)
		helper.PanicIfError(err)
		hobbies = append(hobbies, hobby)
	}
	return hobbies
}

func (h HobbyRepositoryImpl) CreateStudentHobby(ctx context.Context, tx *sql.Tx, studentId int, hobbyId int) int {
	SQL := "insert into student_hobbies(student_id, hobby_id) values(?,?)"
	fmt.Println("student repo", studentId, "hobby repo", hobbyId)
	result, err := tx.ExecContext(ctx, SQL, studentId, hobbyId)
	helper.PanicIfError(err)
	//
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	fmt.Println("success", id)
	//return int(id)
	return int(id)
}

func (h HobbyRepositoryImpl) DeleteHobbiesByStudentId(ctx context.Context, tx *sql.Tx, studentId int) {
	SQL := "delete from student_hobbies where student_id = ?"
	_, err := tx.ExecContext(ctx, SQL, studentId)
	helper.PanicIfError(err)
}
