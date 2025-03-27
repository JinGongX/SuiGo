<template > 
  <el-header>
      <el-menu  class="el-menu-demo" style="height:44px;background: rgb(123 244 92 / 80%);"  mode="horizontal"  :ellipsis="false"  > 
        <div id="blogs_title"></div>
     </el-menu>
  </el-header>
 <!-- <el-breadcrumb separator="/">
   <el-breadcrumb-item :to="{ path: '/Blogs' }">随记</el-breadcrumb-item>
    
   <el-breadcrumb-item><span id="hd_title"></span></el-breadcrumb-item>
 </el-breadcrumb> -->
 <el-form class="demo-rule-form" label-width="80px"  ref="containerRef">
  
   <div id='outline' style="float:left;display: block;position: fixed;height: 85%;max-height: 85%;border-right:1px solid #e0e0e0;margin-top: 60px;"></div>
   <!-- <div id="vditor" class="vditor-reset" name="description" style="float:left;width:77%;padding-left: 300px;"></div> -->
   
 <div id="rd_bgdv_right_conter" class="read_bgdv_right_conter"> 
   <div class="dv_hd_userinfo2"> 
     <div class="dv_op_photo"><img id="op_photo" alt="photo" class="photo" src="../../assets/photo.jpg" /> </div>
     <div class="dv_username">作者: <span id='username'></span></div>
     <div class="dv_hd_title" >创建时间 :  <span id="hd_crdate"></span></div>
   </div>
   <div id="Readvditor"  class="vditor-reset" name="description" style="margin:0 auto;width:100%;max-width: 768px;"></div>
   <div class="dv_foot"></div>
 </div>
 </el-form>
</template>

<script setup>
import Vditor from 'vditor'
//import Vditor from 'vditor/dist/method.min'
import 'vditor/dist/index.css'
import axios from 'axios'
import { progressProps } from 'element-plus';
import{useRoute} from 'vue-router'
import { getCurrentInstance,computed, onMounted,reactive,ref} from 'vue';

import moment from "moment"


moment.locale('zh-cn')

const props=defineProps(['active'],['route'])
const containerRef = ref<HTMLElement | null>(null)
// 获取当前实例
const instance = getCurrentInstance();
const myGlobalVarHost = instance?.appContext.config.globalProperties.$globalHost;

let active=computed({
   get(){
       return props.active;
   }
})

let id=computed({
   get(){
       return props.active;
   }
})
var timeID=null
const IPreviewOptions={
   //"anchor":1,
   theme:{current:props.active?"dark":"light"},
   mode:"dark",
   speech:{"enable":true},
   
   after () {
       const outlineElement = document.getElementById('outline')
       const bgconter=document.getElementById("rd_bgdv_right_conter")
       if (window.innerWidth <= 768) {
          outlineElement.style.display = 'none'
          bgconter.style.left=0
          bgconter.style.width='90%'
         bgconter.style.marginLeft='5%'
         return
       }else{
         bgconter.style.left='250px'
          bgconter.style.width='80%'
          bgconter.style.marginLeft=0
       }
       if(timeID){
        clearTimeout(timeID);
       }
       timeID = setTimeout(()=>{ 
          Vditor.outlineRender(document.getElementById('Readvditor'), outlineElement)
          if (outlineElement.innerText.trim() !== '') {
            outlineElement.style.display = 'block'
            initOutline()
          }
       },500)
     },
}

 
const initOutline = () => { 
   const headingElements = [] 
   Array.from(document.getElementById('Readvditor').children).forEach((item) => {
     if (item.tagName.length === 2 && item.tagName !== 'HR' && item.tagName.indexOf('H') === 0) {
       headingElements.push(item)
      
     }
   }) 
   let toc = []
   window.addEventListener('scroll', () => {
     const scrollTop = window.scrollY
     toc = []
     headingElements.forEach((item) => {
       toc.push({
         id: item.id,
         offsetTop: item.offsetTop,
       })
     })
     
     const currentElement = document.querySelector('.vditor-outline__item--current')
     for (let i = 0, iMax = toc.length; i < iMax; i++) {
       if (scrollTop < toc[i].offsetTop - 30) {
         if (currentElement) {
           currentElement.classList.remove('vditor-outline__item--current')
         }
         let index = i > 0 ? i - 1 : 0
         document.querySelector('span[data-target-id="' + toc[index].id + '"]').classList.add('vditor-outline__item--current')
         break
       }else if(scrollTop > toc[iMax-1].offsetTop - 30){
         if (currentElement) {
           currentElement.classList.remove('vditor-outline__item--current')
         }
         if(document.querySelector('span[data-target-id="' + toc[iMax-1].id + '"]')){
           document.querySelector('span[data-target-id="' + toc[iMax-1].id + '"]').classList.add('vditor-outline__item--current')
         } 
         break
       }
     }
   })
 }
 
const initVditor = async () => {
 const data = msgdata._value// await getContent()
 Vditor.preview(
       document.getElementById('Readvditor'),
       data,IPreviewOptions
   )  
}

const routes=useRoute()
 
 onMounted(async()=>{ 
     if(msgdata._value==null){
     await getContent() 
   }
   
  await addEventListener("resize",initVditor) 
  initVditor()  
})
const msgdata=ref(null) 
const getContent = async () => {
 //const decode =await axios.get('http://127.0.0.1:2023/Getdecrypt/' + routes.query.id).then((res) => {
 //    //const decode=res.data.data;
 //    return res.data.data;
 //  })
 // 重新获取内容的逻辑，可以调用你的 API 获取最新内容
 try {
   const result = await axios.get(myGlobalVarHost+'Readarticles/'+routes.query.id,
   {
     //headers: {
     //  'token': sessionStorage.token,
     //},
   });
    document.getElementsByTagName("title")[0].innerHTML =result.data.data.title
    document.getElementById('blogs_title').innerHTML=result.data.data.title
    document.getElementById('username').innerHTML=result.data.data.account.name
    document.getElementById('op_photo').src=result.data.data.account.photo
    document.getElementById('hd_crdate').innerHTML=moment(result.data.data.created_on*1000).format("YYYY-MM-DD HH:mm") 
    msgdata.value=result.data.data.content;
   //return result.data.data.content;
 } catch (error) {
   console.error(error);
   //return ''; // 处理获取内容失败的情况，确保返回一个默认值
 }
};

</script>
<style>

.vditor-outline__item--current{
 border-left:4px solid rgb(0,189,126);
}
.el-breadcrumb{
 width: 80%;
 margin:0 0 auto 20px;
}
.photo {
   width: 40px;
   height: 40px;
   border-radius: 50%;
   margin-right: 10px;
}
/* .dv_hd_userinfo{
 width:77%;
 height:50px;
 margin:0 0 10px 300px;
} */
.dv_hd_userinfo2{
 width:80%;
 height:50px;
 max-width: 768px;
 margin:0 auto 10px auto;
}
.dv_op_photo{
 float:left;
 width:55px;
}
.dv_username{
 float:left;
 line-height:40px;
}
.dv_hd_title{
 width:300px;
 float:left;
 line-height:40px;
 margin-left:40px;
}
.vditor-reset h1, .vditor-reset h2, .vditor-reset h3, .vditor-reset h4, .vditor-reset h5, .vditor-reset h6 {
 padding-top: 60px;
 margin-top: -60px;
}
.dv_foot{
 width: 80%;
 height:120px;

}
.read_bgdv_right_conter{
 position: absolute;
 left: 250px;
 width: 80%;
 padding-top: 60px;
 margin-top: -10px;
}
#blogs_title{
 width: 100%;
 text-align: center;
 font-size: 18px;
 font-weight: bold;
 line-height: 44px;

}
.el-header {
    z-index: 999;
    position: fixed;
    width: 100%;
    --el-header-padding: 0;
}
</style>