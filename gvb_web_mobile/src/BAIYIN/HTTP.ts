import {Axios} from "@/api";

export namespace Base {
    export interface Response<T> {
        code: number
        message: string
        data: T
    }

    export interface ResponseList<T> {
        code: number
        message: string
        data: List<T>
    }


    export interface List<T> {
        total: number
        list: T[]
    }

    // reactive(Base.List<TypeMessage.Talk>())
    export const List = <T>(): List<T> => ({total: 0, list: []})
}

class _HTTP {
    constructor() {
    }

    Get<T>(url: string, data?: any): Promise<Base.Response<T>> {
        return Axios.get(url, {data})
    }

    GetList<T>(url: string, data?: any): Promise<Base.ResponseList<T>> {
        return Axios.get(url, data) // 查询参数: {data}
    }

    Post<T>(url: string, data?: any): Promise<Base.Response<T>> {
        return Axios.post(url, data)
    }

    Delete<T>(url: string, data?: any): Promise<Base.Response<T>> {
        return Axios.delete(url, data)
    }

}

export const HTTP = new _HTTP();