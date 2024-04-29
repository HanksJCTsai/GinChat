package service

import (
	"ginchat/models"
	"ginchat/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUserList
// @Tags User
// @Success 200 {string} json{"code","message"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	userdatas := make([]*models.UserBasic, 100)
	userdatas = models.GetUserList(utils.DB)
	c.JSON(http.StatusOK, gin.H{
		"message": userdatas,
	})
}
