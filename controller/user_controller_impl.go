package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"golang-fiber/helper"
	"golang-fiber/model/web"
	"golang-fiber/service"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserControllerImpl(userService service.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := web.UsersCreateRequest{}
	helper.BodyToRequest(request, &userCreateRequest)

	response := controller.UserService.Create(request.Context(), userCreateRequest)

	webResponse := web.Response{
		Status: "OK",
		Data:   response,
	}

	helper.WriteToBody(writer, webResponse)
}

func (controller *UserControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userUpdateRequest := web.UsersUpdateRequest{}
	helper.BodyToRequest(request, &userUpdateRequest)

	userUpdateRequest.Id = params.ByName("user_id") // menampung data dari inputan request

	response := controller.UserService.Update(request.Context(), userUpdateRequest)

	webResponse := web.Response{
		Status: "OK",
		Data:   response,
	}

	helper.WriteToBody(writer, webResponse)
}

func (controller *UserControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("user_id") // menampung data dari inputan request

	controller.UserService.Delete(request.Context(), userId)

	webResponse := web.Response{
		Status: "OK",
		Data:   "has been deleted",
	}

	helper.WriteToBody(writer, webResponse)
}

func (controller *UserControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("user_id") // menampung data dari inputan request

	response := controller.UserService.FindById(request.Context(), userId)

	webResponse := web.Response{
		Status: "OK",
		Data:   response,
	}

	helper.WriteToBody(writer, webResponse)
}

func (controller *UserControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	response := controller.UserService.FindAll(request.Context())

	webResponse := web.Response{
		Status: "OK",
		Data:   response,
	}

	helper.WriteToBody(writer, webResponse)
}

func (controller *UserControllerImpl) Auth(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	UserAuthRequest := web.UserAuthRequest{}
	helper.BodyToRequest(request, &UserAuthRequest)

	response := controller.UserService.Auth(request.Context(), UserAuthRequest)

	webResponse := web.Response{
		Status: "OK",
		Data:   response,
	}

	helper.WriteToBody(writer, webResponse)
}

func (controller *UserControllerImpl) CreateWithRefreshToken(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	token := request.Header.Get("Authorization")

	response := controller.UserService.CreateWithRefreshToken(request.Context(), token)

	webResponse := web.Response{
		Status: "OK",
		Data:   response,
	}

	helper.WriteToBody(writer, webResponse)
}
