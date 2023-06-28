package mysql

import (
	"bluebell_backend/models"
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
