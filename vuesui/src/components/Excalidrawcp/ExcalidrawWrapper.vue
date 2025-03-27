<template>
    <div ref="excalidrawContainer"  style="height: 100%; width: 100%;position: fixed;top:0;left: 0;"></div>
  </template>
  
  <script>
  import { renderExcalidraw, unmountExcalidraw } from '../../utils/ExcalidrawComponent';
  
  export default {
    name: 'ExcalidrawWrapper',
    data() {
      return {
        root: null,
      };
    },
    mounted() {
      this.root = renderExcalidraw(this.$refs.excalidrawContainer);
    },
    beforeUnmount() { // Use beforeUnmount instead of beforeDestroy
    unmountExcalidraw(this.root);
    }
  };
  </script>
  
  <style scoped>
  /* 添加样式以适应 Excalidraw 容器 */
  body {
  margin: 0;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', 'Oxygen', 'Ubuntu', 'Cantarell', 'Fira Sans', 'Droid Sans', 'Helvetica Neue', sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
  </style>
  