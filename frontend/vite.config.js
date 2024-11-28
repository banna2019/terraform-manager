import { defineConfig } from 'vite';  
import vue from '@vitejs/plugin-vue';  

export default defineConfig({  
  plugins: [vue()],  
  server: {  
    proxy: {  
      '/api': {  
        target: 'http://localhost:8080', // 后端服务地址  
        changeOrigin: true, // 是否需要改变请求头中的 Origin  
        rewrite: (path) => path.replace(/^\/api/, ''), // 可选：重写路径，移除 /api 前缀  
      }, 
    },
    port: 8080   
  },  
});