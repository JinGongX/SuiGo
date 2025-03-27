<template>
    <div>
      <!-- Vditor 预览模式 -->
      <div id="vditor"></div>
  
      <!-- 展示按钮 -->
      <button @click="showContent">展示</button>
    </div>
  </template>
  
  <script>
  import Vditor from "vditor";
  
  export default {
    data() {
      return {
        vditorInstance: null,
      };
    },
    methods: {
      showContent() {
        // 获取编辑器内容
        const content = this.vditorInstance.getValue();
  
        // 在控制台或者其他地方展示内容
        console.log("Vditor Content:", content);
  
        // 在实际应用中，你可以将内容渲染到页面的其他地方
        // 例如：this.showContentInOtherElement(content);
      },
    },
    mounted() {
      // 初始化 Vditor，直接设置为预览模式
      const vditor = new Vditor(document.getElementById('vditor'), {
        minHeight: 300,
        cache: {
          enable: false, // 禁用缓存以避免报错
        },
        mode: 'sv', // 'sv' 表示预览模式
      });
  
      // 将 Vditor 实例保存在组件的 data 中，以便其他方法中使用
      this.vditorInstance = vditor;
    },
  };
  </script>
  