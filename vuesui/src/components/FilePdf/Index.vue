<template>
    <div style="width: 100%; height: 100%;">
        <div v-if="!$route.meta.showNavOpen">
            <div class="drag-drop-container" @dragover.prevent @drop="handleDrop" >
              <div v-if="!file">
                <p>拖拽 PDF 文件到此处上传</p>
                <input type="file"  ref="fileInput" @change="handleFileInput"  />
              </div>
              <div v-else>
                <p>已选择文件: {{ file.name }}</p>
              </div>
            </div>
            <!-- <router-link :to="{path:'/Openpdf',query:{namepdf:'ch1.pdf'}}"  target="_blank" rel="opener">ch1.pdf</router-link> -->
            <ul>
              <li v-for="pdf in pdfFiles" :key="pdf.fileName">
                <router-link :to="{path:'/Openpdf',query:{namepdf:pdf.fileName}}"  target="_blank" rel="opener">
                    <strong>{{ pdf.fileName }}</strong> 
                    <!-- - {{ pdf.filePath }} -->
                </router-link>
              </li>
            </ul> 
        </div>
       
        <router-view></router-view>
    </div>
</template>
  
<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';

const pdfFiles = ref([]);
const fileInputRef = ref(null);

onMounted(async () => { 
  try {
    // 发起 GET 请求获取 PDF 文件列表
    const response = await axios.get('http://127.0.0.1:2023/getpdfinfo');
    console.log(response.data.pdfs)
    // 将返回的数据赋值给 pdfFiles
    pdfFiles.value = response.data.pdfs;
  } catch (error) {
    console.error('Error fetching PDF files:', error);
  }

  fileInputRef.value = document.getElementById('fileInput');
   // 手动添加点击事件监听器
   if (fileInputRef.value) {
    fileInputRef.value.addEventListener('click', handleClick);
  }
});

const file = ref(null);

const handleDrop = (event) => {
  event.preventDefault();
  handleFiles(event.dataTransfer.files);
};
const handleClick = () => {
 if (fileInputRef.value) {
    fileInputRef.value.click();
  }
};
 

const handleFileInput = (event) => {
  handleFiles(event.target.files);
};

const handleFiles = async (files) => {
  if (files.length > 0) {
    file.value = files[0];
    // 在这里你可以执行上传文件的逻辑
    try {
      const formData = new FormData();
      formData.append('file', file.value);

      const response = await  axios.post('http://127.0.0.1:2023/uploadpdffile', formData, {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
      });
      refreshFileList()
      console.log('File uploaded successfully:', response.data);
       // 清空文件，准备下一次上传
       file.value = null;
    } catch (error) {
      console.error('Error uploading file:', error);
    }
  }
};

const refreshFileList = async () => {
  try {
    // 发起 GET 请求获取 PDF 文件列表
    const response = await axios.get('http://127.0.0.1:2023/getpdfinfo');
    console.log(response.data.pdfs);
    // 将返回的数据赋值给 pdfFiles
    pdfFiles.value = response.data.pdfs;
  } catch (error) {
    console.error('Error fetching PDF files:', error);
  }
};
</script>
<style scoped>
.drag-drop-container {
  border: 2px dashed #ccc;
  padding: 20px;
  text-align: center;
  cursor: pointer;
}

/* .drag-drop-container input {
  display: none;
} */
</style>