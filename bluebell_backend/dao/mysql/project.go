package mysql

import (
	"bluebell_backend/models"
	"database/sql"
	"go.uber.org/zap"
)

func MCreateProject(p *models.Project) (error, bool) {
	sqlStr := `SELECT username,user_status FROM user WHERE username=?`
	userName := p.UserName
	var previousStatus string
	var test1 *models.User
	if err := db.Get(&test1, sqlStr, userName); err == sql.ErrNoRows {
		zap.L().Error("该用户尚未注册信息，请先完成注册。")
		return err, false
	}

	previousStatus = test1.Status
	if previousStatus == "未申请" {
		sqlStr1 := `insert into project(
	           project_user_name,project_university,project_college,project_major,project_email,project_phone,project_projectdirection)
		values(?,?,?,?,?,?,?)`
		if _, err := db.Exec(sqlStr1, p.UserName, p.University, p.College, p.Major, p.Email, p.Phone, p.ProjectDirection); err != nil {
			zap.L().Error("insert project failed", zap.Error(err))
			return err, false
		}
	}
	return nil, true
}
