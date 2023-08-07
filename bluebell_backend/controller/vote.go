package controller

import (
	"bluebell_backend/dao/redis"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type VoteData struct {
	PostID    string  `json:"post_id"`
	Direction float64 `json:"direction"`
}

func (v *VoteData) UnmarshalJSON(data []byte) (err error) {
	var required struct {
		PostID    int    `json:"post_id"`
		Direction string `json:"direction"`
	}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return err // 返回解析错误
	}
	v.PostID = strconv.Itoa(required.PostID) // 将整数转换为字符串类型
	v.Direction, err = strconv.ParseFloat(required.Direction, 64)
	if err != nil {
		return errors.New("direction 不是有效的数字") // 返回direction格式错误
	}
	return nil // 没有错误，返回nil
}

func VoteHandler(c *gin.Context) {
	// 给哪个文章投什么票
	var vote VoteData
	if err := c.BindJSON(&vote); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNotLogin)
		return
	}
	if err1 := redis.PostVote(vote.PostID, fmt.Sprint(userID), vote.Direction); err1 != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
	return
}
