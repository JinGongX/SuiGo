package v1

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	clients        = make(map[*websocket.Conn]bool)
	clientsMutex   = sync.Mutex{}
	broadcast      = make(chan []byte)
	broadcastMutex = sync.Mutex{}
)

// Message 结构
type Message struct {
	Text string `json:"text"`
}

func GetwsConnections(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// 添加连接到连接池
	clientsMutex.Lock()
	clients[conn] = true
	clientsMutex.Unlock()

	for {
		_, p, err := conn.ReadMessage()

		if err != nil {
			log.Println(err)
			break
		}

		// 将消息发送到广播通道
		broadcastMutex.Lock()
		broadcast <- p
		broadcastMutex.Unlock()
	}
}

func GetwsMessages() {
	for {
		select {
		case message := <-broadcast:
			// 发送消息给所有连接的客户端
			clientsMutex.Lock()
			for client := range clients {
				err := client.WriteMessage(websocket.TextMessage, message)
				if err != nil {
					log.Println(err)
					client.Close()
					delete(clients, client)
				}
			}
			clientsMutex.Unlock()
		}
	}
}

// func GetPushMsg(c *gin.Context) {
// 	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	defer conn.Close()

// 	for {
// 		messageType, p, err := conn.ReadMessage()
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}

// 		// 在这里你可以处理收到的消息，然后向特定的WebSocket连接发送响应
// 		// 例如，你可以将消息广播给所有连接的客户端或实现单聊逻辑

// 		// 这里只是一个简单的示例，将接收到的消息原样发送回客户端
// 		if err := conn.WriteMessage(messageType, p); err != nil {
// 			log.Println(err)
// 			return
// 		}
// 	}
// }
