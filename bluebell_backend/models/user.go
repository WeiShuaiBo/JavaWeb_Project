package models

import (
	"encoding/json"
	"errors"
)

type User struct {
	UserID    uint64 `json:"user_id" db:"user_id"`
	UserName  string `json:"username" db:"username"`
	Password  string `json:"password" db:"password"`
	CaptchaId string `json:"captchaId"`
	Status    string `json:"user_status" db:"user_status"`
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
