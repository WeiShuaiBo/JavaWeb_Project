package models

import (
	"encoding/json"
	"fmt"
)

type Project struct {
	ProjectID        uint64 `db:"project_id"`
	UserName         string `json:"username" db:"project_user_name"`
	University       string `json:"university" db:"project_university"`
	College          string `json:"college" db:"project_college"`
	Major            string `json:"major" db:"project_major"`
	Email            string `json:"email" db:"project_email"`
	Phone            string `json:"phone" db:"project_phone"`
	ProjectDirection string `json:"projectDirection" db:"project_projectdirection"`
	Status           string `json:"status" db:"project_status"`
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
	p.Status = required.Status
	return nil
}
