package chatmsg

import (
	"fmt"
	"gotest20230626/models"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/unknwon/com"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	clients        = make(map[*websocket.Conn]string) // 用于存储客户端连接和其标识符的映射
	clientsMutex   = sync.Mutex{}
	groups         = make(map[string]map[*websocket.Conn]bool) // 群组ID -> 群组成员连接映射
	groupsMutex    = sync.Mutex{}                              // 保证群组数据安全
	broadcast      = make(chan Message)
	broadcastMutex = sync.Mutex{}
)

// Message 结构
type Message struct {
	//Text        string `json:"text"`
	//Destination string `json:"destination"` // 添加目标标识符字段
	//FriendId    string `json:"friendId"`
	//FriendName  string `json:"friendname"`
	Messageid string `json:"messageID"`

	Senderid      string    `json:"senderID"`
	Receiverid    string    `json:"receiverID"`
	Messagetype   string    `json:"messageType"`
	Content       string    `json:"content"`
	Mediaurl      string    `json:"mediaURL"`
	Thumbnailurl  string    `json:"thumbnailURL"`
	Senttimestamp time.Time `json:"sentTimestamp"`
	//Uno           string    `json:"uno"`
	//Senttimestamp time.Time `json:"sentTimestamp"`
	//Readtimestamp   *time.Time `json:"readTimestamp"`
	//Recalltimestamp *time.Time `json:"recallTimestamp"`
	//Isrecalled      bool       `json:"isRecalled"`
	//Isread          bool       `json:"isRead"`
	Groupid        string `json:"groupID"`
	SendGroupphoto string `json:"sendgroupPhoto"`
}

func ClientConnections(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// 从客户端请求中获取标识符
	identifier := c.Query("identifier")
	if identifier == "" {
		log.Println("客户端标识符不能为空")
		return
	}
	groupID := c.Query("groupID") // 群组ID (如果适用)

	// 添加连接到连接池
	clientsMutex.Lock()
	clients[conn] = identifier
	clientsMutex.Unlock()

	if groupID != "" {
		// 如果是群组消息，加入到群组
		groupsMutex.Lock()
		if _, ok := groups[groupID]; !ok {
			groups[groupID] = make(map[*websocket.Conn]bool)
		}
		groups[groupID][conn] = true
		groupsMutex.Unlock()
	}

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println(err)
			break
		}

		// 将消息发送到广播通道
		broadcastMutex.Lock()
		broadcast <- msg
		broadcastMutex.Unlock()
	}
}

func ClientMessages() {
	for {
		select {
		case message := <-broadcast:
			log.Println(message)
			data := make(map[string]interface{})
			data["senderID"] = com.StrTo(message.Senderid).MustInt()
			data["receiverID"] = com.StrTo(message.Receiverid).MustInt()
			data["messageType"] = message.Messagetype
			data["content"] = message.Content
			data["mediaURL"] = message.Mediaurl
			data["thumbnailURL"] = message.Thumbnailurl
			var sentime = time.Now()
			data["sentTimestamp"] = sentime
			data["groupID"] = com.StrTo(message.Groupid).MustInt()
			//data["uNO"] = com.StrTo(message.uno).MustInt()
			message.Senttimestamp = sentime
			if message.Groupid == "1" {
				crid := models.AddGroupMessage(data, sentime)
				if crid > 0 {
					AddGroupChats(message, crid)
					//message.Messageid = com.StrTo(crid).String()
				}
			} else {
				crid := models.AddMessage(data, sentime)
				if crid > 0 {
					AddChats(message, crid)
					//message.Messageid = com.StrTo(crid).String()
				}
			}

			// 发送消息给目标客户端
			clientsMutex.Lock()
			if message.Groupid != "" {
				// 群聊消息，广播到该群组的所有成员
				groupsMutex.Lock()
				if groupClients, ok := groups[message.Groupid]; ok {
					for client := range groupClients {
						fmt.Println(message.Groupid)
						//if "1" == message.Groupid || "gp1" == "gp"+message.Groupid {//groupID
						err := client.WriteJSON(message)
						if err != nil {
							log.Println("发送群聊消息失败:", err)
							client.Close()
							delete(groupClients, client)
						}
						//}
					}
				}
				groupsMutex.Unlock()
			} else {
				for client, identifier := range clients {
					fmt.Println(message.Receiverid) //+ "======" + identifier)
					if identifier == message.Receiverid || identifier == "chat"+message.Receiverid {
						err := client.WriteJSON(message)
						if err != nil {
							log.Println(err)
							client.Close()
							delete(clients, client)
						}
					}
				}
			}
			clientsMutex.Unlock()
		}
	}
}
