package models

import (
	"encoding/json"
	"fmt"
)

type Project struct {
	ProjectID        uint64 `gorm:"column:id"`
	UserName         string `json:"username" gorm:"column:project_user_name"`
	University       string `json:"university" gorm:"column:project_university"`
	College          string `json:"college" gorm:"column:project_college"`
	Major            string `json:"major" gorm:"column:project_major"`
	Email            string `json:"email" gorm:"column:project_email"`
	Phone            string `json:"phone" gorm:"column:project_phone"`
	ProjectDirection string `json:"projectDirection" gorm:"column:project_projectdirection"`
}
type ProjectDetail struct {
	ProjectDetailID     uint64 `gorm:"column:id"`
	ProjectDetailName   string `json:"" gorm:"column:project_detail_name"`
	ProjectDetailPerson Person
	ProjectDetailIntro  string `gorm:"column:project_detail_intro"`
	ProjectDetailIdea   string `gorm:"column:project_detail_idea"`
	ProjectDetailAdve   string `gorm:"column:project_detail_advu"`
	ProjectDetailTea    TPerson
}
type Person struct {
	Student []string `gorm:"column:project_detail_person"`
}
type TPerson struct {
	Teacher []string `gorm:"column:project_detail_teacher"`
}

func (p *Project) UnmarshalJSON(data []byte) error {
	required := struct {
		UserName         string `json:"username"`
		University       string `json:"university"`
		College          string `json:"college"`
		Major            string `json:"major"`
		Email            string `json:"email"`
		Phone            string `json:"phone"`
		ProjectDirection string `json:"projectDirection"`
		Status           string `json:"status"`
	}{}
	err := json.Unmarshal(data, &required)
	if err != nil {
		return fmt.Errorf("将 JSON 数据解码进入结构体失败: %s", err.Error())
	}
	p.UserName = required.UserName
	p.University = required.University
	p.College = required.College
	p.Major = required.Major
	p.Email = required.Email
	p.Phone = required.Phone
	p.ProjectDirection = required.ProjectDirection
	return nil
}
func (p Project) TableName() string {
	return "project"
}
