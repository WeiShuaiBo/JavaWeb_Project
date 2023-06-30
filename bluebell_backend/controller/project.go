package controller

import (
	"bluebell_backend/dao/mysql"
	"bluebell_backend/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateProjext1(c *gin.Context) {
	var project models.Project
	err := c.BindJSON(&project)
	if err != nil {
		zap.L().Error("将项目参数绑定到结构体中失败，请你重新尝试")
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	err, flag := mysql.MCreateProject1(&project)
	if flag {
		ResponseSuccess(c, project)
	} else {
		ResponseErrorWithMsg(c, CodeError, err.Error())

	}
}

func CreateProject2(c *gin.Context) {
	var p models.ProjectDetail
	if err := c.BindJSON(&p); err != nil {
		zap.L().Error("createProject2() 方法绑定参数失败，请重新尝试")
		return
	}
	fmt.Println(p.ProjectDetailPerson)
	fmt.Println(p.ProjectDetailTea)
	err, flag := mysql.MCreateProject2(&p)
	if flag {
		ResponseSuccess(c, p)
	} else {
		ResponseErrorWithMsg(c, CodeError, err.Error())
	}
}
func Project() {

}
