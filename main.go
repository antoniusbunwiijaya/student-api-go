package main

import (
	"antoniusbunwijaya/student-api-go/app"
	"antoniusbunwijaya/student-api-go/controller"
	"antoniusbunwijaya/student-api-go/helper"
	"antoniusbunwijaya/student-api-go/middleware"
	"antoniusbunwijaya/student-api-go/repository"
	"antoniusbunwijaya/student-api-go/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {

	db := app.NewDB()
	validate := validator.New()

	hobbyRepository := repository.NewHobbyRepository()
	hobbyService := service.NewHobbyService(hobbyRepository, db)

	majorRepository := repository.NewMajorRepository()
	majorService := service.NewMajorService(majorRepository, db)

	studentRepository := repository.NewStudentRepository()
	studentService := service.NewStudentService(
		studentRepository,
		db,
		validate,
		hobbyService,
		majorService)
	studentController := controller.NewStudentController(studentService)

	router := app.NewRouter(studentController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
