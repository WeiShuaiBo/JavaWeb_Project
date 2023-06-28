package utils

import (
	"bluebell_backend/dao/redis"
	"bytes"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
	"time"
)

// 生成验证码
func Captcha(c *gin.Context, length ...int) {
	l := captcha.DefaultLen  //验证码长度设为默认值
	width, height := 120, 38 // 设置图片宽高
	if len(length) >= 1 {    //调整验证码的长度
		l = length[0]
	}
	if len(length) == 2 { //根据传入参数调整验证码的宽度
		width = length[1]
	}
	if len(length) == 3 { //根据传入的参数调整验证码的高度
		height = length[2]
	}
	captchaID := captcha.NewLen(l) // 生成指定长度的验证码
	redis.Client.Set("bluebell:captcha:1", captchaID, 12*time.Hour)
	_ = Serve(c.Writer, c.Request, captchaID, ".png", "zh", false, width, height)
}

// 生成图片
func Serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	w.Header().Set("Vary", "Accept-Encoding, User-Agent, timestamp")

	var content bytes.Buffer

	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		_ = captcha.WriteImage(&content, id, width, height)
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		_ = captcha.WriteAudio(&content, id, lang)
	default:
		return captcha.ErrNotFound
	}

	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))

	return nil
}

// 验证验证码
func CaptchaVerify(code string) bool {
	// 对传入的验证码值进行 URL 编码，并使用编码后的值作为 Redis 键来获取验证码值
	Values, err := redis.Client.Get("bluebell:captcha:1").Result()
	if err != nil {
		zap.L().Error("从redis中拿出验证码id失败")
		return false
	}
	if captcha.VerifyString(Values, code) {
		log.Println("验证成功")
		return true
	} else {
		log.Println("比对失败")
		return false
	}
}
