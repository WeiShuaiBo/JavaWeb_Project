package mysql

import (
	"bluebell_backend/models"
	"bluebell_backend/pkg/snowflake"
	"crypto/md5"
	"encoding/hex"
	"gorm.io/gorm"
)

const secret = "没头脑学编程"

// MD5算法在现代密码学中已经不再被推荐使用，因为它存在一些安全性上的弱点。
// 可以考虑使用更现代的算法，如SHA-256或bcrypt。
// 对原始密码进行加密  接收一个data参数（字节切片类型）
func encryptPassword(data []byte) (result string) {
	//创建一个MD5哈希对象h 该函数可以将任意长度的数据转换为固定长度的哈希值
	h := md5.New()
	// 将一个名为secret 的字符创转换为字符切片，引入一个秘密因素，以增加密码的安全性
	h.Write([]byte(secret))
	// 同时也将 data同时写进h中，将密码数据与秘密因素一起进行哈希计算
	h.Write(data)
	//h.Sum() 获取完整的哈希结果，参数nil表示不追加任何额外的数据
	//这行代码计算哈希值，并将其转换为十六进制字符串格式，作为函数的返回值
	return hex.EncodeToString(h.Sum(nil))
}
func Register(user *models.User) error {
	//首先查询用户输入的信息是否已经存在了
	var u models.User
	DB.Model(&models.User{}).Select("username").Where("username = ?", user.UserName).First(&u)
	//如果能够从数据库中获取到用户信息，那么该用户的id应该大于0，
	if u.UserID > 0 {
		// 用户已经存在
		return ErrorUserExit
	}
	// 使用雪花算法生成用户id信息
	userID, err := snowflake.GetID()
	if err != nil {
		return ErrorGenIDFailed
	}

	// Generate encrypted password
	password := encryptPassword([]byte(user.Password))

	// Insert user into the database
	err = DB.Create(&models.User{
		UserID:    userID,
		UserName:  user.UserName,
		Password:  password,
		CaptchaId: "",
		Status:    "未申请",
	}).Error

	return err
}

func Login(user *models.User) error {
	// 记录原始密码
	originPassword := user.Password
	//根据用户名查询数据库对应数据，存储到existingUser
	var existingUser models.User
	if err := DB.Where("username = ?", user.UserName).First(&existingUser).Error; err != nil {
		//如果错误信息中包含，用户信息没找到，就返回报错
		if err == gorm.ErrRecordNotFound {
			//用户不存在
			return ErrorUserNotExit
		}
		// 返回对应错误
		return err
	}
	// 将密码加密后， 和数据库中的密码进行对比，查看是否正确
	password := encryptPassword([]byte(originPassword))
	if existingUser.Password != password {
		return ErrorPasswordWrong
	}

	// 查询已经正确，将数据库中的数据id等信息重新赋值给user结构体
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
