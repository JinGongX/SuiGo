package models

import "fmt"

type Chats struct {
	Chatid           int              `gorm:"primaryKey" json:"chatID"`
	User1id          int              `json:"user1ID"` //senderID
	User2id          int              `json:"user2ID"` //receiverID
	Lastmessageid    int              `json:"lastmessageID"`
	Unreadcount      int              `json:"unreadcount"`
	Chattype         int              `json:"chattype"`
	Conversationid   int              `json:"conversationID"`
	Tapp_Account     Tapp_Account     `json:"tapp_account" gorm:"foreignkey:id;association_foreignkey:user2id"`
	Chatmessages     Chatmessages     `json:"chatmessages" gorm:"foreignkey:messageid;association_foreignkey:lastmessageid"`
	Conversationinfo Conversationinfo `json:"conversationinfo" gorm:"foreignkey:groupid;association_foreignkey:conversationid"`
}

type Conversationinfo struct {
	Groupid int `gorm:"primaryKey" json:"groupID"`

	Useridid         int    `json:"userID"`
	Conversationname string `json:"conversationname"` //群聊名称
	Gpphoto          string `json:"gpphoto"`
}

func ExistChatID(id int, id2 int) bool {
	var chats Chats
	db.Where("User1id = ? AND User2id = ?", id, id2).First(&chats) //Select("Chatid").
	return chats.Chatid > 0
}

func ExistGroupChatID(gpid int) bool {
	var chats Chats
	db.Where("Conversationid = ?", gpid).First(&chats) //Select("Chatid").
	return chats.Chatid > 0
}

func AddChats(data map[string]interface{}) bool {
	if err := db.Create(&Chats{
		User1id:       data["user1ID"].(int),
		User2id:       data["user2ID"].(int),
		Lastmessageid: data["lastmessageID"].(int),
		Unreadcount:   0, //data["unreadcount"].(int),
		Chattype:      0,
	}).Error; err != nil {
		// 处理错误
		fmt.Println("Failed to create message:", err)
		return false
	}

	return true
}

func AddGroupChats(data map[string]interface{}) bool {
	if err := db.Create(&Chats{
		User1id:        data["user1ID"].(int),
		User2id:        data["user2ID"].(int),
		Lastmessageid:  data["lastmessageID"].(int),
		Unreadcount:    0, //data["unreadcount"].(int),
		Chattype:       1,
		Conversationid: data["conversationID"].(int),
	}).Error; err != nil {
		// 处理错误
		fmt.Println("Failed to create message:", err)
		return false
	}

	return true
}

func EditChats(id int, id2 int, data interface{}) bool {
	db.Model(&Chats{}).Where("Chattype='0' and ((User1id = ? AND User2id = ?) OR (User1id = ? AND User2id = ?))", id, id2, id2, id).Updates(data)

	return true
}

func EditGroupChats(gpid int, data interface{}) bool {
	db.Model(&Chats{}).Where("(Chattype='1' and  Conversationid = ?)", gpid).Updates(data)

	return true
}

func GetChatsinfo(sendid int, reid int) (chats []Chats) {
	//db.Preload("Chatmessages").Preload("Tapp_Account").Where("User1id = ? AND User2id = ?", sendid, reid).Order("Lastmessageid desc").Find(&chats)
	db.Preload("Chatmessages").Preload("Tapp_Account").Preload("Conversationinfo").Where("User1id = ?", sendid).Order("Lastmessageid asc").Find(&chats)

	return
}
