package controller

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/lazyfure/thefireseed/middleware"
	"github.com/lazyfure/thefireseed/model"

	"github.com/gin-gonic/gin"
	"github.com/lazyfury/go-web-template/controller"
	"github.com/lazyfury/go-web-template/response"
	"github.com/lazyfury/go-web-template/tools/crypto"
	"github.com/lazyfury/go-web-template/tools/types"
	"gorm.io/gorm"
)

type UserController struct {
	controller.Controller
}

// NewUserController NewUserController
func NewUserController() *UserController {
	return &UserController{
		controller.Controller{
			DB:    model.DB,
			Model: &model.User{},
			Auth:  Auth,
		},
	}
}

// Install Install
func (u *UserController) Install(g *gin.RouterGroup, path string) {
	controller.Install(g, u, path)
	g.POST("/reg", u.Reg)
	g.POST("/login", u.Login)
	g.GET("/logout", u.Logout)
	g.GET("/user-profile", u.Profile)
}

func (u *UserController) Profile(c *gin.Context) {
	user := GetUserOrLogin(c)
	id := strconv.Itoa(int(user.ID))
	c.Params = []gin.Param{{Key: "id", Value: id}}
	u.Detail(c)
}

func (u *UserController) Logout(c *gin.Context) {
	if response.IsReqFromHTML(c) {
		c.SetCookie("token", "", -1, "/", "", false, true)
		c.Redirect(http.StatusFound, "/")
		return
	}

	c.JSON(http.StatusOK, response.JSONSuccess("退出登录", nil))
}

// 登录
func (u *UserController) Login(c *gin.Context) {
	user := &model.User{}
	if err := c.Bind(user); err != nil {
		response.Error(response.JSON(response.InvalidJSONData, "", err))
	}
	// check
	user.Name = strings.Trim(user.Name, " ")
	if user.Name == "" {
		response.Error(response.JSONError("用户名不可空", nil))
	}
	if user.Password == "" {
		response.Error(response.JSONError("用户密码不可空", nil))
	}

	// 用昵称查找
	var find = &model.User{}
	if err := u.DB.GetObjectOrNotFound(find, map[string]interface{}{
		"name": user.Name,
	}, func(db *gorm.DB) *gorm.DB {
		return db.Or(map[string]interface{}{
			"email": user.Name,
		})
	}); err != nil {
		response.Error(response.JSON(response.NotFound, "用户不存在", nil))
	}

	// 比对密码
	user.Password = crypto.SHA256String(user.Password)
	if find.Password != user.Password {
		response.Error(response.JSON(response.AuthedError, "用户密码错误", nil))
	}

	u.updateUserInfo(c, find, find.Status)
	// 更新用户信息
	if err := u.DB.Updates(find).Error; err != nil {
		response.Error(err)
	}

	str, _ := middleware.CreateToken(*find)
	if response.IsReqFromHTML(c) {
		c.SetCookie("token", str, 86400, "/", "", false, true)
		c.Redirect(http.StatusFound, "/")
		return
	}
	c.JSON(http.StatusOK, response.JSON(response.LoginSuccess, "", str))
}

// 注册
func (u *UserController) Reg(c *gin.Context) {
	user := &model.User{}

	if err := c.ShouldBind(user); err != nil {
		response.Error(response.JSON(response.InvalidJSONData, "", err))
	}

	if err := user.Validator(); err != nil {
		response.Error(err)
	}

	u.updateUserInfo(c, user, 1)
	user.SetCode()

	// 创建用户
	if err := u.DB.Create(user).Error; err != nil {
		response.Error(err)
	}

	str, _ := middleware.CreateToken(*user)
	if response.IsReqFromHTML(c) {
		c.SetCookie("token", str, 86400, "/", "", false, true)
		c.Redirect(http.StatusFound, "/")
		return
	}

	c.JSON(http.StatusCreated, response.JSON(response.StatusCreated, "注册成功", &struct {
		*model.User
		A string `json:"password,omitempty"`
	}{User: user}))
}

func (u *UserController) updateUserInfo(c *gin.Context, user *model.User, status int) {
	user.IP = c.ClientIP()
	user.Ua = c.Request.UserAgent()
	user.Status = status //数据库忘记default了
	user.LoginTime = types.LocalTime{Time: time.Now()}

}
