package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"gotest20230626/models"
	"gotest20230626/pkg/e"
)

func GetUserlist(c *gin.Context) {
	data := make(map[string]interface{})
	code := e.SUCCESS
	data["lists"] = models.GetUserlist()

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func GetFriendlist(c *gin.Context) {
	userid := com.StrTo(c.Query("userid")).MustInt()

	data := make(map[string]interface{})
	code := e.SUCCESS
	data["lists"] = models.GetFriendlist(userid)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
func GetGrouplist(c *gin.Context) {
	userid := com.StrTo(c.Query("userid")).MustInt()

	data := make(map[string]interface{})
	code := e.SUCCESS
	data["lists"] = models.GetGrouplist(userid)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func GetUserInfo(c *gin.Context) {
	userid := com.StrTo(c.Query("userid")).MustInt()
	//data := make(map[string]interface{})
	code := e.SUCCESS
	//var data interface{}
	data := models.GetUserInfo(userid)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func EditUserInfo(c *gin.Context) {
	userid := com.StrTo(c.Query("userid")).MustInt()
	username := com.StrTo(c.Query("username"))

	data := make(map[string]interface{})
	if username != "" {
		data["name"] = username
	}

	code := e.SUCCESS
	//var data interface{}
	models.EditUserInfo(userid, data)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
