import vue from '@vitejs/plugin-vue'
import { defineConfig, loadEnv } from "vite"

// https://vitejs.dev/config/
export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd(), '');
  return {
    plugins: [vue()],
    server: {
      proxy: {
        "/api": env.HOST
      }
    }
  }
})