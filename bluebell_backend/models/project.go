package models

import "encoding/json"

type Project struct {
	UserName         string `json:"username" db:"UserName"`
	University       string `json:"university" db:"University"`
	College          string `json:"college" db:"College"`
	Major            string `json:"major" db:"Major"`
	Email            string `json:"email" db:"Email"`
	Phone            int    `json:"phone" db:"Phone"`
	ProjectDirection string `json:"projectDirection" db:"ProjectDirection"`
	Status           string `json:"status" db:"Status"`
}

func (p *Project) UnmarshalJSON(data []byte) (err error) {
	requied := struct {
		UserName         string `json:"username" db:"UserName"`
		University       string `json:"university" db:"University"`
		College          string `json:"college" db:"College"`
		Major            string `json:"major" db:"Major"`
		Email            string `json:"email" db:"Email"`
		Phone            int    `json:"phone" db:"Phone"`
		ProjectDirection string `json:"projectDirection" db:"ProjectDirection"`
		Status           string `json:"status" db:"Status"`
	}{}
	err = json.Unmarshal(data, &requied)
}
