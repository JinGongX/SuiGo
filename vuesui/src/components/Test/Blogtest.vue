<template>
    <div @mouseup="handleMouseUp">
    <div v-if="showMenu" :style="menuStyle"  class="custom-menu">
      <button class="btn" @click="underlineSelection">划线</button>
      <!-- <div @click="underlineSelection" >
        划线
      </div>  -->
    </div>
        <MdPreview   style="margin-top:-20px;"  :autoFoldThreshold="9999" :editorId="id" :modelValue="messageb" />
    </div>
     

  </template>
  <script setup>
  //import { ref } from 'vue';
  import { MdPreview } from 'md-editor-v3';
  import 'md-editor-v3/lib/preview.css';
  
  const id = 'preview-only';
  //const text = ref('# Hello Editor');
  </script>
  <script>  
  import axios from 'axios' 

  export default {
    data() {
      return {
        form: {
          desc: ''
        },
        messageb:"",
        username:'Jin',
        loadingtype:false, 
        controller: null, 
        arequestData : {
            model: "qwen2:0.5b",//"llama3.1",
            messages: []
        },
        showMenu:false,
        selectedText:'',
        menuStyle:{
            top: '0px',
            left: '0px'
        },
      }
    },
    mounted() {  
        axios({
         method: "get",
         url: "http://127.0.0.1:2023/Readarticles/aa90f92c89b6153f32cf62d776129e780784", //"http://127.0.0.1:2023/api/v1/softinfo",
         
        })
        .then((res) => {
          //console.log(res)
          console.log(res.data.data.content)
          
          this.messageb=res.data.data.content
          //this.totals = response.body.count // this.items.length
          //this.page = response.body.totalpages
          //console.log(response.body.TotalPages)
        }) 
         
        //try {
        //  const result =  axios.get('http://127.0.0.1:2023/Readarticles/aa90f92c89b6153f32cf62d776129e780784', { //'http://127.0.0.1:2023/api/v1/articles/' + routes.query.id, {
        //   
        //  })
        //  console.log(result.data)
        //  this.messageb=result.data.data.content;
        //  console.log(this.messageb)
        //  //return result.data.data.content;
        //} catch (error) {
        //  
        //  console.error(error);
        //  //return ''; // 处理获取内容失败的情况，确保返回一个默认值
        //}
        setTimeout(()=>{
          this.$nextTick(()=>{
            
            //Vditor.preview(document.getElementById('vditorsuichat'),"fsafsafsdf");
              //let scroll=this.$refs["scrollDiv"].wrap;
              //scroll.scrollTop = scroll.scrollHeight+30
          })
        },100)
    },
    methods: {
       handleMouseUp(event){
            const selection = window.getSelection()
            const text = selection.toString()
            if (text) {
              this.selectedText = text
              const range = selection.getRangeAt(0)
              const rect = range.getBoundingClientRect()
              this.menuStyle = {
                top: `${rect.bottom + window.scrollY+10}px`,
                left: `${rect.left + window.scrollX}px`
              };
              this.showMenu = true
            } else {
                this.showMenu = false
            }
          },
        underlineSelection(){
        const selection = window.getSelection();
        if(!selection.rangeCount) return;
        const range = selection.getRangeAt(0)
        const selectedContents = range.cloneContents();
        const allUnderlinedElements=document.querySelectorAll("span[style*='border-bottom: 2px solid rgb(160, 207, 29)']")
        let affectedElements=[];
        allUnderlinedElements.forEach(underlineElement=>{
          if(range.intersectsNode(underlineElement)){
            affectedElements.push(underlineElement);
          }
        });
        const containsNonUnderlined=this.containsNonUnderlinedText(selectedContents);
        affectedElements.forEach(underlineElement=>{
          this.removeEntireUnderline(underlineElement);
        });
        if(containsNonUnderlined){
          const newContents=document.createDocumentFragment();
          newContents.appendChild(range.extractContents());
          const span=document.createElement('span');
          span.style.borderBottom='2px solid #a0cf1d';
          span.appendChild(newContents);
          range.deleteContents();
          range.insertNode(span);
        }
      },
      containsNonUnderlinedText(selectedContents){
        return Array.from(selectedContents.childNodes).some(node=>{
          return node.nodeType===Node.TEXT_NODE&&node.textContent.trim()!==""||(node.nodeType===Node.ELEMENT_NODE&&node.style.borderBottom!=='border-bottom: 2px solid rgb(160, 207, 29)');
        });
      },
      removeEntireUnderline(underlineElement){
         const parent=underlineElement.parentNode;
         while(underlineElement.firstChild){
           parent.insertBefore(underlineElement.firstChild,underlineElement);
         }
         parent.removeChild(underlineElement);
      }
    }
  }
  </script>
  <style scoped>
  .btn{
    background: #fff;
    border:0px;
    cursor: pointer;
  }
  .custom-menu {
  position: absolute;
  background-color: #fff;
  border: 1px solid #ccc;
  z-index: 1000;

  padding: 5px;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.2);
}
  </style>