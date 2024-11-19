import {Axios} from "@/api/index";
import {Base} from "@/BAIYIN/HTTP";

export namespace TypeApi {
    export interface PageInfo {
        search?: string
        key?: string | number
        page?: number
        limit?: number
        sort?: string
        source?: string
    }

    export function PageInfo(): PageInfo {
        return {
            page: 1,
            limit: 100,
        }
    }

}

class _Api {

    URL = {
        User: "/api/user",
        Promotion: "/api/advert",
        Menu: "/api/menu",
        Image: "/api/image",
        Article: "/api/article",
        Message: "/api/message",
        Comment: "/api/comment",
        Chat: "/api/chat",
        Date: "/api/date",
    }

    Search = (url: string, params?: TypeApi.PageInfo): Promise<Base.ResponseList<any>> => {
        return Axios.get(url + "_search", {params})
    }

    List = (url: string, params?: TypeApi.PageInfo): Promise<Base.ResponseList<any>> => {
        return Axios.get(url + "_list", {params})
    }

    Create = (url: string, data: any): Promise<Base.ResponseList<any>> => {
        return Axios.post(url + "_create", data)
    }

    Update = (url: string, data: any, id: number | string): Promise<Base.ResponseList<any>> => {
        return Axios.put(url + "_update/" + id.toString(), data)
    }

    Remove = <T>(url: string, idList: T[]): Promise<Base.ResponseList<any>> => {
        return Axios.delete(url + "_remove", {
            data: {
                id_list: idList,
            }
        })
    }
}

export const Api = new _Api()