<style setup>
.pdf-app #thumbnailView{
    min-width: 100%;
}
</style>
<script setup lang="ts">
// import{useRoute} from 'vue-router'
import {ref} from 'vue'
import VuePdfApp from "vue3-pdf-app"; 
import "vue3-pdf-app/dist/icons/main.css";
// const routes= useRoute()
// const ids =routes.query.namepdf
// console.log(routes.query.namepdf)
// import {ref} from 'vue'  pdf="/src/files/pdffile/{{ids}}" 
const fileName = window.location.search.split("=")[1];
// const pdfUrl = ref('publicmulu.pdf')
 const pdfUrl =ref("/src/files/pdffile/"+fileName);//ref('/src/files/pdffile/'+ids) //ref('/src/assets/pdf/srcassetspdfmulu.pdf')
 
</script>


<template>
    <div style="width: 100%; height: 100%;">
        <vue-pdf-app style="height: 100vh;" :pdf="pdfUrl" ></vue-pdf-app>
    </div>
</template>

  
