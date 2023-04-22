package app

import (
	"antoniusbunwijaya/student-api-go/controller"
	"antoniusbunwijaya/student-api-go/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(studentController controller.StudentController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/students", studentController.FindAll)
	router.GET("/api/students/:studentId", studentController.FindById)
	router.POST("/api/students", studentController.Create)
	router.PUT("/api/students/:studentId", studentController.Update)
	router.DELETE("/api/students/:studentId", studentController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
