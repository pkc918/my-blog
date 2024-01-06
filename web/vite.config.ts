import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import UnoCSS from 'unocss/vite'
import path from "path";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue(), UnoCSS()],
  resolve:{
    alias:{
      "@": path.resolve(__dirname, "./src")
    }
  },
  server:{
    proxy: {
      '/api': {
        target: 'http://localhost:8080/api',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, ''),
      },
    }
  }
})
