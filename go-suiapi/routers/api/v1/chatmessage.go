package v1

import (
	"gotest20230626/models"
	"gotest20230626/pkg/e"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetChatMessage(c *gin.Context) {
	sendid := com.StrTo(c.Query("sendid")).MustInt()
	reid := com.StrTo(c.Query("reid")).MustInt()
	//data := make(map[string]interface{})
	code := e.SUCCESS
	//var data interface{}
	data := models.GetChatMessage(sendid, reid)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func GetGroupChatMessage(c *gin.Context) {
	sendid := com.StrTo(c.Query("sendid")).MustInt()
	gpid := com.StrTo(c.Query("gpid")).MustInt()
	//data := make(map[string]interface{})
	code := e.SUCCESS
	//var data interface{}
	data := models.GetGroupChatMessage(sendid, gpid)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
