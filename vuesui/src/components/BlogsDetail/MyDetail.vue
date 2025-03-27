<template> 
  <el-breadcrumb separator="/">
    <el-breadcrumb-item :to="{ path: '/myBlogs' }">随记</el-breadcrumb-item>
    <!-- <el-breadcrumb-item
      ><a href="/">promotion management</a></el-breadcrumb-item
    > -->
    <el-breadcrumb-item><span id="hd_title"></span></el-breadcrumb-item>
  </el-breadcrumb>
  
  <el-form class="demo-rule-form" id="test" label-width="80px" style="margin-top: 16px;">
    <div id='outline' style="float:left;display: block;position: fixed;height: 85%;max-height: 85%;border-right:1px solid #e0e0e0"></div>
   
    <div id="bgdv_right_conter" class="bgrddv_right_conter"> 
    <div class="dv_hd_userinfo" > 
      <div class="dv_op_photo"><img id="op_photo" alt="photo" class="photo" src="../../assets/photo.jpg" /> </div>
      <div class="dv_username">作者: <span id='username'></span></div>
      <div class="dv_hd_title" >创建时间 :  <span id="hd_crdate"></span></div>

      <el-button type="warning" title="分享"  @click="CopyPage" style="margin-top: 4px;"><el-icon><Share /></el-icon></el-button> 
      <el-button type="warning" title="导出&打印"  @click="printPage" style="margin-top: 4px;"><el-icon><Printer /></el-icon></el-button> 
      
    </div>
    
   
    <div id="vditorsview" class="vditor-reset" name="description" ></div>
    <div class="dv_foot"></div>
  </div>
  </el-form> 
  <div class="content-to-print">
      <div id="export_pdf" class="vditor-reset" name="description"></div>
  </div>
</template>
<script setup>

import Vditor from 'vditor'
//import Vditor from 'vditor/dist/method.min'
import 'vditor/dist/index.css'
import axios from 'axios'
import { ElMessage } from 'element-plus';
import{useRoute} from 'vue-router'

import { getCurrentInstance ,computed, onMounted,reactive,ref} from 'vue';

import moment from "moment" 

moment.locale('zh-cn')

// 获取当前实例
const instance = getCurrentInstance();
const myGlobalVarHost = instance?.appContext.config.globalProperties.$globalHost;
const appglobalHost=instance?.appContext.config.globalProperties.$appglobalHost;
const props=defineProps(['active'],['route'])
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
 console.log(id)
 const IPreviewOptions={
    //"anchor":1,
    theme:{current:props.active?"dark":"light"},
    mode:"dark",
    speech:{"enable":true},
    
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
       
        Vditor.outlineRender(document.getElementById('vditorsview'), outlineElement)
        
        if (outlineElement.innerText.trim() !== '') {
          outlineElement.style.display = 'block'
          initOutline()
        }
      },
 }
 

 const initOutline = () => {
    const headingElements = []
    Array.from(document.getElementById('vditorsview').children).forEach((item) => {
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
        document.getElementById('vditorsview'),
        data,IPreviewOptions
    ) 
    Vditor.preview(
        document.getElementById('export_pdf'),
        data,IPreviewOptions
    ) 
 }

 const routes=useRoute()
  onMounted(async()=>{
   await fetchdata()
   if(msgdata._value==null){
    await getContent()
   }
   
   //const routes=useRoute()
   //console.log("id:"+routes.query.id)
   
   //const  result= axios({
   //      method: "get",
   //      url: myGlobalVarHost+"api/v1/articles/"+routes.query.id,//"http://127.0.0.1:2023/api/v1/articles/"+routes.query.id,
   //      headers:{
   //         "token":sessionStorage.token
   //      }
   //     })
   //     .then((res) => {
   //      //test=res.data.data.content 
   //      ////console.log(res.data.data.content)
   //      //console.log("tests:"+res.data.data.account.name)
   //      document.getElementById('hd_title').innerHTML=res.data.data.title
   //      document.getElementById('username').innerHTML=res.data.data.account.name
   //      document.getElementById('op_photo').src=res.data.data.account.photo
   //      document.getElementById('hd_crdate').innerHTML=moment(res.data.data.created_on*1000).format("YYYY-MM-DD HH:mm") 
   //      
   //      return res.data.data.content //res.data.data.content 
   //     }) 
   
  await addEventListener("resize",initVditor)
 initVditor()
 })

  

 const printPage=(event)=> {
   event.preventDefault(); // 阻止默认行为
   window.print();
  }
 
  const msgdata=ref(null) 
  const getContent = async () => {
  // 重新获取内容的逻辑，可以调用你的 API 获取最新内容
  try {
    const result = await axios.get(myGlobalVarHost+'api/v1/articles/' + routes.query.id, { //'http://127.0.0.1:2023/api/v1/articles/' + routes.query.id, {
      headers: {
        'token': sessionStorage.token,
      },
    });
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
@media print { 
  body * {
    visibility: hidden;
  }  
  .content-to-print,.content-to-print * {
    visibility: visible;
    top: 10;
    width:94%;
    /*width: 100%;  设置内容宽度为100% */
    /*page-break-before: avoid;  防止额外分页 */  
  }
  #vditorsview{
    width:0px;
    height:0px;
  }
}
.content-to-print{
  visibility: hidden;   
  float:left;
  width:94%;
  max-width: 786px;
  margin: 20px 10%;
  padding: 0;
  position: absolute;
  top:10px;
} 
#vditorsview{
  /*float:left;width:77%;padding-left: 300px;padding-bottom: 20px;*/
  width:100%;max-width: 768px;margin:0 auto;
}
.vditor-outline__item--current{
  /* border-left:2px solid rgb(106, 146, 226) */
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
.dv_hd_userinfo{
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
/* .vditor-reset h1, .vditor-reset h2, .vditor-reset h3, .vditor-reset h4, .vditor-reset h5, .vditor-reset h6 {
  padding-top: 60px;
  margin-top: -60px;
} */

::-webkit-scrollbar {
    width: 6px;
    height: 6px;
    background-color: #f5f5f5;
}

/*滚动条 阴影~圆角*/
::-webkit-scrollbar-track {
    -webkit-box-shadow: inset 0 0 6px rgba(128, 128, 128, 0.7);
    border-radius: 10px;
    background-color: #f5f5f5;
}

/*滑块 阴影~圆角*/
::-webkit-scrollbar-thumb {
    border-radius: 10px;
    -webkit-box-shadow: inset 0 0 6px rgba(128, 128, 128, 0.7);
    background-color: rgb(149, 147, 147);
}
.dv_foot{
  width: 80%;
  height:120px;
}
.bgrddv_right_conter{
  position: absolute;
  left: 250px;
  width: 80%;
  /* margin-top: -90px; */
  padding-top: 90px;
  margin-top: -90px;
}
</style>