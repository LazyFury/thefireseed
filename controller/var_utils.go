package controller

import (
	"net/http"

	"github.com/lazyfure/thefireseed/model"

	"github.com/gin-gonic/gin"
	_model "github.com/lazyfury/go-web-template/model"
	"github.com/lazyfury/go-web-template/response"

	"gorm.io/gorm"
)

func Auth(c *gin.Context, must bool) _model.Middleware {
	return func(db *gorm.DB) *gorm.DB {
		//公开的接口，列表和详情不需要验证
		if c.Request.Method != http.MethodGet {
			GetUserOrLogin(c)
		}
		return db
	}
}

// GetUserOrLogin GetUser
func GetUserOrLogin(c *gin.Context) *model.User {
	_user, exists := c.Get("user")
	if !exists {
		response.Error(response.JSON(response.AuthedError, "请先登录1", nil))
	}
	user, ok := _user.(*model.User)
	if !ok {
		response.Error(response.JSON(response.AuthedError, "请先登录", nil))
	}
	return user
}

// GetUserOrEmpty GetUser
func GetUserOrEmpty(c *gin.Context) *model.User {
	_user, exists := c.Get("user")
	if exists {
		user, ok := _user.(*model.User)
		if ok {
			return user
		}
	}
	return &model.User{}
}
