<template>
    
    <div id="showWindow" class="showWindow">
         <div id="chat_menu">
             <div class="chat_menu_index">
              <router-link to="./MsgInfo" rel="opener">
                <el-icon size="25px"><ChatRound /></el-icon>
              </router-link></div>
             <div class="chat_menu_index"><el-icon size="25px"><User /></el-icon></div>
             <div class="chat_menu_index"><el-icon size="25px"><Notification /></el-icon></div>
         </div>
         <div id="chat_main">
          <router-view  class="chat_main_index"></router-view>
          
         </div>
    </div>
    <div style="display: none;">
            <h1>WebSocket Chat</h1>
            <button @click="connectWebSocket">连接WebSocket</button>
            <div>
              <div v-for="(message, index) in chatMessages" :key="index">{{ message }}</div>
            </div>
            <input v-model="messageToSend" @keyup.enter="sendMessage" placeholder="输入要发送的消息" />
            <input v-model="destination" placeholder="输入目标标识符" />
            <button @click="sendMessage">发送消息</button>
          </div>
  </template>
  
  <script> 
  
  export default {
    data() {
      return {
        chatMessages: [],
        messageToSend: '',
        destination: '',
        socket: null,
        identifier:''
      };
    },
    methods: { 
      connectWebSocket() { 
        const urlParams = new URLSearchParams(window.location.search);
        this.identifier = urlParams.get("identifier");
        console.log("identifier:"+this.identifier)
        this.socket = new WebSocket('ws://localhost:2023/chatmsg?identifier='+this.identifier);
  
        this.socket.onopen = () => {
          console.log('WebSocket连接已打开');
        };
  
        this.socket.onmessage = (event) => {
          console.log('接收到WebSocket消息:', event.data);
          this.chatMessages.push(event.data);
        };
  
        this.socket.onerror = (error) => {
          console.error('WebSocket连接错误:', error);
        };
  
        this.socket.onclose = () => {
          console.log('WebSocket连接已关闭');
        };
      },
      sendMessage() {
        if (this.messageToSend && this.destination) {
          const messageObject = {
            text: this.messageToSend,
            destination: this.destination,
          };
          this.socket.send(JSON.stringify(messageObject));
          this.chatMessages.push(`我: ${this.messageToSend}`);
          this.messageToSend = '';
        }
      },
    },
    beforeDestroy() {
      // 在组件销毁时关闭WebSocket连接
      if (this.socket) {
        this.socket.close();
      }
    },
  };
  </script>
  <style>
  #showWindow{
    width:70%;
    height:90%;
    min-height:550px;
    margin:0 auto;
    border:1px solid red;
  }
  #chat_menu{
    width:5%;
    height:100%;
    border:1px solid blue;
    float:left;
  }
.chat_menu_index{
  width:100%;
  height:40px;
  text-align: center;
  margin-top: 10px;
}
#chat_main{
  width: 95%;
  height:100%;
  border:1px solid #ccc;
  float:left;
}
.chat_main_index{
  width:100%;
  height:100%;
  border:1px solid #bc2525;
}
  </style>
  