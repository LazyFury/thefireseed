package controller

import (
	"thefireseed/model"

	"github.com/lazyfury/go-web-template/controller"
)

type CodeController struct {
	controller.Controller
}

func NewCodeController() *CodeController {
	return &CodeController{
		controller.Controller{
			DB:    model.DB,
			Model: &model.CodeModel{},
			Auth:  Auth,
		},
	}
}
