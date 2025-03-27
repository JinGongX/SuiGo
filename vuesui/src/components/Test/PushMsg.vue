<template>
    <div>
      <h1>WebSocket Broadcast Example</h1>
      <button @click="connectWebSocket">连接WebSocket</button>
      <div v-for="(receivedMessage, index) in receivedMessages" :key="index">{{ receivedMessage }}</div>
      <input v-model="broadcastMessage" placeholder="输入要广播的消息" />
      <button @click="sendBroadcast">发送广播</button>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        receivedMessages: [],
        broadcastMessage: "",
        socket: null,
      };
    },
    methods: {
      connectWebSocket() {
        this.socket = new WebSocket("ws://localhost:2023/pushmsg");
  
        this.socket.onopen = (event) => {
          console.log("WebSocket连接已打开:", event);
        };
  
        this.socket.onmessage = (event) => {
          console.log("接收到WebSocket消息:", event.data);
          this.receivedMessages.push(event.data);
        };
  
        this.socket.onclose = (event) => {
          console.log("WebSocket连接已关闭:", event);
        };
  
        this.socket.onerror = (error) => {
          console.error("WebSocket连接错误:", error);
        };
      },
      sendBroadcast() {
        if (this.socket && this.socket.readyState === WebSocket.OPEN) {
          this.socket.send(this.broadcastMessage);
          this.broadcastMessage = ""; // 清空输入框
        }
      },
    },
  };
  </script>
  