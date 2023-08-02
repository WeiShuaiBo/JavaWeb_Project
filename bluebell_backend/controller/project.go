package controller

import (
	"bluebell_backend/dao/mysql"
	"bluebell_backend/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateProject(c *gin.Context) {
	var project models.Project
	err := c.ShouldBind(&project)
	fmt.Println(project)
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
		zap.L().Info("项目1创建成功，正在向项目2过渡")
		return
	}
}

func CreateProject2(c *gin.Context) {
	var p models.ProjectDetail
	if err := c.ShouldBind(&p); err != nil {
		zap.L().Error("createProject2() 方法绑定参数失败，请重新尝试" + err.Error())
		return
	}
	err, flag := mysql.MCreateProject2(p)
	if flag {
		ResponseSuccess(c, p)
	} else {
		ResponseErrorWithMsg(c, CodeError, err.Error())
	}
}
