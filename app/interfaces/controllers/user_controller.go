package controllers

import (
	"viet_college_api/domain"
	"viet_college_api/interfaces/database"
	"viet_college_api/usecase"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create() {
	u := domain.User{}
	controller.Interactor.Add(u)

	return
}

func (controller *UserController) GetUser() domain.Users {
	res := controller.Interactor.GetInfo()
	return res
}

func (controller *UserController) Delete(id string) {
	controller.Interactor.Delete(id)
}
