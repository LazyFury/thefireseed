package controller

import (
	"thefireseed/model"

	"github.com/lazyfury/go-web-template/controller"
)

func NewAppController() *controller.Controller {
	return &controller.Controller{
		DB:    model.DB,
		Model: &model.AppModel{},
		Auth:  Auth,
	}
}
