<template>
    <el-form class="demo-rule-form" label-width="80px" style="width: 88%;margin: 16px auto 0 auto;">
      <div><img alt="xixi" style="width: 100%;" src="../../assets/other/1.gif"/> </div>
      <div v-if="contenitem==false" style="text-align: center"><img alt="loading"  class="loading" src="../../assets/loading.gif"/></div>
      <div id="vditor" class="vditor-reset" name="description"></div>
      <div id='outline'></div>
    </el-form>
    <!-- <audio autoplay loop>
      <source src="../../assets/other/青花瓷.mp3" type="audio/mpeg">
    </audio> -->
    <audio ref="backgroundMusic" loop>
      <source src="../../assets/other/青花瓷.mp3" type="audio/mpeg">
    </audio>
    <!-- <audio src="../../assets/other/青花瓷.mp3" autoplay loop></audio> -->
  </template>
  <script setup>
  import Vditor from 'vditor'
  //import Vditor from 'vditor/dist/method.min'
  import 'vditor/dist/index.css'
  import axios from 'axios'
  import { progressProps } from 'element-plus';
  import { ref,computed, onMounted,reactive} from 'vue';
 

   const IPreviewOptions={
      "anchor":1,
      //theme:{current:props.active?"dark":"light"},
      mode:"dark",
      speech:{"enable":true},
      
      after () {
          if (window.innerWidth <= 768) {
            return
          }
          const outlineElement = document.getElementById('outline')
          Vditor.outlineRender(document.getElementById('preview'), outlineElement)
          //if (outlineElement.innerText.trim() !== '') {
          //  outlineElement.style.display = 'block'
          //  initOutline()
          //}
        },
   }
    
   const contenitem= ref(false);
   const backgroundMusic = ref(null);
   onMounted(() => {
    if (backgroundMusic.value) {
    const audio = backgroundMusic.value;
    audio.volume = 0.2;  // 设置音量
    audio.play().catch(error => {
      console.error("Auto-play failed:", error);
    });
  } else {
    console.error("backgroundMusic reference is null.");
  }
});

   onMounted(async()=>{
    const  result=axios({
          method: "get",
          url: "http://127.0.0.1:2023/PusthChatkimi", 
         })
         .then((res) => {
          //test=res.data.data.content 
          console.log(res.data.choices[0].message.content)
         
          return res.data.choices[0].message.content
         }) 
    let mdStrs=await result.then(res=>{return res})
    if(mdStrs!=""){ 
        contenitem.value = true; // 更新标志变量
    }
    Vditor.preview(
          document.getElementById('vditor'),
          mdStrs
      )

   
    //console.log("test"+mdStrs)
    ////addEventListener("resize",ReInitVidor)
  ////
    ////ReInitVidor(mdStrs)
     
  })
  
 



  
   
  </script>