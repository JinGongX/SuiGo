<script setup>
import { MdPreview, MdCatalog } from 'md-editor-v3';
// preview.css相比style.css少了编辑器那部分样式
import 'md-editor-v3/lib/preview.css';

const id = 'adpreview-only'; 
</script>
<script>
import Vditor from 'vditor'
import 'vditor/dist/index.css'
import axios from 'axios'
import { useTransitionFallthroughEmits } from 'element-plus'
 

export default { 
  data() {
    return {
      contentEditor: {}, //3.声明一个变量
      ruleForm: {
        title: '',
        tags: '',
        content: ''
      },
      tagID:1,
      createdBy:"",
      dialogVisible:false,

      form: {
        desc: ''
      },
      message:[],
      username:'Jin',
      loadingtype:false,
      controller: null,
      arequestData : {
          model: "qwen2:0.5b",//"llama3.1",
          messages: []
      }
    }
  },
  mounted() {
    let that = this;
    this.contentEditor = new Vditor('vditor', { //4.刚刚声明的变量contentEditor被赋值为一个Vditor实例,
      height: 500,
      placeholder: '# 输个标题吧……',
      theme: 'classic',
      //counter: {
      //  enable: true,
      //  type: 'markdown'
      //},
      outline: {
       enable: true,
       position: "left"
      },
      preview: {
        delay: 0,
        hljs: {
          style: 'monokai',
          lineNumber: true
        },
        anchor:1,
      },
      
      tab: '\t',
      typewriterMode: true,
      toolbarConfig: {
        pin: true
      },
      cache: {
        enable: false
      },
      mode: 'sv',//'sv',
      upload: {
        accept: 'image/*,.mp3, .wav, .rar',
        token: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QiLCJwYXNzd29yZCI6InRlc3QxMjM0NTYiLCJleHAiOjE2ODg2Njg0NjYsImlzcyI6Imdpbi1ibG9nIn0.K4CLvzR51WwTMM2QfzMMiyKFxlZtX9mwv5Oh7eqcsxs',
        multiple:false,
        url: 'api/upload',
        linkToImgUrl:this.$globalHost+'upload', //'http://127.0.0.1:2023/upload',
        fieldName:'file',
        filename(name) {
          return name.replace(/[^(a-zA-Z0-9\u4e00-\u9fa5\.)]/g, '').
            replace(/[\?\\/:|<>\*\[\]\(\)\$%\{\}@~]/g, '').
            replace('/\\s/g', '')
        },
        format(files,responseText){ 
          const res=JSON.parse(responseText); 
          const name=files[0].name; 
          const url="F://vscodepj2023/gotest20230626/files/"+res.name;
          console.log(window.location.origin+'/src/files/'+res.name)
          const result=JSON.stringify({
              msg:'',
              code :0,
              data:{errFiles:[],succMap:{[res.name]:window.location.origin+'/src/files/'+res.name,}
            },
          }); 
          return result;
        }
      },
      after:()=>{
          this.contentEditor.setValue('# 输入个标题~')

          const currentElement = document.querySelectorAll('.vditor-tooltipped')
          // 遍历这些元素并修改它们的内容
          currentElement.forEach(element => {
           if (element.className == 'vditor-tooltipped vditor-tooltipped__ne') {
             element.classList.remove('vditor-tooltipped__ne')
             element.classList.add('vditor-tooltipped__se')
           }
           if (element.className == 'vditor-tooltipped vditor-tooltipped__n') {
             element.classList.remove('vditor-tooltipped__n')
             element.classList.add('vditor-tooltipped__s')
           }
           if (element.className == 'vditor-tooltipped vditor-tooltipped__nw') {
             element.classList.remove('vditor-tooltipped__nw')
             element.classList.add('vditor-tooltipped__sw')
           }
          });
      },
       toolbar: [
         //'emoji',
         'headings',
         'bold',
         'italic',
         'strike',
         'link',
         '|',
         'list',
         'ordered-list',
         'check',
         'outdent',
         'indent',
         '|',
         'quote',
         'line',
         'code',
         'inline-code',
         'insert-before',
         'insert-after',
         '|',
         // 'record',
         'table',
         '|',
         'undo',
         'redo',
         '|',
         'upload',
         'edit-mode',
         // 'content-theme',
         'code-theme',
         'export',
         
         {
           name: 'more',
           toolbar: [
             'fullscreen',
             'both',
             'preview',
             'info',
             'help',
           ],
         },{
           name:'提交',
           tip:'提交',
           className:'mkbtn',
           icon:'<button aria-disabled="false" type="button" class="el-button el-button--primary"><span class="">提交</span></button>',
           click(){
            that.submitForm('ruleForm')
            //alert('test')
           }

         },{
           name:'谁可见',
           tip:"谁可见",
           className:'mklabel',
           icon:'<select id="sel_stare" class="mkselect"><option selected value="1">公开</option><option value="2">仅我可见</option></select>'
         },{
           name:'标签',
           tip:"标签",
           className:'mklabel_b',
           icon:'<select id="sel_type" class="mkselect" ><option selected value="1">默认</option><option value="2">计算机</option><option value="3">工程</option><option value="4">经济</option></select>'
         },{
           name:'AI随推',
           tip:"AI随推",
           className:'mkgpt',
           icon:'<svg t="1715692443706" class="mkgpt" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="4579" width="128" height="128"><path d="M966.392 448.741L851.745 249.689a27.039 27.039 0 0 0-23.431-13.544h-42.126l-78.386-135.769a27.04 27.04 0 0 0-23.417-13.52H454.539a27.039 27.039 0 0 0-23.4 13.49l-20.99 36.249-156.729 0.025a27.04 27.04 0 0 0-23.413 13.52L115.084 349.191a27.041 27.041 0 0 0 0 27.04l20.925 36.243-78.386 135.769a27.041 27.041 0 0 0 0 27.04l114.923 199.052a27.04 27.04 0 0 0 23.417 13.52h41.85l78.386 135.769a27.04 27.04 0 0 0 23.417 13.52h229.845a27.04 27.04 0 0 0 23.417-13.52l20.925-36.243h156.773a27.04 27.04 0 0 0 23.417-13.52l114.923-199.052a27.041 27.041 0 0 0 0-27.04l-20.925-36.243 78.386-135.769a27.04 27.04 0 0 0 0.015-27.016zM767.96 557.282l-219.424-0.026 82.291-142.529a4.517 4.517 0 0 0 0-4.507L462.345 118.403h219.439l84.889 147.035 0.003 0.004 84.888 147.032-83.604 144.808z m62.68-90.538h104.516l-78.388 135.767L804.51 512l26.13-45.256z m-208.921-54.27L540.73 552.748 517.204 512l27.429-47.509c0.031-0.054 0.042-0.115 0.07-0.17 0.125-0.24 0.228-0.49 0.31-0.753 0.028-0.091 0.062-0.179 0.084-0.271 0.082-0.341 0.139-0.693 0.139-1.059a4.46 4.46 0 0 0-0.139-1.059c-0.022-0.092-0.056-0.181-0.084-0.272a4.565 4.565 0 0 0-0.31-0.753c-0.029-0.055-0.039-0.116-0.07-0.17L429.711 260.932a4.509 4.509 0 0 0-3.904-2.254h-49.653l78.384-135.768 167.181 289.564zM365.744 258.678h-52.255l-52.259-90.514 156.787-0.024-52.273 90.538z m-62.664 0h-49.655a4.508 4.508 0 0 0-3.904 2.254l-82.29 142.528-23.526-40.748 109.719-190.038 49.656 86.004z m-131.946 156.05l84.891-147.036h54.856l0.009 0.001 0.009-0.001h112.307l109.718 190.039H368.347a4.508 4.508 0 0 0-3.904 2.254L195.961 751.802 86.244 561.763l84.89-147.035z m199.814 52.016h161.975l-23.527 40.749h-54.859c-0.037 0-0.073 0.014-0.11 0.015a4.484 4.484 0 0 0-2.044 0.548c-0.08 0.044-0.154 0.093-0.231 0.142a4.454 4.454 0 0 0-0.652 0.499c-0.068 0.063-0.139 0.122-0.203 0.189a4.439 4.439 0 0 0-0.662 0.86L335.712 708.799a4.515 4.515 0 0 0 0 4.506l24.828 43.003H203.769l167.179-289.564z m-5.205 298.578l26.129 45.257-52.258 90.511-78.384-135.768h104.513z m31.334 54.272l24.826 43.001a4.508 4.508 0 0 0 3.904 2.253h164.579l-23.524 40.749h-219.44l49.655-86.003z m31.331 36.241L344.82 711.052l109.719-190.038 82.288 142.528a4.508 4.508 0 0 0 3.904 2.253h336.963l-109.718 190.04H428.408z m114.923-199.053l-80.987-140.276h47.055l27.427 47.51c0.044 0.076 0.103 0.137 0.152 0.21 0.087 0.134 0.177 0.264 0.279 0.389 0.095 0.115 0.197 0.221 0.301 0.326 0.1 0.1 0.2 0.197 0.311 0.288 0.13 0.107 0.266 0.201 0.406 0.292 0.069 0.046 0.128 0.102 0.2 0.144 0.034 0.02 0.073 0.025 0.107 0.044 0.314 0.17 0.64 0.315 0.984 0.406 0.035 0.01 0.071 0.009 0.107 0.018 0.302 0.072 0.61 0.11 0.92 0.119 0.044 0.001 0.088 0.017 0.132 0.017 0.051 0 0.099-0.018 0.15-0.019 0.062 0.002 0.119 0.018 0.18 0.018l229.505 0.027a4.509 4.509 0 0 0 3.904-2.253l24.841-43.027 53.556 92.759 0.001 0.003 24.83 43.005H543.331z m292.512-199.051l24.828-43.003a4.517 4.517 0 0 0 0-4.507l-82.289-142.529h47.325l109.457 190.039h-99.321z" p-id="4580"></path></svg>',
           click(){
            //that.submitForm('ruleForm')
            that.dialogVisible=true;
            //alert('test')
           }
          }],
    })

    setTimeout(()=>{
        this.$nextTick(()=>{
            let scroll=this.$refs["scrollDiv"].wrap;
            scroll.scrollTop = scroll.scrollHeight+30
        })
      },100)
  },
  methods: {
    // 我的提交表单代码大致如下,这段代码核心是this.ruleForm.content = this.contentEditor.getValue()
    
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          if (
            this.contentEditor.getValue().length === 1 ||
            this.contentEditor.getValue() == null ||
            this.contentEditor.getValue() === ''
          ) {
            alert('话题内容不可为空')
            return false
          }
          
          var myselect=document.getElementById("sel_stare")
          var index=myselect.selectedIndex
          var sel_tate=parseInt(myselect.options[index].value)

          //if(myselect.options[index].value!=""){
          //  var t=parseInt(myselect.options[index].value)
          //  alert(parseInt(myselect.options[index].value))
          //  return false
          //}
          var myselect_type=document.getElementById("sel_type")
          var index=myselect_type.selectedIndex
          var sel_types=parseInt(myselect_type.options[index].value)
          //if(myselect_type.options[index].value!=""){
          //  alert(myselect_type.options[index].value)
          //  return false
          //}
          
          //5.通过this.contentEditor.getValue()获取编辑器内容
          this.ruleForm.content = this.contentEditor.getValue()
          console.log(this.ruleForm.content)
          console.log(this.contentEditor.getHTML())
          document.getElementById("test").innerHTML=this.contentEditor.getHTML()
          //console.log(document.getElementById("test").firstChild.textContent)
          //console.log(document.getElementById("test").querySelector("h1").innerText)
          //
          //console.log(document.getElementById("test").querySelector("img").src)
          var imgsrc=""
          if(document.getElementById("test").querySelector("img")!=null)
              imgsrc=document.getElementById("test").querySelector("img").src
          this.createdBy=sessionStorage.id //"ming_test"
          console.log(this.setText(this.contentEditor.getHTML()))
          axios({
            method: "post",
            url: this.$globalHost+"api/v1/articles", //"http://127.0.0.1:2023/api/v1/articles",
            headers:{
              "token":sessionStorage.token
            },
            data: {
              tag_id:sel_types,//this.tagID,
              Content: this.contentEditor.getValue(),//getHTML(),
              Title:document.getElementById("test").firstChild.textContent,
              Desc:this.setText(this.contentEditor.getHTML()),
              created_by:this.createdBy,
              Imgcover:imgsrc!=""?imgsrc:"", //document.getElementById("test").querySelector("img").src
              state:sel_tate
            }
          }).then((res) => {
            console.log(res);
            alert("提交完成！")
           location.href= "/Mkdown"
          })


        }
      })
    },
    async startStreaming() {
      // 如果已经有一个正在进行的流式请求，则中止它
      if (this.controller) {
        this.controller.abort();
      }

      setTimeout(()=>{
        this.scrollToBottom();
       },50);
        var mymsg=this.form.desc.trim(); 
        if(mymsg.length>0){
               this.form.desc='';
               this.message.push({
                   user:'Jin',
                   msg:mymsg
               })           
               this.message.push({
                       user:'GPT', //'assistant',//'GPT',
                       msg:'',
                       dot:''
               });
        
       
             // 创建一个新的 AbortController 实例
             this.controller = new AbortController();
             const signal = this.controller.signal;
             this.arequestData.messages.push({role:"user",content:mymsg});
              
             console.log(JSON.stringify(this.arequestData));
             try {
               const response = await fetch('http://127.0.0.1:11434/api/chat', {//'http://127.0.0.1:11434/api/chat', 
                 method: 'POST',
                 headers: {
                   'Content-Type': 'application/json'
                 },
                 body:JSON.stringify(this.arequestData),
                 signal
               });
       
               if (!response.body) {
                 throw new Error('ReadableStream not yet supported in this browser.');
               }
       
               const reader = response.body.getReader();
               const decoder = new TextDecoder();
               let result = '';
               //this.dot='●'
               this.message[this.message.length-1].dot='⚪';
               while (true) {
                 const { done, value } = await reader.read();
                 if (done) {
                   break;
                 }
                 result += decoder.decode(value, { stream: true });
        
                 // 处理流中的每一块数据，这里假设每块数据都是完整的 JSON 对象
                 const jsonChunks = result.split('\n').filter(line => line.trim());
                 //console.log(result)
                 for (const chunk of jsonChunks) {
                   try {
                     const data = JSON.parse(chunk);
                     //console.log(data.message.content) 
                     this.message[this.message.length-1].msg+=data.message.content;
                     setTimeout(()=>{
                       this.scrollToBottom();
                      },50); 
                   } catch (e) {
                     // 处理 JSON 解析错误
                     //console.error('Failed to parse JSON:', e);
                   }
                 } 
                 // 清空 result 以便处理下一块数据
                 result = '';
               }
             } catch (error) {
               if (error.name === 'AbortError') {
                 console.log('Stream aborted');
               } else {
                 console.error('Streaming error:', error);
               }
             }
             //this.dot='';
             this.message[this.message.length-1].dot='';
             this.arequestData.messages.push({
                     role: 'assistant',//this.message[this.message.length-1].user,//"GPT",
                     content: this.message[this.message.length-1].msg
                   })
             setTimeout(()=>{
                this.scrollToBottom();
               },50); 
      }else{
          this.form.desc='';
        }
    },

    //截取前100字符做简述
    setText(val){ 
      if(val!=null && val!=""){
        var rel=new RegExp("<.+?>|&.+?;","g")
        var msg=val.replace(rel,"")
        msg=msg.replace(/\s/g,"")
        msg=msg.replace(/[\r\n]/g,"")
        return msg.substr(0,100);
      }else return ''
    },
    scrollToBottom() {
        let elscroll=this.$refs["scrollDiv"];
        elscroll.scrollTop = elscroll.scrollHeight+30
       
    },
    submitFormgpt(formName) {
        //this.$refs[formName].validate((valid) => {
        //    if (valid) {
        //        this.form.desc
        //    }
        //})
        //alert(this.form.desc)
        // var div=document.getElementById("messgebox");
	    // div.scrollTop = div.scrollHeight;

       setTimeout(()=>{
        this.scrollToBottom();
       },50);
        var mymsg=this.form.desc.trim(); 
        if(mymsg.length>0){
        this.form.desc='';
        this.message.push({
            user:'Jin',
            msg:mymsg
        })           
        this.message.push({
                user:'GPT',
                msg:''
        });
       
        const requestData = {
          model: 'moonshot-v1-32k',
          messages: [
            {
              role: 'user',
              content: mymsg
            }
          ],
          temperature: 0.7
        };


        axios({
            method: "post",
            url: this.$globalHost+"Chatkimi",//"http://127.0.0.1:2023/Chatkimi",
            data: requestData
          }).then((res) => {
            //console.log(res.data.choices[0].message.content);
            //this.message.push(res.data.data);
            this.message[this.message.length-1]={
                user:'GPT',
                msg:res.data.choices[0].message.content
            };

            setTimeout(()=>{
              this.scrollToBottom();
             },50);
          })
        }else{
          this.form.desc='';
        }
    },
    clearForm(formName){
      this.form.desc='';
    },cp_put(desc){
      this.contentEditor.setValue(desc)
    },cp_append(desc){
      this.contentEditor.setValue(this.contentEditor.getValue()+desc)
    }
  }
   
  

}
</script>

<template>
  <el-form ref="ruleForm" :model="ruleForm" class="demo-rule-form" label-width="80px" style="margin-top: 16px;">
    <div id="vditor" name="description" class="vditor vditor--fullscreen"></div>
    <!--<el-button type="primary" @click="submitForm('ruleForm')">提交</el-button>-->

  </el-form>
  <div class="test" id="test">

  </div>
  <el-drawer v-model="dialogVisible" :show-close="false" size="50%">
    <template #header="{ close, titleId, titleClass }">
      <h5 :id="titleId" :class="titleClass"><img alt="AI" class="ailoge" src="../assets/aislogo.png"/> AI随推</h5>
      <el-button type="danger" @click="close">
        <el-icon class="el-icon--left"><CircleCloseFilled /></el-icon>
        Close
      </el-button>
    </template>
    <el-row :gutter="12" class="demo-radius">
    <div
        class="radius"
        :style="{
          borderRadius: 'base'
        }">
        <div class="messge" id="messgebox"  ref="scrollDiv"> 
            <ul style="padding: 10px;">
  <li v-for="(item, index) in message" :key="index" style="list-style-type:none;">
    <div v-if="item.user == username" class="mymsginfo" style="float:right">
        <div><el-avatar style="float: right;margin-right: 30px;">  {{ item.user }}  </el-avatar>  
    </div><div style="float: right;margin-right: 10px;margin-top:10px;width:80%;text-align: right;"> {{ item.msg }} </div> 
     </div>
     <div v-else class="chatmsginfo" >
        <div>
        <el-avatar style="float: left;margin-right: 10px;">  {{ item.user }}  </el-avatar> 
    </div>
      <div style="float: left;margin-top:10px;width:86%;">
        <img alt="loading" v-if="item.msg == ''" class="loading" src="../assets/loading.gif"/> 
        <MdPreview style="margin-top:-20px;"  :autoFoldThreshold="9999" :editorId="id" :modelValue=" item.msg + item.dot  " /> 
        </div> 
        <div v-if="item.msg.length>0" style="float: right;margin-top: 10px;margin-right: 8px;">
          <div class="gpt_cp" @click="cp_put(item.msg)" title="嵌入随记"><img alt="嵌入随记"  class="cp_option" src="../assets/cp_start.png"/> </div>
          <div  class="gpt_cp" @click="cp_append(item.msg)" title="追加内容"><img alt="追加内容"  class="cp_option" src="../assets/cp_append.png"/> </div> 
        </div>
      </div>
      
  </li>
</ul>
        </div>
        <div class="inputmsg">
            <el-form :model="form" label-width="120px">
                <el-form-item label="Jin:">
                    <el-input v-model="form.desc"  placeholder="说说你想问点啥...." @keyup.enter.native="startStreaming" type="textarea" />
                </el-form-item>
                <el-form-item>
                  <el-button type="primary" @click="startStreaming">发送</el-button>
                  <el-button @click="clearForm('form')">清空</el-button>
                </el-form-item>
            </el-form>
        </div>
      </div>    
  </el-row>
  </el-drawer>
</template>
<style scoped>
.radius{
  margin:0 auto;
}
.demo-radius .title {
  color: var(--el-text-color-regular);
  font-size: 18px;
  margin: 10px 0;
}
.demo-radius .value {
  color: var(--el-text-color-primary);
  font-size: 16px;
  margin: 10px 0;
}
.demo-radius{
 height: -webkit-fill-available;
}
.demo-radius .radius {
  min-height: 580px;
  height: 85vh;
  width: 94%;
  border: 1px solid var(--el-border-color);
  border-radius: 14px;
  /* margin-top: 20px; */
}
.messge{
    width:96%;
    height:75%;
    /* border:1px solid red; */
    margin: 6px auto;
    overflow: hidden;
    overflow-y: auto;
}
.inputmsg{
    width:96%;
    height:21%;
    /* border:1px solid blue; */
    border-top:2px solid #ccc;
    margin: 4px auto;
    padding-top: 10px;
}
.mymsginfo{
    width:100%;
    height:auto;
    min-height:50px;
}
.cp_option{
  width:24px;
  height:24px;
}
.gpt_cp{
  margin-bottom: 6px;
  cursor: pointer;
}

::-webkit-scrollbar {width: 6px;height: 5px;
}
::-webkit-scrollbar-track {background-color: rgba(0, 0, 0, 0.2);border-radius: 10px;
}
::-webkit-scrollbar-thumb {background-color: rgba(0, 0, 0, 0.5);border-radius: 10px;
}
::-webkit-scrollbar-button {background-color: #7c2929;height: 0;width: 0px;
}
::-webkit-scrollbar-corner {background-color: black;
} 
</style>
<style>
.mkbtn{
  float:right;
  margin-right: 50px;
  margin-top: -8px;
}
.mklabel{
  float:right;
  margin-right: 64px;
  margin-top: -8px;
}
.mklabel_b{
  float:right;
  margin-right: 44px;
  margin-top: -8px;
}
.mkselect{
  height: 32px;
    /* border: 1px solid #4285f4;
    background-color: #4285f4; 
    color: #fff;*/
    border-radius: 5px;
}
.mkgpt{
  height: 32px;
    /* border: 1px solid #4285f4;
    background-color: #4285f4; 
    color: #fff;*/
    border-radius: 5px;
}
.ailoge{
  width:24px;
  height:24px;
  vertical-align: bottom;
}
</style>