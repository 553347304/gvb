import axios, {type InternalAxiosRequestConfig} from "axios";
import {useStore} from "@/stores";
import {Message} from "@arco-design/web-vue";


export const Axios = axios.create({
    baseURL: ""
})

Axios.interceptors.request.use((config): InternalAxiosRequestConfig<any> => {
    const store = useStore()
    config.headers["token"] = store.userInfo.token
    return config
})

Axios.interceptors.response.use((response) => {
    if (response.status !== 200) {
        Message.error(response.statusText)
        return Promise.resolve(response.statusText)
    }
    return response.data
}, (error) => {
    Message.error(error.message)
    return Promise.resolve(error.message)
})


export function cacheRequest<T>(func: (...args: any) => Promise<T>): () => Promise<T> {
    let lastRequestTime: number = 0
    let cacheData: T | null = null // 上次缓存的数据
    let currentRequest: Promise<T> | null = null

    return function () {

        const currentTime: number = Date.now()
        const timeDiff: number = currentTime - lastRequestTime

        // 1秒内重复请求   返回缓存速度
        if (timeDiff < 1000 && cacheData) {
            return Promise.resolve(cacheData)
        }

        // 发送请求
        if (!currentRequest) {
            currentRequest = func(...arguments).then((result: T) => {
                lastRequestTime = currentTime
                cacheData = result
                currentRequest = null       //重置当前请求
                return result
            })
        }
        return currentRequest
    }
}
