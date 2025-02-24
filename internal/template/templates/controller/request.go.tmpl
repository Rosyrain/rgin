package controller

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	CtxUserIDKey = "userID"
	CtxUserName  = "username"
)

var ErrorUserNotLogin = errors.New("用户未登录")

// GetCurrentUser  获取当前登录的用户ID
func GetCurrentUser(c *gin.Context) (userID int64, username string, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	uname, ok := c.Get(CtxUserName)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	username, ok = uname.(string)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

func GetPageInfo(c *gin.Context) (int64, int64) {
	//获取分页参数
	offsetStr := c.Query("offset")
	limitStr := c.Query("limit")

	var (
		offset int64
		limit  int64
		err    error
	)
	offset, err = strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		offset = 1
	}

	limit, err = strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		limit = 10
	}
	return offset, limit
}
