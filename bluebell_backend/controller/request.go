package controller

import "github.com/gin-gonic/gin"

func getCurrentUserID(c *gin.Context) (userID uint64, err error) {
	_userID, ok := c.Get("userid")
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = _userID.(uint64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
