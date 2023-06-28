package mysql

import (
	"bluebell_backend/models"
	"bluebell_backend/pkg/snowflake"
	"crypto/md5"
	"encoding/hex"
	"gorm.io/gorm"
)

const secret = "liwenzhou.com"

func encryptPassword(data []byte) (result string) {
	h := md5.New()
	h.Write([]byte(secret))
	h.Write(data)
	//这行代码计算哈希值，并将其转换为十六进制字符串格式，作为函数的返回值
	return hex.EncodeToString(h.Sum(nil))
}
func Register(user *models.User) error {
	var u models.User
	if err := DB.Model(&models.User{}).Select("username,user_id").Where("username = ?", user.UserName).First(&u).Error; err != nil {
		return err
	}

	if u.UserID > 0 {
		// User already exists
		return ErrorUserExit
	}

	// Generate user_id
	userID, err := snowflake.GetID()
	if err != nil {
		return ErrorGenIDFailed
	}

	// Generate encrypted password
	password := encryptPassword([]byte(user.Password))

	// Insert user into the database
	err = DB.Create(&models.User{
		UserID:   userID,
		UserName: user.UserName,
		Password: password,
	}).Error

	return err
}

func Login(user *models.User) error {
	originPassword := user.Password // Record the original password

	var existingUser models.User
	if err := DB.Where("username = ?", user.UserName).First(&existingUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// User does not exist
			return ErrorUserExit
		}
		// Error querying the database
		return err
	}

	// Compare encrypted passwords
	password := encryptPassword([]byte(originPassword))
	if existingUser.Password != password {
		return ErrorPasswordWrong
	}

	// Update the user object with the retrieved data
	user.UserID = existingUser.UserID
	user.UserName = existingUser.UserName
	user.Password = existingUser.Password

	return nil
}

func GetUserByID(idStr string) (*models.User, error) {
	user := new(models.User)
	if err := DB.Select("user_id, username").Where("user_id = ?", idStr).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// User not found
			return nil, ErrorUserNotExit
		}
		// Error querying the database
		return nil, err
	}

	return user, nil
}
