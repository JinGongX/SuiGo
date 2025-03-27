<script>
import Vditor from 'vditor'
import 'vditor/dist/index.css'
import axios from 'axios'


export default {
  data() {
    return {
      contentEditor: {}, //3.声明一个变量
      ruleForms: {
        title: '',
        tags: '',
        content: ''
      },
      tagID:1,
      createdBy:"",
      contess:'',
    }
  },
  mounted() {
    let that = this;
    //this.getdata()
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
        linkToImgUrl:this.$globalHost+'upload',//'http://127.0.0.1:2023/upload',
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
        //const testddd=  this.getdata()
        //console.log("test:"+contess)
        //this.contentEditor.setValue(contess) 
        
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
         
         this.getdata()
         //this.contentEditor.setValue(this.contess,true)
      },
       toolbar: [
         'emoji',
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
         '|',
         'fullscreen',
         {
           name: 'more',
           toolbar: [
             'fullscreen',
             'both',
             'preview',
             'info',
             'help',
           ]
         },{
           name:'提交',
           tip:'提交',
           className:'mkbtn',
           icon:'<button aria-disabled="false" type="button" class="el-button el-button--primary"><span class="">编辑提交</span></button>',
           click(){
            that.submitForm('ruleForms')
            //alert('test')
           }

         }],
         fullscreen:{
          index:9999
         },tipPosition: 'se'
    })

   
     
    //this.getdata()
    //const currentElement = document.querySelector('.vditor-tooltipped')
    //currentElement.classList.remove('vditor-tooltipped__ne')
    //currentElement.classList.add('vditor-tooltipped__se')
         //document.getElementsByClassName('vditor-tooltipped').classList.remove('vditor-tooltipped__ne')
  },
  methods: {
    // 我的提交表单代码大致如下,这段代码核心是this.ruleForm.content = this.contentEditor.getValue()
    getdata(){ 
        console.log(this.$route.query.id)
        
        axios({
         method: "get",
         url: this.$globalHost+"api/v1/articles/"+this.$route.query.id, //"http://127.0.0.1:2023/api/v1/articles/"+this.$route.query.id,
         headers:{
            "token":sessionStorage.token
         }
        })
        .then((res) => {
         //test=res.data.data.content 
         //console.log(res.data.data.content) 
         this.contess=res.data.data.content
         //console.log("dddddddddddddddddddddddddddd:"+this.contess)
         this.contentEditor.setValue(this.contess,true)
         //this.textconts=res.data.data.content
         //console.log(sessionStorage.token)
          
        }) 
        
    },
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

          //5.通过this.contentEditor.getValue()获取编辑器内容
          this.ruleForms.content = this.contentEditor.getValue()
          console.log(this.ruleForms.content)
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
          var tid=this.$route.query.id
          const intValue = parseInt(this.$route.query.id);
          console.log(this.setText(this.contentEditor.getHTML()))
          axios({
            method: "put",
            url: this.$globalHost+"api/v1/articles/"+this.$route.query.id, //"http://127.0.0.1:2023/api/v1/articles/"+this.$route.query.id,
            headers:{
              "token":sessionStorage.token
            },
            data: {
              //TagID:this.tagID,
              ID:intValue,
              Content:this.contentEditor.getValue(),//getHTML(),
              Title:document.getElementById("test").firstChild.textContent,
              Desc:this.setText(this.contentEditor.getHTML()),
              modified_by:this.createdBy,
              Imgcover:imgsrc!=""?imgsrc:"" //document.getElementById("test").querySelector("img").src
            }
          }).then((res) => {
            console.log(res);
            alert("提交完成！")
           location.href= "/EditBlogs?id="+tid
          })
          //axios({
          //  method: "PUT",
          //  url: "http://127.0.0.1:2023/api/v1/articles/"+this.$route.query.id,
          //  headers:{
          //    "token":sessionStorage.token
          //  },
          //  data: {
          //    //TagID:this.tagID,
          //    ID:this.$route.query.id,
          //    Content: this.contentEditor.getValue(),//getHTML(),
          //    Title:document.getElementById("test").firstChild.textContent,
          //    Desc:this.setText(this.contentEditor.getHTML()),
          //    ModifiedBy:this.createdBy,
          //    Imgcover:imgsrc!=""?imgsrc:"" //document.getElementById("test").querySelector("img").src
          //  }
          //}).then((res) => {
          //  console.log(res);
          //  alert("提交完成！")
          // location.href= "/EditBlogs?id="+this.$route.query.id
          //})


        }
      })
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

  }
   
  

}
</script>

<template>
  <el-form ref="ruleForms" :model="ruleForms" class="edemo-rule-form" label-width="80px" style="margin-top: 16px;">
    <div id="vditor" name="description" class="vditor  vditor--fullscreen"></div>
    <el-button type="primary" @click="submitForm('ruleForms')">编辑提交</el-button>

  </el-form>
  <div class="test" id="test">

  </div>
</template>
<style>
.mkbtn{
  float:right;
  margin-right: 80px;
  margin-top: -8px;
}
</style>