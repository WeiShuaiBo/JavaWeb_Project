package controller

import (
	"bluebell_backend/dao/mysql"
	"bluebell_backend/dao/redis"
	"bluebell_backend/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func CreateProject(c *gin.Context) {
	var project models.Project
	err := c.ShouldBind(&project)
	userid, _ := c.Get("userid")
	userId := fmt.Sprintf("%v", userid)
	projectJSON, err := json.Marshal(&project)
	fmt.Println(projectJSON)
	err2 := redis.Client.Set("bluebell:captain:project1:"+userId, projectJSON, 7*24*time.Hour).Err()
	if err2 != nil {
		zap.L().Error("将项目数据存储到redis中出错了")
		ResponseError(c, CodeError)
		return
	}
	err1 := redis.Client.Set("bluebell:captain:Dui", project.UserName, 7*24*time.Hour).Err()
	if err1 != nil {
		zap.L().Error("redis存储出错")
		return
	}
	if err != nil {
		zap.L().Error("将项目参数绑定到结构体中失败，请你重新尝试")
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	err, flag := mysql.MCreateProject1(&project)
	if flag {
		ResponseSuccess(c, project)
		zap.L().Info("项目表单已申请，进入页面填写具体信息")
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
	userid, _ := c.Get("userid")
	userId := fmt.Sprintf("%v", userid)
	redis.Client.Set("bluebell:captain:project2:"+userId, p, 7*24*time.Hour)
	//userid, _ := c.Get("userid")
	username, _ := redis.Client.Get("bluebell:captain:Dui").Result()
	err, flag := mysql.MCreateProject2(p, username)
	if flag {
		zap.L().Info("项目2创建成功")
		var project1 models.Project
		userid1, _ := c.Get("userid")
		userId1 := fmt.Sprintf("%v", userid1)
		p1, _ := redis.Client.Get("bluebell:captain:project1:" + userId1).Result()
		err1 := json.Unmarshal([]byte(p1), &project1)
		if err1 != nil {
			zap.L().Error("将字符串解码进去结构体失败")
			return
		}
		p3 := &models.ProjectData{
			Name:             project1.UserName,
			University:       project1.University,
			College:          project1.College,
			Major:            project1.Major,
			Email:            project1.Email,
			Phone:            project1.Phone,
			ProjectDirection: project1.ProjectDirection,
			ProjectSort:      p.ProjectDetailSort,
			ProjectName:      p.ProjectDetailName,
			Member:           p.ProjectDetailPerson,
			Introduction:     p.ProjectDetailIntro,
			Creativity:       p.ProjectDetailIdea,
			Advantage:        p.ProjectDetailAdv,
			Instructor:       p.ProjectDetailTea,
			Status:           "待审批",
		}
		fmt.Println(p3)
		err2 := mysql.DB.Debug().Create(&p3).Error
		if err2 != nil {
			zap.L().Error("创建项目信息列表失败")
			return
		}
		err3 := mysql.DB.Table("user").Debug().Where("user_id=?", userId1).Update("user_status", "已申请").Error
		if err3 != nil {
			zap.L().Error("修改用户申请状态失败")
			ResponseError(c, CodeError)
			return
		}
		ResponseSuccess(c, p)
		return
	} else {
		zap.L().Error("该用户已经创建过对应的项目了")
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}
}
func ListProject(c *gin.Context) {
	var ps []*models.ProjectData
	err := mysql.DB.Debug().Find(&ps).Error
	if err != nil {
		zap.L().Error("获取项目列表失败")
		ResponseError(c, CodeError)
		return
	}

	zap.L().Info("获取项目信息成功")
	ResponseSuccess(c, &ResponseData{
		Data:    ps,
		Code:    CodeSuccess,
		Message: CodeSuccess.Msg(),
	})
	return
}

type Shenpi struct {
	ProjectName string `json:"projectname" form:"projectname"`
	TickKind    string `json:"tickKind" form:"tickKind"`
	TickReason  string `json:"tickReason" form:"tickReason"`
	Context     string `json:"context" form:"context"`
	ShenpiTime  string `json:"ticketCreateTime" form:"ticketCreateTime"`
}

func ShenPi(c *gin.Context) {
	var s Shenpi
	err := c.ShouldBind(&s)
	if err != nil {
		zap.L().Error("绑定失败")
		ResponseError(c, CodeError)
		return
	}

}
