package controller

import (
	"bluebell_backend/dao/mysql"
	"bluebell_backend/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateProjext(c *gin.Context) {
	var project models.Project
	err := c.ShouldBindJSON(project)
	if err != nil {
		zap.L().Error("将项目参数绑定到结构体中失败，请你重新尝试")
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	err, flag := mysql.MCreateProject(&project)
	if flag {
		ResponseSuccess(c, project)
	}
}
