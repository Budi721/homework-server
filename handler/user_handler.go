package handler

import (
	"github.com/go-chi/chi/v5"
	"homework-server/helper"
	"homework-server/service"
	"net/http"
)

type UserHandler interface {
	GetDetail(w http.ResponseWriter, r *http.Request)
	GetFollowers(w http.ResponseWriter, r *http.Request)
}

type userHandlerImpl struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandlerImpl{userService: userService}
}

func (u userHandlerImpl) GetDetail(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	response, err := u.userService.GetDetail(param)
	if err != nil {
		helper.NewHttpResponse(r, w, http.StatusNotFound, nil, err)
		return
	}
	helper.NewHttpResponse(r, w, http.StatusOK, response, err)
}

func (u userHandlerImpl) GetFollowers(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "username")
	response, err := u.userService.GetFollower(param)
	if err != nil {
		helper.NewHttpResponse(r, w, http.StatusNotFound,nil, err)
		return
	}
	helper.NewHttpResponse(r, w, http.StatusOK, response, err)
}
