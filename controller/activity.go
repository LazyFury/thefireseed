package controller

import (
	"shareInviteCode/model"

	"github.com/lazyfury/go-web-template/controller"
)

func NewActivityController() *controller.Controller {
	return &controller.Controller{
		DB:    model.DB,
		Model: &model.ActivityModel{},
		Auth:  Auth,
	}
}
