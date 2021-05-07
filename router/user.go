package router

import (
	"github.com/lazyfure/thefireseed/utils"

	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	utils.Render(c, "login")
}
func reg(c *gin.Context) {
	utils.Render(c, "reg")
}

func profile(c *gin.Context) {
	utils.Render(c, "home/user/profile.html", utils.UserParam{})
}
