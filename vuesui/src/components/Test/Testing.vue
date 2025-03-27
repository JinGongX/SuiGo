<template>
  <el-breadcrumb separator="/">
    <el-breadcrumb-item :to="{ path: '/Blogs' }">随记</el-breadcrumb-item>
     
    <el-breadcrumb-item><span id="hd_title"></span></el-breadcrumb-item>
  </el-breadcrumb>
  <el-form class="demo-rule-form" label-width="80px" style="margin-top: 16px;" ref="containerRef">
   
    <div id='outline' style="float:left;display: block;position: fixed;height: 85%;max-height: 85%;border-right:1px solid #e0e0e0"></div>
    <!-- <div id="vditor" class="vditor-reset" name="description" style="float:left;width:77%;padding-left: 300px;"></div> -->
   
  <div id="bgdv_right_conter" class="bgdv_right_conter"> 
    <div class="dv_hd_userinfo1"> 
      <div class="dv_op_photo"><img id="op_photo" alt="photo" class="photo" src="../../assets/photo.jpg" /> </div>
      <div class="dv_username">作者: <span id='username'></span></div>
      <div class="dv_hd_title" >创建时间 :  <span id="hd_crdate"></span></div>
      <el-button type="warning" title="分享"  @click="CopyPage" style="margin-top: 4px;"><el-icon><Share /></el-icon></el-button> 
    </div>
    <div id="vditor"  class="vditor-reset"   name="description" style="margin:0 auto;width:100%;max-width: 768px;"></div>
    <div class="dv_foot"></div>
  </div>
    <!--  <div style="width:200px;float:right;position: fixed;top: 160px;right: 0;">
     <el-anchor id="md_ancho"   :container="containerRef"
          direction="vertical"
          bound="-15"
          type="underline"
          :offset="30"
          >
        
          <el-anchor-link v-for="(o, index) in  statehh.message "   :key="o"   :href="`#${statehh.message[index]}`" :title="statehh.message[index]" ></el-anchor-link>
           
        </el-anchor>
    </div> -->
  </el-form>
   
</template>
<script setup>
import Vditor from 'vditor'
//import Vditor from 'vditor/dist/method.min'
import 'vditor/dist/index.css'
import axios from 'axios'
import { ElMessage  } from 'element-plus';
import{useRoute} from 'vue-router'
import { getCurrentInstance,computed, onMounted,reactive,ref} from 'vue';

import moment from "moment"


moment.locale('zh-cn')

const props=defineProps(['active'],['route'])
const containerRef = ref<HTMLElement | null>(null)
// 获取当前实例
const instance = getCurrentInstance();
const myGlobalVarHost = instance?.appContext.config.globalProperties.$globalHost;
const appglobalHost=instance?.appContext.config.globalProperties.$appglobalHost;

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
    "anchor":1,
    theme:{current:props.active?"dark":"light"},
    mode:"dark",
    speech:{"enable":true},
    hljs: {
          style: 'monokai',
        },
    after () {
        const outlineElement = document.getElementById('outline')
        const bgconter=document.getElementById("bgdv_right_conter")
        if (window.innerWidth <= 768) {
           outlineElement.style.display = 'none'
           bgconter.style.left=0
           bgconter.style.width='90%'
          return
        }else{
          bgconter.style.left='250px'
          bgconter.style.width='80%'
        }
        if(timeID){
         clearTimeout(timeID);
        }
        timeID = setTimeout(()=>{ 
           Vditor.outlineRender(document.getElementById('vditor'), outlineElement)
           if (outlineElement.innerText.trim() !== '') {
             outlineElement.style.display = 'block'
             initOutline()
           }
        },500)
      },
 }

 
 
 const initOutline = () => { 
    const headingElements = []
  
    Array.from(document.getElementById('vditor').children).forEach((item) => {
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
  const data = msgdata._value //await getContent()
  Vditor.preview(
        document.getElementById('vditor'),
        data,IPreviewOptions
    )  
 }

 const routes=useRoute()

  onMounted(async()=>{ 
   await fetchdata()
   if(msgdata._value==null){
    await getContent() 
  }
   
   await addEventListener("resize",initVditor)
   initVditor() 
   
 })
 const msgdata=ref(null) 
 const getContent = async() => {
  // 重新获取内容的逻辑，可以调用你的 API 获取最新内容
  try {
    const result = await axios.get(myGlobalVarHost+'api/v1/articles/' + routes.query.id, { //'http://127.0.0.1:2023/api/v1/articles/' + routes.query.id, {
      headers: {
        'token': sessionStorage.token,
      },
    })
    document.getElementById('hd_title').innerHTML=result.data.data.title
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

const encode=ref('') 
const fetchdata= async()=>{
    try {
        await axios.get(myGlobalVarHost+'Getencrypt/' + routes.query.id).then((res) => { 
        encode.value=res.data.data
        })
    }catch (error) {
        console.error(error); 
    }
}
const CopyPage=(event)=> {
   event.preventDefault(); // 阻止默认行为 
   try{ 
        var cptext=appglobalHost+"BlogRead?id="+encode._value;  //"http://127.0.0.1:2024/BlogRead?id="+encode
        const textarea = document.createElement('textarea')
        textarea.value =cptext
        document.body.appendChild(textarea)
        textarea.select()
        document.execCommand('copy')
        document.body.removeChild(textarea)
         ElMessage({message: '分享链接已拷贝至剪切板',type: 'success',}); 
      }catch{
         ElMessage({message: '当前版本不支持直接拷贝，可分享网址：'+cptext,type: 'warning',});
      } 
  }
 
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
.dv_hd_userinfo1{
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
.bgdv_right_conter{
  position: absolute;
  left: 250px;
  width: 80%;
  /* margin-top: -90px; */
  padding-top: 90px;
  margin-top: -90px;
}
</style>