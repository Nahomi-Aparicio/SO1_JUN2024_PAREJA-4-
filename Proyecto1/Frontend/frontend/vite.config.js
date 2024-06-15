import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react-swc'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    watch: {
      usePolling: true,
    },
    host: true,
    port: 5173,
    proxy:{
  '/insertRam':'http://backend:8000',
  '/InsertCPU':'http://backend:8000',
  '/insertProcess':'http://backend:8000',
  '/killProcess':'http://backend:8000'

    }
  }
});