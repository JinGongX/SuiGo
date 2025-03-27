package models

import (
	"fmt"
	"time"

	_ "gorm.io/gorm"
)

// ChatMessages 结构定义
type Chatmessages struct {
	Messageid int `gorm:"primaryKey" json:"messageID"`

	Senderid        int          `json:"senderID"`
	Receiverid      int          `json:"receiverID"`
	Messagetype     string       `json:"messageType"`
	Content         string       `json:"content"`
	Mediaurl        string       `json:"mediaURL"`
	Thumbnailurl    string       `json:"thumbnailURL"`
	Senttimestamp   time.Time    `json:"sentTimestamp"`
	Readtimestamp   *time.Time   `json:"readTimestamp"`
	Recalltimestamp *time.Time   `json:"recallTimestamp"`
	Isrecalled      bool         `json:"isRecalled"`
	Isread          bool         `json:"isRead"`
	Groupid         int          `json:"groupID"`
	Tapp_Account    Tapp_Account `json:"tapp_account" gorm:"foreignkey:id;association_foreignkey:senderid"`
}

func AddMessage(data map[string]interface{}, senttime time.Time) int {
	result := db.Create(&Chatmessages{
		Senderid:      data["senderID"].(int),
		Receiverid:    data["receiverID"].(int),
		Messagetype:   data["messageType"].(string),
		Content:       data["content"].(string),
		Mediaurl:      data["mediaURL"].(string),
		Thumbnailurl:  data["thumbnailURL"].(string),
		Senttimestamp: senttime, // time.Now(),
	})
	// 检查是否有错误发生
	if result.Error != nil {
		return -1
	}

	// 获取自增ID（适用于MySQL）
	//createdID := int(result.Value.(*Chatmessages).Messageid)
	createdID := int(result.RowsAffected)

	// 尝试获取自增ID（适用于MySQL）
	if createdID == 1 {
		// 通过查询获取最后插入的自增ID
		var lastInsertID int
		db.Raw("SELECT LAST_INSERT_ID()").Row().Scan(&lastInsertID)
		createdID = lastInsertID
	}
	fmt.Println("message" + fmt.Sprint(createdID))
	return int(createdID)
}

func GetChatMessage(sendid int, reid int) (chatmessage []Chatmessages) {
	//db.Where("(senderid = ? AND receiverid = ?) or (senderid = ? AND receiverid = ?) )", sendid, reid, reid, sendid).Order("sentTimestamp desc").Find(&chatmessage) //db.Preload("Account")
	var last50MessageIDs []uint

	db.Model(&chatmessage).
		Select("messageid").
		Where("(senderid = ? AND receiverid = ?) OR (senderid = ? AND receiverid = ?)", sendid, reid, reid, sendid).
		Order("sentTimestamp DESC").
		Limit(50).
		Pluck("messageid", &last50MessageIDs)

		// 根据子查询的结果，查询完整的消息内容，并按发送时间正序排列
	db.Where("messageid IN (?)", last50MessageIDs).Order("sentTimestamp ASC").Find(&chatmessage)
	return
}

func AddGroupMessage(data map[string]interface{}, senttime time.Time) int {
	result := db.Create(&Chatmessages{
		Senderid:      data["senderID"].(int),
		Receiverid:    data["senderID"].(int),
		Messagetype:   data["messageType"].(string),
		Content:       data["content"].(string),
		Mediaurl:      data["mediaURL"].(string),
		Thumbnailurl:  data["thumbnailURL"].(string),
		Senttimestamp: senttime, // time.Now(),
		Groupid:       data["groupID"].(int),
	})
	// 检查是否有错误发生
	if result.Error != nil {
		return -1
	}

	// 获取自增ID（适用于MySQL）
	//createdID := int(result.Value.(*Chatmessages).Messageid)
	createdID := int(result.RowsAffected)

	// 尝试获取自增ID（适用于MySQL）
	if createdID == 1 {
		// 通过查询获取最后插入的自增ID
		var lastInsertID int
		db.Raw("SELECT LAST_INSERT_ID()").Row().Scan(&lastInsertID)
		createdID = lastInsertID
	}
	fmt.Println("message" + fmt.Sprint(createdID))
	return int(createdID)
}

func GetGroupChatMessage(sendid int, grid int) (chatmessage []Chatmessages) {
	//db.Where("(senderid = ? AND receiverid = ?) or (senderid = ? AND receiverid = ?) )", sendid, reid, reid, sendid).Order("sentTimestamp desc").Find(&chatmessage) //db.Preload("Account")
	var last50MessageIDs []uint
	//db.Model(&chatmessage).
	db.Model(&chatmessage).Select("messageid").
		Where("groupid = ?", grid).
		Order("sentTimestamp DESC").
		Limit(50).
		Pluck("messageid", &last50MessageIDs)

		// 根据子查询的结果，查询完整的消息内容，并按发送时间正序排列
	//db.Where("messageid IN (?)", last50MessageIDs).Order("sentTimestamp ASC").Find(&chatmessage)
	db.Preload("Tapp_Account").Where("messageid IN (?)", last50MessageIDs).Order("sentTimestamp ASC").Find(&chatmessage)
	return
}
