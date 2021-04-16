package controller

import (
	"shareInviteCode/model"

	"github.com/lazyfury/go-web-template/controller"
)

func NewAppController() *controller.Controller {
	return &controller.Controller{
		DB:    model.DB,
		Model: &model.AppModel{},
		Auth:  Auth,
	}
}

func NewActivityController() *controller.Controller {
	return &controller.Controller{
		DB:    model.DB,
		Model: &model.ActivityModel{},
		Auth:  Auth,
	}
}

func NewCodeController() *controller.Controller {
	return &controller.Controller{
		DB:    model.DB,
		Model: &model.CodeModel{},
		Auth:  Auth,
	}
}
