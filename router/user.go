package router

import (
	"shareInviteCode/utils"

	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	utils.Render(c, "login")
}
func reg(c *gin.Context) {
	utils.Render(c, "reg")
}
