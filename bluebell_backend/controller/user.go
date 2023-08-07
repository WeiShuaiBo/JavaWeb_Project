package controller

import (
	"bluebell_backend/dao/mysql"
	"bluebell_backend/dao/redis"
	"bluebell_backend/models"
	"bluebell_backend/pkg/jwt"
	"bluebell_backend/utils"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {
	// 拿到用户输入的注册信息 存储到 fo 结构体中
	var fo models.RegisterForm
	if err := c.ShouldBind(&fo); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParams, "注册时候，出错，无法将用户输入信息存储到结构体中")
		return
	}
	// 查询用户是否存在
	err := mysql.Register(&models.User{
		UserName: fo.UserName,
		Password: fo.Password,
	})
	if errors.Is(err, mysql.ErrorUserExit) {
		ResponseError(c, CodeUserExist)
		return
	}
	if err != nil {
		zap.L().Error("mysql.Register() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	Rd := &ResponseData{
		Code:    CodeSuccess,
		Message: CodeSuccess.Msg(),
		Data:    fo,
	}
	ResponseSuccess(c, Rd)
}

// 登录页面
func LoginHandler(c *gin.Context) {
	var u models.User
	if err := c.ShouldBind(&u); err != nil {
		zap.L().Error("用户登录时，信息绑定到结构体失败", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	fmt.Println(u)
	if err := mysql.Login(&u); err != nil {
		zap.L().Error("mysql.Login(&u) failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	if flag := utils.CaptchaVerify(u.CaptchaId); flag == false {
		zap.L().Error("验证码输入错误")
		ResponseErrorWithMsg(c, CodeParam, "验证码错误")
		return
	}
	redis.Client.Set("bluebell:userID:", u.UserID, 12*time.Hour)
	//生成Token
	aToken, rToken, _ := jwt.GenToken(u.UserID, u.UserName)
	ResponseSuccess(c, gin.H{
		"accessToken":  aToken,
		"refreshToken": rToken,
		"userID":       u.UserID,
		"username":     u.UserName,
		"code":         http.StatusOK,
	})
	c.Request.Header.Set("Authorization", aToken)
	return
}

func LoginHandler1(c *gin.Context) {
	var u models.Admin
	if err := c.ShouldBind(&u); err != nil {
		zap.L().Error("用户登录时，信息绑定到结构体失败", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	fmt.Println(u)
	err := mysql.DB.Table("admins").Where("username = ? AND pasword =?", u.UserName, u.Password).Error
	if err != nil {
		zap.L().Error("管理员登录失败")
		ResponseError(c, CodeError)
		return
	}
	err1 := mysql.DB.Table("admins").Save(&u).Error
	if err1 != nil {
		zap.L().Error("保存失败")
		ResponseError(c, CodeError)
		return
	}
	redis.Client.Set("bluebell:userID:", u.UserID, 12*time.Hour)
	//生成Token
	aToken, rToken, _ := jwt.GenToken(u.UserID, u.UserName)
	ResponseSuccess(c, gin.H{
		"accessToken":  aToken,
		"refreshToken": rToken,
		"userID":       u.UserID,
		"username":     u.UserName,
		"code":         http.StatusOK,
	})
	c.Request.Header.Set("Authorization", aToken)
	return
}

func RefreshTokenHandler(c *gin.Context) {
	rt := c.Query("refresh_token")
	// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
	// 这里假设Token放在Header的Authorization中，并使用Bearer开头
	// 这里的具体实现方式要依据你的实际业务情况决定
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		ResponseErrorWithMsg(c, CodeInvalidToken, "请求头缺少Auth Token")
		c.Abort()
		return
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		ResponseErrorWithMsg(c, CodeInvalidToken, "Token格式不对")
		c.Abort()
		return
	}
	aToken, rToken, err := jwt.RefreshToken(parts[1], rt)
	fmt.Println(err)
	c.JSON(http.StatusOK, gin.H{
		"access_token":  aToken,
		"refresh_token": rToken,
	})
}

// 展示用户信息 /api/v1/getInf
func ListUserInformation(c *gin.Context) {
	var user models.User
	var u models.UserInformation
	userid, _ := c.Get("userid")
	if err1 := mysql.DB.Table("user").Where("user_id=?", userid).First(&user).Error; err1 != nil {
		ResponseErrorWithMsg(c, CodeError, "从数据库中拿取相应的用户信息失败")
		return
	}
	u.Name = user.UserName
	u.Gender = user.Sex
	u.IdCard = user.PostId
	u.Address = user.Address
	u.Email = user.Email
	u.BirthDate = user.Birth
	zap.L().Info("成功获取用户的信息")
	ResponseSuccess(c, u)
	return
}

// 修改用户信息 /api/v1/updateInf
func UpdateUserInformation(c *gin.Context) {
	var u models.User
	var user models.UserInformation
	err := c.ShouldBind(&user)
	if err != nil {
		zap.L().Error("修改用户信息失败")
		ResponseError(c, CodeError)
		return
	}
	userid, _ := c.Get("userid")
	err1 := mysql.DB.Where("user_id=?", userid).First(&u).Error
	if err1 != nil {
		zap.L().Error("获取用户的数据失败")
		ResponseError(c, CodeError)
		return
	}
	u.UserName = user.Name
	zap.L().Info("jwt已经更新了")
	u.Sex = user.Gender
	u.PostId = user.IdCard
	u.Address = user.Address
	u.Email = user.Email
	u.Birth = user.BirthDate
	//保存更改后的数据
	err2 := mysql.DB.Debug().Table("user").Where("user_id=?", userid).Save(&u).Error
	if err2 != nil {
		zap.L().Error("保存数据失败")
		ResponseError(c, CodeError)
		return
	}
	zap.L().Info("修改成功")
	ResponseSuccess(c, CodeSuccess)
	return
}

func Uploads(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		zap.L().Error("未上传文件")
		c.JSON(http.StatusBadRequest, gin.H{"error": "未上传文件"})
		return
	}

	extName := path.Ext(file.Filename)
	allowedExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".jpeg": true,
	}
	if !allowedExtMap[extName] {
		zap.L().Error("无效的图片类型，请上传JPG、PNG或JPEG格式的图片")
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的图片类型，请上传JPG、PNG或JPEG格式的图片"})
		return
	}

	currentTime := time.Now().Format("2006-01-02")
	saveDir := fmt.Sprintf("./temp/pic/%s", currentTime)
	if err1 := os.MkdirAll(saveDir, 0755); err1 != nil {
		zap.L().Error("目录创建失败")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "目录创建失败"})
		return
	}

	savePath := path.Join(saveDir, file.Filename)

	err2 := c.SaveUploadedFile(file, savePath)
	if err2 != nil {
		zap.L().Error("图片保存失败")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "图片保存失败"})
		return
	}

	imageURL := "http://localhost:8081/static/" + currentTime + "/" + file.Filename
	ResponseSuccess(c, &ResponseData{
		Data:    imageURL,
		Code:    CodeSuccess,
		Message: CodeSuccess.Msg(),
	})
}
