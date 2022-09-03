package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 处理跨域请求,支持options访问
func GetSysUserHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, bol := c.GetQuery("id")
		if !bol {
			return
		}
		parseInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return
		}
		if err != nil {
			return
		}
		var testservice TestService
		account, err := testservice.GetSysUser(c, parseInt)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, account)
		c.Next()
	}
}

func UpdateSysUserHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, bol := c.GetQuery("id")
		if !bol {
			return
		}
		parseInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return
		}
		if err != nil {
			return
		}
		var testservice TestService
		err = testservice.UpdateSysUser(c, parseInt)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, "Success")
		c.Next()
	}
}
