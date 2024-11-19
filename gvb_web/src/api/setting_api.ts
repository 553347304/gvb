import {Axios} from "@/api/api_index";
import {Base} from "@/BAIYIN/HTTP";

export namespace DefineSetting {

}

export namespace TypeSetting {
    export interface info {
        created_at: string
        bei_an: string
        title: string
        qq_image: string
        version: string
        email: string
        wechat_image: string
        name: string
        job: string
        addr: string
        slogan: string
        slogan_en: string
        web: string
        bilibili_url: string
        gitee_url: string
        github_url: string
    }

    export interface Email {
        host: string
        port: number
        user: string
        password: string
        default_from_email: string
        use_ssl: number
        user_tls: number
    }

    export interface QQ {
        app_id: string
        key: string
        redirect: string
    }

    export interface Send {
        receive: string
        body?: string
    }

    export const info = (): info => ({
        created_at: "",
        bei_an: "",
        title: "",
        qq_image: "",
        version: "",
        email: "",
        wechat_image: "",
        name: "",
        job: "",
        addr: "",
        slogan: "",
        slogan_en: "",
        web: "",
        bilibili_url: "",
        gitee_url: "",
        github_url: "",
    })
    export const Email = (): Email => ({
        host: "",
        port: 0,
        user: "",
        password: "",
        default_from_email: "",
        use_ssl: 0,
        user_tls: 0,
    })

}

class _ApiSetting {
    List(url: string): Promise<Base.Response<any>> {
        return Axios.get("/api/settings/" + url)
    }

    Update(url: string, data: any): Promise<Base.Response<any>> {
        return Axios.put("/api/settings/" + url, data)
    }

    Send(data: TypeSetting.Send): Promise<Base.Response<any>> {
        return Axios.post("/api/send_email", data)
    }
}

export const ApiSetting = new _ApiSetting()