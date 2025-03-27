package chatmsg

import (
	"gotest20230626/models"
	"gotest20230626/pkg/e"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetChats(c *gin.Context) {
	sendid := com.StrTo(c.Query("sendid")).MustInt()
	reid := com.StrTo(c.Query("reid")).MustInt()
	code := e.SUCCESS
	data := models.GetChatsinfo(sendid, reid)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func PostChats(c *gin.Context) {
	sendid := com.StrTo(c.Query("sendid")).MustInt()
	reid := com.StrTo(c.Query("reid")).MustInt()
	lastmessageID := com.StrTo(c.Query("lastmessageID")).MustInt()
	code := e.SUCCESS
	data := ptAddChats(sendid, reid, lastmessageID)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
func ptAddChats(sendid int, reid int, lastmessageID int) bool {
	data := make(map[string]interface{})
	data["user1ID"] = sendid
	data["user2ID"] = reid
	data["lastmessageID"] = lastmessageID
	data["Unreadcount"] = 0
	if !models.ExistChatID(sendid, reid) {
		return models.AddChats(data)
	}
	return false
}

func AddChats(message Message, lastmessageID int) {
	data := make(map[string]interface{})
	data["user1ID"] = com.StrTo(message.Senderid).MustInt()
	data["user2ID"] = com.StrTo(message.Receiverid).MustInt()
	data["lastmessageID"] = lastmessageID
	data["Unreadcount"] = 0
	if !models.ExistChatID(data["user1ID"].(int), data["user2ID"].(int)) {
		models.AddChats(data)
		data["user1ID"] = com.StrTo(message.Receiverid).MustInt()
		data["user2ID"] = com.StrTo(message.Senderid).MustInt()
		models.AddChats(data)
	} else {
		UPdata := make(map[string]interface{})
		UPdata["lastmessageID"] = lastmessageID
		UPdata["Unreadcount"] = 0
		models.EditChats(data["user1ID"].(int), data["user2ID"].(int), UPdata)
		//data["user1ID"] = data["user2ID"].(int)
		//data["user2ID"] = data["user1ID"].(int)
		//models.EditChats(data["user2ID"].(int), data["user1ID"].(int), data)
		//models.EditChats(data["user2ID"].(int), data["user1ID"].(int), data)
	}
}

func AddGroupChats(message Message, lastmessageID int) {
	data := make(map[string]interface{})
	data["user1ID"] = com.StrTo(message.Senderid).MustInt()
	data["user2ID"] = com.StrTo(message.Receiverid).MustInt()
	data["Groupid"] = com.StrTo(message.Groupid).MustInt()
	data["lastmessageID"] = lastmessageID
	data["Unreadcount"] = 0
	//if !models.ExistGroupChatID(data["Groupid"].(int)) {
	//	models.AddGroupChats(data)
	//	//data["user1ID"] = com.StrTo(message.Receiverid).MustInt()
	//	//data["user2ID"] = com.StrTo(message.Senderid).MustInt()
	//	//models.AddGroupChats(data)
	//} else {
	if models.ExistGroupChatID(data["Groupid"].(int)) {
		UPdata := make(map[string]interface{})
		UPdata["lastmessageID"] = lastmessageID
		UPdata["Unreadcount"] = 0
		models.EditGroupChats(data["Groupid"].(int), UPdata)
		//data["user1ID"] = data["user2ID"].(int)
		//data["user2ID"] = data["user1ID"].(int)
		//models.EditChats(data["user2ID"].(int), data["user1ID"].(int), data)
		//models.EditChats(data["user2ID"].(int), data["user1ID"].(int), data)
	}
}
