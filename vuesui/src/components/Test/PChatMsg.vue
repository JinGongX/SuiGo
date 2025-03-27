<template>
   <div class="chatinfo_left">
    <div class="chatinfo_index" @click="connectWebSocket">
        <div class="chatinfo_photo"><img src="../../assets/photo.jpg"></div>
        <div class="chatinfo_msg">
            <div class="info_name">.至于.<div class="chatinfo_time">22:29</div></div>
            <div class="info_desc">hello 你好！</div>
        </div>
        
    </div>
    <div class="chatinfo_index" @click="connectWebSocket">
        <div class="chatinfo_photo"><img src="../../assets/liff.jpg"></div>
        <div class="chatinfo_msg">
            <div class="info_name">JinGong.<div class="chatinfo_time">22:29</div></div>
            <div class="info_desc">hello 你好！</div>
        </div>
        
    </div>
    <div class="chatinfo_index" @click="connectWebSocket">
        <div class="chatinfo_photo"><img src="../../assets/photo.jpg"></div>
        <div class="chatinfo_msg">
            <div class="info_name">.至于.<div class="chatinfo_time">22:29</div></div>
            <div class="info_desc">hello 你好！</div>
        </div>
        
    </div>
   </div>
   <div class="chatinfo_main">
       <div class="chat_info_head">
         <div class="chat_info_head_name">.至于.</div>
       </div>
       <div id="chat_info_conter" ref="messageContainer" class="chat_info_conter">
        <div v-for="(receivedMessage, index) in receivedMessages" :key="index">
           <div class="chat_info_s">
            <div class="chat_info_photo_s">
                <img src="../../assets/photo.jpg">
            </div>
            <div class="chat_info_msg_s">
                <div class="triangle_s"></div>
                {{ receivedMessage }}
            </div>
           </div> 

           <div class="chat_info_f">
            <div class="chat_info_photo_f">
                <img src="../../assets/photo.jpg">
            </div>
            <div class="chat_info_msg_f">
                <div class="triangle_f"></div>
                {{ receivedMessage }}
            </div>
           </div>
        </div>
       </div>
       <div class="chat_info_input">
        <el-input class="txt_inputchat_msg"
          v-model="broadcastMessage"
          :rows="4"
          :autosize="{ minRows: 3, maxRows: 4 }"
          type="textarea"
          placeholder="Please input"
        />
        <el-button type="success" style="float:right"  @click="sendBroadcast">发送</el-button>
       </div>
   </div>
</template>
<script>
import { ref } from 'vue'
const textarea = ref('')

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
        //   var element = document.getElementById("chat_info_conter");
        //   // 将滚动条滚动到底部
        //   element.scrollTop = element.scrollHeight;
        // 使用 $nextTick 来确保 Vue 已经更新 DOM
           this.$nextTick(() => {
             this.scrollToBottom();
           });
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
      scrollToBottom() {
      // 获取消息容器的引用
      const container = this.$refs.messageContainer;

      // 将滚动条滚动到底部
      container.scrollTop = container.scrollHeight;
      },
    }, 
  };

</script>
<style>
.chatinfo_left{
    width:24%;
    height:100%;
    border:1px solid #d7c5c5;
    float:left;
}
.chatinfo_index{
    width:100%;
    height:55px;
    /* border:1px solid #810505; */
}
.chatinfo_photo{
    width:22%;
    height:50px;
    float:left;
    margin-left:10px;
    margin-top:5px;
}
.chatinfo_photo img{
    width:42px;
    height:42px;
    border-radius: 5px;
}
.chatinfo_msg{
    width:67%;
    height:50px;
    float:left;
    margin-top:5px;
}
.info_name{
    width:100%;
    height:20px;
    /* border:1px solid #ccc; */
    
}
.info_desc{
    width:100%;
    height:20px;
    /* border:1px solid #ccc; */
    font-size: 13px;
    color:#848484;
}
.chatinfo_time{
   width:40px;
   height:30px;
   float:right;
   color:#848484;
   font-size: 13px;
}
.chatinfo_main{
    width:76%;
    height:100%;
    border:1px solid rgb(42, 33, 33);
    float:left;
}
.chat_info_head{
    width:100%;
    height:50px;
    border:1px solid #ccc;
}
.chat_info_head_photo {

}
.chat_info_head_name {
   width: 30%;
   height:50px;
   margin-left:20px;
   float:left;
   line-height: 50px;
   font-weight: 500;
}
.chat_info_conter{
    width:100%;
    height:70%;
    border:1px solid rgb(148, 224, 27);
    overflow-y: auto;
    max-height: 388px;
}
.chat_info_input{
    width:100%;
    height:20%;
    border:1px solid rgb(57, 61, 49);
}
.chat_info_s{
    width:100%;
    float:left;
    margin-top: 5px;
    padding-left: 10px;
} 
.chat_info_photo_s{
    width:50px;
    height:50px;
    float: left;
}
.chat_info_photo_s img{
    width:50px;
    height:50px;
    border-radius:50%;
}
.chat_info_msg_s{
    max-width: 70%;
    width:auto;
    height:auto;
    min-height:30px;
    background-color: #e6e5e5;
    float:left;
    padding-top: 2px;
    padding-left:4px;
    padding-right:4px;
    padding-bottom: 2px;
    margin-left:10px;
    margin-top:10px;
    line-height: 1.3;
}
.triangle_s {
    width: 0;
    height: 0;
    /* border-left: 8px; */
    border-right: 8px solid #e6e5e5;
    border-top: 6px solid transparent;
    border-bottom: 4px solid transparent;
    float: left;
    margin-left: -11px;
}


.chat_info_f{
    width:100%;
    float:right;
    margin-top: 5px;
    
    padding-right: 10px;
} 
.chat_info_photo_f{
    width:50px;
    height:50px;
    float: right;
}
.chat_info_photo_f img{
    width:50px;
    height:50px;
    border-radius:50%;
}
.chat_info_msg_f{
    max-width: 70%;
    width:auto;
    height:auto;
    min-height:30px;
    background-color: rgba(40, 233, 40, 0.751);
    float:right;
    padding-top: 2px;
    padding-bottom: 2px;
    padding-left:4px;
    padding-right:4px;
    margin-right:10px;
    margin-top:10px;
    line-height: 1.3;
}
.triangle_f {
    width: 0;
    height: 0;
    /* border-left: 8px; */
    border-left: 8px solid rgba(40, 233, 40, 0.751);
    border-top: 6px solid transparent;
    border-bottom: 4px solid transparent;
    float: right;
    margin-right: -11px;
}
</style>