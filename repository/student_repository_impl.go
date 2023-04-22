package repository

import (
	"antoniusbunwijaya/student-api-go/helper"
	"antoniusbunwijaya/student-api-go/model/domain"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type StudentRepositoryImpl struct {
}

func NewStudentRepository() StudentRepository {
	return &StudentRepositoryImpl{}
}

func (StudentRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, student domain.Student) domain.Student {
	SQL := "insert into students(name,age,gender,created_at,major_id) values(?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, student.Name, student.Age, student.Gender, student.CreatedAt, student.Major.Id)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	student.Id = int(id)
	return student
}

func (s StudentRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, student domain.Student) domain.Student {
	SQL := "update students set name = ?,age =?, gender=?  where id = ?"
	_, err := tx.ExecContext(ctx, SQL, student.Name, student.Age, student.Gender, student.Id)
	helper.PanicIfError(err)
	return student
}

func (s StudentRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, student domain.Student) {
	SQL := "delete from students where id = ?"
	_, err := tx.ExecContext(ctx, SQL, student.Id)
	helper.PanicIfError(err)
}

func (s StudentRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, studentId int) (domain.Student, error) {
	SQL := "select students.id, students.name, students.age, students.gender, students.created_at, students.major_id, majors.major_name from students join majors on majors.id = students.major_id where students.id = ?"
	rows, err := tx.QueryContext(ctx, SQL, studentId)
	helper.PanicIfError(err)
	defer rows.Close()

	student := domain.Student{}
	if rows.Next() {
		err := rows.Scan(&student.Id, &student.Name, &student.Age, &student.Gender, &student.CreatedAt, &student.Major.Id, &student.Major.MajorName)
		fmt.Println("student_repo_impl2 findById")
		helper.PanicIfError(err)
		fmt.Println("student_repo_impl3 findById")
		return student, nil
	} else {
		return student, errors.New("student is not found")
	}
}

func (s StudentRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Student {
	SQL := "select students.id, students.name, students.age, students.gender, students.created_at, students.major_id, majors.major_name from students join majors on majors.id = students.major_id"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	defer rows.Close()
	var students []domain.Student
	for rows.Next() {
		student := domain.Student{}
		err := rows.Scan(&student.Id, &student.Name, &student.Age, &student.Gender, &student.CreatedAt, &student.Major.Id, &student.Major.MajorName)
		helper.PanicIfError(err)
		students = append(students, student)
	}
	return students
}
