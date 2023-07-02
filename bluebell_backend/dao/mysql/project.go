package mysql

import (
	"bluebell_backend/models"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func MCreateProject1(p *models.Project) (error, bool) {
	var test1 models.User
	result := DB.Where("username = ?", p.UserName).First(&test1)
	if result.Error == gorm.ErrRecordNotFound {
		zap.L().Error("该用户尚未注册信息，请先完成注册。")
		return result.Error, false
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
		if err2 := DB.Model(&models.User{}).Where("username=?", test1.UserName).Update("Status", "已申请").Error; err2 != nil {
			zap.L().Error("修改用户状态失败", zap.Error(err2))
		}
	}
	return nil, true
}

func MCreateProject2(p models.ProjectDetail) (error, bool) {
	var project models.ProjectDetail
	result := DB.Where("project_detail_sort=? AND project_detail_name=?", p.ProjectDetailSort, p.ProjectDetailName).First(&project)
	//如果记录不存在  就创建数据
	fmt.Println(p)
	if result.Error == gorm.ErrRecordNotFound {

		if err := DB.Create(&p).Error; err != nil {
			zap.L().Error("createProject2（）创建 数据失败" + err.Error())
			return err, false
		}
		return nil, true
	}
	//如果存在的话
	return result.Error, false
}
