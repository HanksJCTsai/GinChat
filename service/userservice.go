package service

import (
	"ginchat/models"
	"ginchat/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) {
	userdatas := make([]*models.UserBasic, 100)
	userdatas = models.GetUserList(utils.DB)
	c.JSON(http.StatusOK, gin.H{
		"users": userdatas,
	})
}
