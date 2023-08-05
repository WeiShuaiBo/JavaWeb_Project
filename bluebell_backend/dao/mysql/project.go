package mysql

import (
	"bluebell_backend/models"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func MCreateProject1(p *models.Project) (error, bool) {
	var test1 models.User
	result := DB.Where("username = ?", p.UserName).First(&test1)
	if result.Error == gorm.ErrRecordNotFound {
		zap.L().Error("该用户尚未注册信息，请先完成注册。")
		return errors.New("该队长尚未注册信息，要想申请必须先进行完成平台注册"), false
	}
	previousStatus := test1.Status
	if previousStatus == "未申请" {
		project := models.Project{
			UserName:         p.UserName,
			University:       p.University,
			College:          p.College,
			Major:            p.Major,
			Email:            p.Email,
			Phone:            p.Phone,
			ProjectDirection: p.ProjectDirection,
		}
		err1 := DB.Create(&project).Error
		if err1 != nil {
			zap.L().Error("create project failed", zap.Error(err1))
			return result.Error, false
		}
	}
	return nil, true
}

func MCreateProject2(p models.ProjectDetail, captain string) (error, bool) {
	var project models.ProjectDetail
	err1 := DB.Where("project_captain=?", captain).First(&models.ProjectDetail{}).Error
	if err1 != nil {
		if errors.Is(err1, gorm.ErrRecordNotFound) {
			//用户并旗下并没有参与过相关的项目组织，所以可以继续向下进行
			//然后再查询 项目类型和项目名称是否同时存在
			result := DB.Where("project_detail_sort=? AND project_detail_name=?", p.ProjectDetailSort, p.ProjectDetailName).First(&project)
			//如果记录不存在 就创建数据
			fmt.Println(p)
			if result.Error == gorm.ErrRecordNotFound {
				if err := DB.Create(&p).Error; err != nil {
					zap.L().Error("createProject2（）创建 数据失败" + err.Error())
					return err, false
				}
				return nil, true
			}
		} else {
			zap.L().Error("该用户已经参与过其他项目了")
			return nil, false
		}
	}
	return nil, true
}
