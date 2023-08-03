package controller

import (
	"bluebell_backend/dao/mysql"
	"bluebell_backend/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateProject(c *gin.Context) {
	var project models.Project
	err := c.ShouldBind(&project)
	c.Set("captain", project.UserName)
	if err != nil {
		zap.L().Error("将项目参数绑定到结构体中失败，请你重新尝试")
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	err, flag := mysql.MCreateProject1(&project)
	if flag {
		ResponseSuccess(c, project)
		return
	} else {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		zap.L().Info("创建项目失败，失败原因：" + err.Error())
		return
	}
}

func CreateProject2(c *gin.Context) {
	var p models.ProjectDetail
	if err := c.ShouldBind(&p); err != nil {
		zap.L().Error("createProject2() 方法绑定参数失败，请重新尝试" + err.Error())
		return
	}
	username, _ := c.Get("captain")
	err, flag := mysql.MCreateProject2(p, username.(string))
	if flag {
		zap.L().Info("项目2创建成功")
		ResponseSuccess(c, p)
		return
	} else {
		zap.L().Error("该用户已经创建过对应的项目了")
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}
}
