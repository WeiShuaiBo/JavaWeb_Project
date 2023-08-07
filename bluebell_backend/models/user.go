package models

import (
	"encoding/json"
	"errors"
)

type User struct {
	UserID    uint64 `json:"user_id" gorm:"column:user_id"`
	UserName  string `json:"user_name" gorm:"column:username"`
	Password  string `json:"user_password" gorm:"column:password"`
	Sex       string `json:"user_sex" gorm:"column:sex"`
	Birth     string `json:"birth" gorm:"column:age"`
	Address   string `json:"user_address" gorm:"column:address"`
	PostId    string `json:"user_postId" gorm:"column:postId"`
	Email     string `json:"user_email" gorm:"column:email"`
	CaptchaId string `json:"captchaId" gorm:"column:captcha_id"`
	Status    string `json:"user_status" gorm:"column:user_status"`
}
type UserInformation struct {
	Name      string `json:"name" form:"name"`
	Gender    string `json:"gender" form:"gender"`
	BirthDate string `json:"birthDate" form:"birthDate"`
	IdCard    string `json:"idCard" form:"idCard"`
	AdatarUrl string `json:"avatarUrl" form:"avatarUrl"`
	Address   string `json:"address" form:"address"`
	Email     string `json:"email" form:"email"`
}
type Admin struct {
	UserID   uint64 `json:"user_id" gorm:"column:user_id" `
	UserName string `json:"username" gorm:"column:username" form:"username"`
	Password string `json:"password" gorm:"column:password" form:"password"`
}

func (a Admin) TableName() string {
	return "admins"
}

func (u *User) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		UserName string `json:"username" db:"username"`
		Password string `json:"password" db:"password"`
		Captcha  string `json:"captchaId"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.UserName) == 0 {
		err = errors.New("缺少必填字段username")
	} else if len(required.Password) == 0 {
		err = errors.New("缺少必填字段password")
	} else if len(required.Captcha) == 0 {
		err = errors.New("缺少验证码字段captchaid")
	} else {
		u.UserName = required.UserName
		u.Password = required.Password
		u.CaptchaId = required.Captcha
	}
	return
}

type RegisterForm struct {
	UserName        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"re_password"`
}

func (r *RegisterForm) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		UserName        string `json:"username"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"re_password"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.UserName) == 0 {
		err = errors.New("缺少必填字段username")
	} else if len(required.Password) == 0 {
		err = errors.New("缺少必填字段password")
	} else if required.Password != required.ConfirmPassword {
		err = errors.New("两次密码不一致")
	} else {
		r.UserName = required.UserName
		r.Password = required.Password
		r.ConfirmPassword = required.ConfirmPassword
	}
	return
}
func (user User) TableName() string {
	return "user"
}
