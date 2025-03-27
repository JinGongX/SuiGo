import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
//import Login from './components/Account/Login.vue'
import axios from 'axios'
import VueAxios from 'vue-axios'
//import * as router from 'vue-router'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import router from './router/index' 

import * as EleIcons from '@element-plus/icons-vue'

//import VueNativeSock from 'vue-native-websocket'


const app=createApp(App)
//createApp(App).mount('#app')
// 配置WebSocket连接
// Vue.use(VueNativeSock, 'ws://localhost:2023/ws', {
//   format: 'json',
//   reconnection: true, // 如果连接断开，是否自动重新连接
// })
for(const name in EleIcons){
    app.component(name,EleIcons[name])
}
app.config.globalProperties.$globalHost="http://127.0.0.1:2023/"
app.config.globalProperties.$appglobalHost="http://127.0.0.1:2024/"
app.use(VueAxios,axios)
app.use(router)
app.use(ElementPlus)
app.mount('#app')
