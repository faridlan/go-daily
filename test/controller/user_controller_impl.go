package controller

import (
	"net/http"
	"strconv"

	"github.com/faridlan/daily/test/helper"
	"github.com/faridlan/daily/test/model/web"
	"github.com/faridlan/daily/test/service"
	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func (controller *UserControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	UserCreate := web.UserCreate{}
	helper.ReadFromReqBody(request, &UserCreate)

	user := controller.UserService.Create(request.Context(), UserCreate)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   user,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	UserUpdate := web.UserUpdate{}
	helper.ReadFromReqBody(request, &UserUpdate)

	userId := params.ByName("id")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	UserUpdate.Id = id

	user := controller.UserService.Update(request.Context(), UserUpdate)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   user,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("id")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	controller.UserService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("id")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	user := controller.UserService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   user,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	users := controller.UserService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   users,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
