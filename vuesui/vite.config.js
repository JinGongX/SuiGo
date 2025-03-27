import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  server:{
    open:true,
    host:'127.0.0.1',
    port:2024,
    proxy:{
      '^/api':{
        target:'http://127.0.0.1:2023/',
        ws:true,
        changeOrigin:true,
        secure:true,
        rewrite:(path)=>path.replace('/^\/api/','')
          
      }
    }

  }
})
