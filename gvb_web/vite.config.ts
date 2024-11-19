import {fileURLToPath, URL} from 'node:url'

import {defineConfig, loadEnv} from 'vite'
import vue from '@vitejs/plugin-vue'
import type {ImportMetaEnv} from "./env";

export default defineConfig(({mode}) => {
    let env: Record<keyof ImportMetaEnv, string> = loadEnv(mode, process.cwd()) // 加载.env环境变量
    console.log(mode, env, process.cwd())
    return {
        envPrefix: ["VITE_"],
        plugins: [vue()],
        resolve: {
            alias: {
                '@': fileURLToPath(new URL('./src', import.meta.url))
            },
        },
        // 跨域代理
        server: {
            host: "0.0.0.0",
            port: 80,
            proxy: {
                "/api": {
                    target: env.VITE_SERVER_URL,
                    changeOrigin: true,
                },
                "/file": {
                    target: env.VITE_SERVER_URL,
                    changeOrigin: true,
                },
                "/ws": {
                    target: env.VITE_SERVER_URL.replace("http", "ws"),
                    changeOrigin: true,
                    ws: true,
                    rewrite: (path)=>path.replace(/^\/ws/, "")
                },
            }
        },
    }
})

