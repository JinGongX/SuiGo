<script  lang="ts" setup>

import { getCurrentInstance ,computed, ref } from 'vue'
import { ElMessage } from 'element-plus'
import {
  Document,
  Menu as IconMenu,
  Location,
  Setting,
  HomeFilled,
  House,
  Avatar,Promotion,Tools,RemoveFilled,Comment,ChatDotRound,Grid,Search,Connection,Message,Plus,
} from '@element-plus/icons-vue'
import type { DropdownInstance,UploadProps  } from 'element-plus'
import axios from 'axios'
import { useRouter } from 'vue-router';

// 获取当前实例
const instance = getCurrentInstance();
const myGlobalVarHost = instance?.appContext.config.globalProperties.$globalHost;
//console.log("testhaha:"+myGlobalVarHost)
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
         url: myGlobalVarHost+"api/v1/GetSeachs?title="+event,//"http://127.0.0.1:2023/api/v1/GetSeachs?title="+event,
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
         url: myGlobalVarHost+"api/v1/GetSeachs?title="+event.target.value,//"http://127.0.0.1:2023/api/v1/GetSeachs?title="+event.target.value,
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

const containerRef = ref<HTMLElement | null>(null)

const handleClick = (e: MouseEvent) => {
  e.preventDefault()
}

const imageUrl = ref('')

const handleAvatarSuccess: UploadProps['onSuccess'] = (
  response,
  uploadFile
) => {
  imageUrl.value = URL.createObjectURL(uploadFile.raw!)
}

const beforeAvatarUpload: UploadProps['beforeUpload'] = (rawFile) => {
  if (rawFile.type !== 'image/jpeg') {
    ElMessage.error('Avatar picture must be JPG format!')
    return false
  } else if (rawFile.size / 1024 / 1024 > 2) {
    ElMessage.error('Avatar picture size can not exceed 2MB!')
    return false
  }
  return true
}
const router = useRouter();
const btexit= () => { 
    //console.log(event.target.value) 
   if(!confirm("确定退出?")){
      return 
   }else{
      // 清除所有记录 
     sessionStorage.clear();
     router.push("Account")
   }
}
const onhomeload= () => {  
     router.push("Blogs") 
}
</script>

<template>
    <div class="common-layout">
        <el-container> 
            <el-header v-if="!$route.meta.showNav">
                 <el-menu  class="el-menu-demo" mode="horizontal"  :ellipsis="false" style="background: rgb(240 243 70 / 80%);height: 60px;overflow: hidden;text-overflow: ellipsis;" > 
                    <!--:default-active="activeIndex" class="el-menu-demo" mode="horizontal" :ellipsis="false"
                    @select="handleSelect" -->
                    <!-- <el-menu-item index="0" style="height:auto"><img alt="Vue logo" class="logo" src="../../assets/sui.png"
                            width="38" height="38" /><span
                            style="font-family: cursive;display: block;font-size: 24px; vertical-align: bottom;width: 24px; height: auto;">随</span>
                    </el-menu-item> -->
                    <div class="menu_convenient_top" @click="onhomeload" style="cursor: pointer;overflow: hidden;text-overflow: ellipsis;">
                      
                        <img alt="Vue logo" class="logo" src="../../assets/logo1.svg"
                            width="38" height="38" /><span
                            style="font-family: cursive;display: block;font-size: 24px; vertical-align: bottom;width: 24px; height: auto;float: right;">随</span> 
                    </div>
                    <!-- <div class="menu_convenient">
                        <router-link to="./Blogs">
                            <el-icon title="随记"  :size="27" style="cursor:pointer" >
                            <Grid /></el-icon> <span class="menu_font">随记</span>
                        </router-link> 
                    </div>  -->
                    <!-- <div class="menu_convenient">
                          <router-link to="./ExcalidrawWrapper" target="_blank" rel="opener"> 
                        <el-icon title="白板"  :size="27" style="cursor:pointer" >
                        <EditPen /></el-icon>  <span class="menu_font">白板</span>  </router-link>
                    </div> 
                    <div class="menu_convenient">
                      <router-link to="./BookInfo">
                    <el-icon title="读书"  :size="27" style="cursor:pointer" >  <Notebook /></el-icon> 
                     <span class="menu_font">读书</span>  </router-link>
                  </div>
                        <div class="menu_convenient">
                            <router-link to="./SoftInfo">
                        <el-icon title="开源软件"  :size="27" style="cursor:pointer" >
                        <Connection /></el-icon>  <span class="menu_font">开源软件</span>  </router-link>
                    </div> -->
                   
                    
                    <div class="flex-grow" style="line-height: 60px;font-family: cursive;color: #8d8989;overflow: hidden;text-overflow: ellipsis;">欲买桂花同载酒,终不似,少年游</div> 
                    <!-- <el-menu-item index="1">
                       
                    </el-menu-item> -->
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
                        <!-- 搜索结果下拉列表 -->
                       <!-- <el-dropdown v-model="isDropdownVisible" >
                         <el-dropdown-menu slot="dropdown">
                           <el-dropdown-item v-for="(result, index) in searchResults" :key="index" :command="result.id">
                             {{ result.name }}
                           </el-dropdown-item>
                         </el-dropdown-menu>
                       </el-dropdown> -->
                         <!-- <span class="el-dropdown-link"> Dropdown List1 </span> -->
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
                                    <!-- {{o["title"]}} -->
                                  </div>
                                </div> 
                                  <div v-if="sesrchcount>=10" class="sj_result_select_conter_li">
                                    相关结果大于10条记录，查看更多....
                                  </div>
                                  <div class="sj_result_select_foot">
                                 <!-- 其他 -->
                              </div>
                              </div>
                              <div class="sj_result_select_foot">
                                 <!-- 其他 -->
                              </div>
                            </div>
                             <!-- <el-dropdown-item disabled>随记</el-dropdown-item>
                             <el-dropdown-item>Action 1</el-dropdown-item>
                             <el-dropdown-item>Action 2</el-dropdown-item>
                             <el-dropdown-item divided>Action 3</el-dropdown-item>
                             <el-dropdown-item disabled>Action 4</el-dropdown-item>
                             <el-dropdown-item divided>Action 5</el-dropdown-item> -->
                           </el-dropdown-menu>
                         </template>
                       </el-dropdown>
                        <!-- <div id="seach_result" :style="disstyle"   class="seach_result">

                        </div> -->
                    </div>
                    <div class="menu_loge"  >
                        <router-link to="./Blogs">
                            <el-icon title="首页"  :size="27" style="cursor:pointer" >
                              <svg t="1724855793595" class="icon" viewBox="0 0 1152 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="5720" id="mx_n_1724855793595" width="200" height="200"><path d="M583.807877 778.756066c-84.099572 0-152.514558-70.908769-152.514558-158.092758s68.414986-158.059945 152.514558-158.059945 152.481745 70.908769 152.481744 158.092758-68.414986 158.059945-152.481744 158.059945z m0-261.453527c-53.944478 0-97.848195 46.364688-97.848196 103.360769s43.903717 103.393582 97.848196 103.393583 97.68413-46.364688 97.68413-103.36077-43.772466-103.393582-97.68413-103.393582z" fill="#00bd7e" p-id="5721"></path><path d="M234.284411 1023.967187a109.16866 109.16866 0 0 1-109.16866-109.004596v-607.03944H179.814926v607.03944a54.403859 54.403859 0 0 0 54.338233 54.338234l685.462373-3.084417a54.436672 54.436672 0 0 0 54.469485-54.371046V311.073194h54.699176v600.772168a109.16866 109.16866 0 0 1-109.037409 109.16866z" fill="#00bd7e" p-id="5722"></path><path d="M27.365994 419.060588a27.333181 27.333181 0 0 1-15.258018-50.006924L530.191528 18.873939a108.676466 108.676466 0 0 1 123.409478 0.689072l487.469077 337.251425a27.346306 27.346306 0 0 1-31.10667 44.986544L622.42871 64.516742a54.371046 54.371046 0 0 0-61.688332-0.328129L42.656826 414.368337a27.365994 27.365994 0 0 1-15.290832 4.692251z" fill="#00bd7e" p-id="5723"></path></svg>
                            </el-icon>
                        </router-link> &nbsp;&nbsp;
                        <router-link to="./Mkdown" target="_blank" rel="opener"><el-icon title="发表随记"  :size="27" style="cursor:pointer" >
                            <Promotion /></el-icon>
                        </router-link> &nbsp;&nbsp;
                        <router-link to="./Chatsui">
                            <el-icon title="随问"  :size="27" style="cursor:pointer"><ChatDotRound /></el-icon>
                        </router-link>
                        <!-- <router-link to="./Chatsui">  
                            <el-icon  title="消息"  :size="27" style="cursor:pointer"><Message /></el-icon> 
                        </router-link>-->
                        <!-- <router-link to="./Chatsui"> <el-icon  title="功能"  :size="27" style="cursor:pointer"><Compass /></el-icon></router-link> -->
                    </div>
                    <el-sub-menu index="2">
                        <template #title><img alt="photo" class="photo" :src=suiphoto /> {{ suiname }}</template>
                        <el-menu-item index="2-1" @click="visible = true"> <el-icon><Avatar /></el-icon><span>个人资料</span></el-menu-item>
                        <!-- <el-menu-item index="2-2"><el-icon><Promotion /></el-icon><span>随记</span></el-menu-item>
                        <el-menu-item index="2-3"> <el-icon><Tools /></el-icon><span>设置</span></el-menu-item> -->
                        <el-menu-item index="2-4" @click="btexit"> <el-icon><RemoveFilled /></el-icon><span>退出</span></el-menu-item>
                        <el-menu-item index="2-5" > <el-icon><Opportunity /></el-icon><span>关于</span></el-menu-item>
                        <!-- <el-sub-menu index="2-4">
                            <template #title>item four</template>
                            <el-menu-item index="2-4-1">item one</el-menu-item>
                            <el-menu-item index="2-4-2">item two</el-menu-item>
                            <el-menu-item index="2-4-3">item three</el-menu-item>
                        </el-sub-menu> -->
                    </el-sub-menu>
                </el-menu>
            </el-header>
            <el-container> 
                <el-container>
                    <el-main><router-view></router-view></el-main>
                    <!-- <el-footer>Footer</el-footer> -->
                </el-container>
            </el-container>
        </el-container>
        <el-dialog v-model="visible" :show-close="false">
    <template #header="{ close, titleId, titleClass }">
      <div class="my-header">
        <h5 :id="titleId" :class="titleClass">个人资料</h5>
        <el-button type="danger" @click="close">
          <el-icon class="el-icon--left"><CircleCloseFilled /></el-icon>
          
        </el-button>
      </div>
    </template>
    
    <div class="myinfo-content">
      <div class="myinfo-content-left">
        <!-- <ul class="myinfo_menu_ul">
          <li>头像&昵称</li>
          <li>随记</li>
          <li>密码</li>
          <li>其他</li>
        </ul> -->
        <el-anchor  :container="containerRef"
          direction="vertical"
          bound="-15"
          type="underline"
          :offset="30"
          @click="handleClick">
          <!-- <el-link href="#part1" title="头像&昵称"/> 
          <el-link href="#part2" title="随记"/> 
          <el-link href="#part3" title="密码"/>  
          <el-link href="#part4" title="其他"/>   -->
          <el-anchor-link href="#part1" title="头像&昵称" >头像&昵称</el-anchor-link>
          <el-anchor-link href="#part2" title="随记" >随记</el-anchor-link>
          <el-anchor-link href="#part3" title="密码" >密码</el-anchor-link>
          <el-anchor-link href="#part3" title="其他"> 其他 </el-anchor-link>
        </el-anchor>
      </div>
      <div class="myinfo-content-right" ref="containerRef"> 
        <div class="myinfo-content-photo" id="part1">
             <div class="myinfo-content-title">
                 个人信息
             </div>
             <div class="myinfo-content-body">
                  <div class="myinfo-content-body-left">
                      <div class="myinifo-content-body-title">昵称：</div>
                      <div><el-input 
                          class="w-20 m-2" 
                          placeholder="请输入昵称 " 
                        /> </div>
                      <div class="myinifo-content-body-title">个人简介：</div>
                      <div>
                        <el-input  v-model="textarea_desc" style="width: 100%" :rows="2"  type="textarea"  />
                       </div>
                  </div>
                  <div class="myinfo-content-body-right">
                    <div style="width:100%;margin-top:10px;margin-bottom:6px;">
                      <!-- <img class="myinifo-photo"   src="../../assets/sysimg/Emay.png" />  -->
                      <el-upload
                        class="avatar-uploader"
                        action="https://run.mocky.io/v3/9d059bf9-4660-45f2-925d-ce80ad6c4d15"
                        :show-file-list="false"
                        :on-success="handleAvatarSuccess"
                        :before-upload="beforeAvatarUpload"
                      >
                        <img v-if="imageUrl" :src="imageUrl" class="avatar" />
                        <el-icon v-else class="avatar-uploader-icon"><Plus /></el-icon>
                      </el-upload>
                    </div>
                    <div style="font-size:12px"><el-button  size="small" ><el-icon class="el-icon--right"><Upload /></el-icon>更换头像</el-button> </div>
                  </div>
             </div>
        </div>
        <div class="myinfo-content-sui" id="part2">
            <div class="myinfo-content-title">
              随记
            </div>
            <div class="myinfo-content-body">
              <div style="margin-top: 20px;height:40px;">
               <el-radio-group  v-model="suiradio1" size="small">
                 <el-radio label="1" border>公开</el-radio>
                 <el-radio label="2" border>仅我可见</el-radio>
               </el-radio-group>
             </div>
             <div class="mt-4">
              <el-checkbox-group v-model="checkboxGroup1" size="small">
                <el-checkbox label="默认" value="Value1" border />
                <el-checkbox label="计算机" value="Value2" border />
                <el-checkbox label="工程" value="Value3" border />
                <el-checkbox label="经济" value="Value4" border />
              </el-checkbox-group>
             </div>
            </div>
        </div>
        <div class="myinfo-content-password" id="part3">
          <div class="myinfo-content-title">
              密码
            </div>
            <div class="myinfo-content-body">
              <div style="margin-top: 20px;height:40px;">
                <el-button type="warning"><el-icon class="el-icon--left"><Message /></el-icon> 通过辅助邮箱更换密码</el-button>
              </div>
            </div>
           </div>
        <div class="myinfo-content-other" id="part4">
          <div class="myinfo-content-title">
             其他
           </div>
           <div class="myinfo-content-body">
            <div style="margin-top: 20px;height:40px;"></div>
            </div> 
        </div>
      </div>
    </div>
  </el-dialog>
    </div>
</template>
 
<style scoped>
.avatar-uploader .avatar {
  width:60%;
  min-width:80px;
  min-height:80px;
  display: block;
}
</style>
<style>
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
   margin-top:-10px;
   
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

.avatar-uploader .el-upload {
  border: 1px dashed var(--el-border-color);
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  transition: var(--el-transition-duration-fast);
}

.avatar-uploader .el-upload:hover {
  border-color: var(--el-color-primary);
}

.el-icon.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  min-width:80px;
  min-height:80px;
  width:60%;
  text-align: center;
}
</style>