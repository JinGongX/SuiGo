<template>
   <el-menu  class="el-menu-demo" mode="horizontal"   :ellipsis="false"> 
    <el-menu-item index="0">
      <img
        style="width: 38px"
        src="../../assets/20230910.jpg"
        alt="Element logo"
      /><span
                            style="font-family: cursive;display: block;font-size: 24px; vertical-align: bottom;width: 24px; height: auto;float: right;">随</span>
    </el-menu-item>

    <el-menu-item index="1">Processing Center</el-menu-item>           
                   
                   <div class="menu_convenient_top">
                        <img alt="Vue logo" class="logo" src="../../assets/20230910.jpg"
                            width="38" height="38" /><span
                            style="font-family: cursive;display: block;font-size: 24px; vertical-align: bottom;width: 24px; height: auto;float: right;">随</span>
                    </div>
                    
                    <div class="menu_convenient">
                        <router-link to="./Blogs">
                            <el-icon title="随记"  :size="27" style="cursor:pointer" >
                            <Grid /></el-icon> <span class="menu_font">随记</span>
                        </router-link> 
                    </div>  
                        <div class="menu_convenient">
                        <router-link to="./SoftInfo">
                            <el-icon title="开源软件"  :size="27" style="cursor:pointer" >
                            <Connection /></el-icon>  <span class="menu_font">开源软件</span>  </router-link>
                    </div> 
                    <div class="flex-grow" />
                    <!-- <el-menu-item index="1">
                       
                    </el-menu-item> x-->
                    <div class="menu_seach" >
                        <el-dropdown ref="dropdown1" :v-model="isDropdownVisible" trigger="contextmenu" style="margin-right: 30px;margin-top:15px;">
                       
                        <el-input 
                          class="w-50 m-2"
                          v-model="inputSeach"
                          placeholder="Please Search"
                          :prefix-icon="Search" 
                          @focus="handlefocus"
                          @input="handleSearching"
                          @blur="handleblur" 
                        />  
                         <template #dropdown>
                           <el-dropdown-menu style="width:240px;">
                            <div class="sj_result_select">
                              <div class="sj_result_select_title">
                                <span><el-icon><Position /></el-icon>随记搜索 {{sesrchcount}}结果</span>
                              </div>
                              <div class="sj_result_select_conter" >
                                <div  v-for="(o, index) in  blogData.values"  :key="o">
                                  <div class="sj_result_select_conter_li">
                                    <router-link :to="{path:'/BlogDeail',query:{id:o['id']}}" target="_blank" rel="opener">{{o["title"]}}</router-link>
                                    
                                  </div>
                                </div> 
                                  <div v-if="sesrchcount>=10" class="sj_result_select_conter_li">
                                    相关结果大于10条记录，查看更多....
                                  </div>
                                  <div class="sj_result_select_foot">
                                  
                              </div>
                              </div>
                              <div class="sj_result_select_foot">
                                  
                              </div>
                            </div> 
                           </el-dropdown-menu>
                         </template>
                       </el-dropdown>

                    </div>
                   
                    <div class="menu_loge"  >
                        <router-link to="./Mkdown" target="_blank" rel="opener"><el-icon title="发表随记"  :size="27" style="cursor:pointer" >
                            <Promotion /></el-icon>
                        </router-link> &nbsp;&nbsp;
                        <router-link to="./ChatMsg">
                            <el-icon title="随问"  :size="27" style="cursor:pointer"><ChatDotRound /></el-icon>
                        </router-link>&nbsp;&nbsp;
                        <router-link to="./Chatsui">
                             
                              <el-badge is-dot class="item">
                            <el-icon  title="消息"  :size="27" style="cursor:pointer"><Message /></el-icon></el-badge>
                           
                        </router-link>
                    </div> 
                    <el-sub-menu index="3">
      <template #title>Workspace</template>
      <el-menu-item index="2-1">item one</el-menu-item>
      <el-menu-item index="2-2">item two</el-menu-item>
      <el-menu-item index="2-3">item three</el-menu-item>
      <el-sub-menu index="2-4">
        <template #title>item four</template>
        <el-menu-item index="2-4-1">item one</el-menu-item>
        <el-menu-item index="2-4-2">item two</el-menu-item>
        <el-menu-item index="2-4-3">item three</el-menu-item>
      </el-sub-menu>
    </el-sub-menu>

                    <el-sub-menu index="4">
                        <template #title>123 {{ suiname }}</template>
                        <el-menu-item index="2-1" @click="visible = true"> <el-icon><Avatar /></el-icon><span>个人资料</span></el-menu-item>
                        <!-- <el-menu-item index="2-2"><el-icon><Promotion /></el-icon><span>随记</span></el-menu-item>
                        <el-menu-item index="2-3"> <el-icon><Tools /></el-icon><span>设置</span></el-menu-item> -->
                        <el-menu-item index="2-2"> <el-icon><RemoveFilled /></el-icon><span>退出</span></el-menu-item>
                    </el-sub-menu>
                </el-menu>
  <div>
    <el-row>
      <el-col :span="18">
        <div
          style="
            height: 30px;
            width: 70%;
            background: #000;
            position: absolute;
            top: 0;
            left: 0;
            color: #fff;
          "
        >
          Fixed Top Block
        </div>
        <div ref="containerRef" style="height: 300px; overflow-y: auto">
          <div
            id="part1"
            style="
              height: 300px;
              background: rgba(255, 0, 0, 0.02);
              margin-top: 30px;
            "
          >
            part1
          </div>
          <div
            id="part2"
            style="
              height: 300px;
              background: rgba(0, 255, 0, 0.02);
              margin-top: 30px;
            "
          >
            part2
          </div>
          <div
            id="part3"
            style="
              height: 300px;
              background: rgba(0, 0, 255, 0.02);
              margin-top: 30px;
            "
          >
            part3
          </div>
        </div>
      </el-col>
      <el-col :span="6">
        <el-anchor
          :container="containerRef"
          direction="vertical"
          type="default"
          :offset="30"
          @click="handleClick"
        >
          <el-anchor-link href="#part1" title="part1" >1</el-anchor-link>
          <el-anchor-link href="#part2" title="part2" >2</el-anchor-link>
          <el-anchor-link href="#part3" title="part3" >3</el-anchor-link>
          <el-anchor-link href="#part3">
            part3
      </el-anchor-link>
        </el-anchor>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts"> 
import { computed, ref } from 'vue'
import {
  Document,
  Menu as IconMenu,
  Location,
  Setting,
  HomeFilled,
  House,
  Avatar,Promotion,Tools,RemoveFilled,Comment,ChatDotRound,Grid,Search,Connection,Message,
} from '@element-plus/icons-vue'
import type { DropdownInstance } from 'element-plus'
import axios from 'axios'


//const imgUrl = new URL('./img.png', import.meta.url).href
//const  logoimg= "@/assets/sui.png"
const activeIndex = ref('1')
const handleSelect = (key: string, keyPath: string[]) => {
    console.log(key, keyPath)
}

const handleOpen = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
  
}
const handleClose = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}
const isstyle=ref(false)
const inputSeach = ref('')
const isDropdownVisible=ref(false)
const dropdown1 = ref<DropdownInstance>()
const sesrchcount=ref(0)
const visible = ref(false)

const suiname=sessionStorage.name
const suiphoto=sessionStorage.photo
// const searchResults= [          // 搜索结果列表
//         { id: 1, name: '结果1' },
//         { id: 2, name: '结果2' },
//         { id: 3, name: '结果3' },
//         // 在这里添加更多的搜索结果
//       ] as Array<{ id: number; name: string }>
const blogData=[]

 const disstyle=computed(()=>{
    return {
        display:isstyle.value?'block':'none'
    }
 })
const handleSearching = (event) => { 
   
    console.log(event) 
    isDropdownVisible.value=true
    if (!dropdown1.value) return
       dropdown1.value.handleOpen()

    axios({
         method: "get",
         url: "http://127.0.0.1:2023/api/v1/GetSeachs?title="+event,
         headers:{
            "token":sessionStorage.token
         }
        })
        .then((res) => {
         sesrchcount.value=res.data.data.lists.length
         blogData.values=res.data.data.lists
        //  console.log(blogData.values)
     
          //  let tds=this.moments().unix(1690728099*1000)
          //  console.log("test:"+ this.moments(1691674049*1000).format("YYYY-MM-DD HH:mm:ss"))
         
    }) 
}
const handlefocus = (event) => { 
    console.log(event.target.value)
    isDropdownVisible.value=true
    if (!dropdown1.value) return
       dropdown1.value.handleOpen()

    axios({
         method: "get",
         url: "http://127.0.0.1:2023/api/v1/GetSeachs?title="+event.target.value,
         headers:{
            "token":sessionStorage.token
         }
        })
        .then((res) => {
         sesrchcount.value=res.data.data.lists.length
         blogData.values=res.data.data.lists
        //  console.log(blogData.values)
     
          //  let tds=this.moments().unix(1690728099*1000)
          //  console.log("test:"+ this.moments(1691674049*1000).format("YYYY-MM-DD HH:mm:ss"))
         
    }) 
    //isstyle.value=true
}
const handleblur = () => { 
    //console.log(event.target.value) 
    isstyle.value=false
}
const handlechang = () => { 
    //console.log(event.target.value)
    isstyle.value=false
}
const textarea_desc = ref('')
const suiradio1 = ref('1')
const checkboxGroup1 = ref(['Value1'])

 
</script>
<style>
.flex-grow {
  flex-grow: 1;
}

.el-header{
    z-index: 999;
    position: fixed;
    width: 100%; 
    --el-header-padding: 0;
}
.el-main{
    margin-top: 50px;
}
.flex-grow {
    flex-grow: 1;
}

.photo {
    width: 40px;
    height: 40px;
    border-radius: 50%;
    margin-right: 10px;
}

.el-menu--popup {
    min-width: 140px;
}

.infoimg {
    width: 24px;
    height: 24px;
    margin-right: 10px;
}
.menu_loge{
    padding-top: 17px;
    margin-right: 20px;
}
.menu_convenient{
    padding-top: 17px;
    margin-left: 20px;
    
}
.menu_convenient_top{
    padding-top: 10px;
    margin-left: 50px;
    margin-right: 20px;
}
.menu_seach{
    line-height:57px;
    margin-right:80px;
}
.menu_font{
    vertical-align: super;
}

.seach_result{
    width:240px;
    min-height: 200px;
    border:1px solid #ccc;
    position:fixed;
    margin-top: -5px;
    background-color: #f3f3f3;
    
}
.sj_result_select{
    width:100%;
    height:180px;
    /* border-bottom:1px solid #ccc; */
}
.sj_result_select_title{
    width:100%;
    height:30px;
    border-bottom:1px solid #ccc;
    line-height: 30px;
    text-indent: 10px;
}
.sj_result_select_conter{
    width:100%;
    /* height:120px;
    border:1px solid blue; */
}
.sj_result_select_conter_li{
    width:100%;
    height:30px;
    border-bottom:1px solid bisque;
    line-height: 30px;
    text-indent: 10px;
}
.sj_result_select_foot{
    width:100%;
    height:30px;

}

.my-header {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}
.item {
  margin-top: -10px;
}
.myinfo-content{
   width:100%;
   height:300px;
   margin-top:-16px;
   
}
.myinfo-content-left{
   width:24%;
   height:300px;
   border-right:1px solid #ccc;
   float:left;
}
.myinfo-content-right{
   width:72%;
   height:300px;
   overflow-y: auto;
   float:left;
   margin-left:10px;
}
.myinfo-content-photo{
  width:80%;
  height:auto;
}
.myinfo-content-title{
  width:100%;
  height:20px;
  font-size:13px;
  border-bottom: 1px dotted #ccc;
  font-weight: 600;
}
.myinfo-content-body-left{
  width:60%;
  height:120px;
  float:left;
}
.myinfo-content-body-right{
  width:40%;
  height:120px;
  float:left;
  text-align: center;
  margin-top: 16px;
}
.myinifo-photo{
  width:60%;
}
.myinifo-content-body-title{
  margin-top: 6px;
  font-size: 13px;
}
.myinfo-content-sui,.myinfo-content-password,.myinfo-content-other{
  width: 80%;
    height: auto;
    float: left;
    margin-top: 14px;
}
.myinfo_menu_ul{
  list-style-type: none;
  padding: 0;
}

/* 去掉点击后的线条 */
.el-menu--horizontal {
  .el-menu-item {
    &:focus,
    &:hover,
    &.is-active {
      outline: none;
      border-bottom: none;
    }
  }
}
 
.el-menu--vertical {
  .el-menu-item {
    &:focus,
    &:hover,
    &.is-active {
      outline: none;
      border-bottom: none;
    }
  }
}
</style>