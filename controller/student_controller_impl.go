package controller

import (
	"antoniusbunwijaya/student-api-go/helper"
	"antoniusbunwijaya/student-api-go/model/web"
	"antoniusbunwijaya/student-api-go/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type StudentControllerImpl struct {
	StudentService service.StudentService
}

func NewStudentController(studentService service.StudentService) StudentController {
	return &StudentControllerImpl{
		StudentService: studentService,
	}
}

func (s StudentControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	studentCreateRequest := web.StudentCreateRequest{}
	helper.ReadFromRequestBody(request, &studentCreateRequest)

	studentResponse := s.StudentService.Create(request.Context(), studentCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   studentResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (s StudentControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	studentUpdateRequest := web.StudentUpdateRequest{}
	helper.ReadFromRequestBody(request, &studentUpdateRequest)

	studentId := params.ByName("studentId")
	id, err := strconv.Atoi(studentId)
	helper.PanicIfError(err)

	studentUpdateRequest.Id = id

	studentResponse := s.StudentService.Update(request.Context(), studentUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   studentResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (s StudentControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	studentId := params.ByName("studentId")
	id, err := strconv.Atoi(studentId)
	helper.PanicIfError(err)

	s.StudentService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (s StudentControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	studentId := params.ByName("studentId")
	id, err := strconv.Atoi(studentId)
	helper.PanicIfError(err)

	studentResponse := s.StudentService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   studentResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (s StudentControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	studentResponses := s.StudentService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   studentResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
