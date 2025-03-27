<template>
    <el-form class="demo-rule-form" label-width="80px" style="margin-top: 16px;">
      <div id="vditor" class="vditor-reset" name="description"></div>
      <div id='outline'></div>

      <div style="width:200px;float:right;position: fixed;top: 160px;right: 0;">
     
     <el-anchor id="md_ancho"   :container="containerRef"
         direction="vertical"
         bound="-15"
         type="underline"
         :offset="30"
         >
       
         <el-anchor-link v-for="(o, index) in  statehh.message "   :key="o"   :href="`#${statehh.message[index]}`" :title="statehh.message[index]" ></el-anchor-link>
          
       </el-anchor>
   </div>
    </el-form>
     
  </template>
  <script setup>
  import Vditor from 'vditor'
  //import Vditor from 'vditor/dist/method.min'
  import 'vditor/dist/index.css'
  import axios from 'axios'
  import { progressProps } from 'element-plus';
  import { computed, onMounted,reactive} from 'vue';
  const props=defineProps(['active'])
  let active=computed({
      get(){
          return props.active;
      }
  })
  const statehh = reactive({
      message: [],
    });
 

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
        }
      }
    })
  }


   const IPreviewOptions={
      "anchor":1,
      theme:{current:props.active?"dark":"light"},
      mode:"dark",
      speech:{"enable":true},
      
      after () {
          if (window.innerWidth <= 768) {
            return
          }
          const outlineElement = document.getElementById('outline')
          Vditor.outlineRender(document.getElementById('vditor'), outlineElement)
          if (outlineElement.innerText.trim() !== '') {
            outlineElement.style.display = 'block'
            initOutline()
          }
        },
   }
  
  const datacon=reactive({
     datas:""
  })
   let test=""
   
  //  const  result=axios({
  //          method: "get",
  //          url: "http://127.0.0.1:2023/api/v1/articles/7",
  //          headers:{
  //             "token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QiLCJwYXNzd29yZCI6InRlc3QxMjM0NTYiLCJleHAiOjE2OTA3NDE0NjEsImlzcyI6Imdpbi1ibG9nIn0.v0NiAuZaCSou82WJF5dt4TfXiwczCGHDYt4HuGgcZW0"
  //          }
  //         })
  //         .then((res) => {
  //          //console.log(res);
  //          //console.log(res.data.data.content)
  //          //data =res.data.data.content
  //          //console.log(data)
  //          //datacon.datas= res.data.data.content
  //          test=res.data.data.content
            
  //          console.log(res.data.data.content)
  //         }) 
          
  
   //var mdStrs= result.then(res=>{return res})
   
   
    
  //   let mdStrs=await result.then(res=>{return res})
  //   const test=mdStrs
  //   console.log(test)
    //const mdStr=`## 标题1 
    //发法沙发沙发撒法
    //## 标题二 ` 

 
  const mdStr=
     `<h1>输入个标题~</h1>
    <pre><code>public string test()
    {
       return &quot;hello world&quot;;
    }
    </code></pre>
    <blockquote>
    <p>毛主席语录：</p>
    <p>好好学习，天天向上！</p>
    <p>--致同学们</p>
    </blockquote>
    <h2>输入个标题2~</h1>
    <h2>输入个标题3~</h1>

      `
  
   function ReInitVidor(data){ 
    //console.log(GetData()) 
     
      Vditor.preview(
          document.getElementById('vditor'),
          mdStr,IPreviewOptions
      )
      //Vditor.outlineRender(document.getElementId('vditor'), document.getElementId('vditor'))
   } 
  
    onMounted(async()=>{
     //const  result=axios({
     //      method: "get",
     //      url: "http://127.0.0.1:2023/api/v1/articles/7",
     //      headers:{
     //         "token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QiLCJwYXNzd29yZCI6InRlc3QxMjM0NTYiLCJleHAiOjE2OTA5MDUyNTYsImlzcyI6Imdpbi1ibG9nIn0.mJHQgLO6c5fF4lQkQ7zkpiOjh3l9KPI8s8wBuOYp9lg"
     //      }
     //     })
     //     .then((res) => {
     //      //test=res.data.data.content 
     //      console.log(res.data.data.content)
     //      return res.data.data.content 
     //     }) 
     //let mdStrs=await result.then(res=>{return res})
     //console.log("test"+mdStrs)
     addEventListener("resize",ReInitVidor)
  
     ReInitVidor(mdStr)
     
     setTimeout(() => {
         Array.from(document.getElementById('vditor').children).forEach((item) => {
           if (item.tagName.length === 2 && item.tagName !== 'HR' && item.tagName.indexOf('H') === 0) {
             statehh.message.push(item.id) 
           }
         })
      }, 1000);
   })
  
  
   
  </script>