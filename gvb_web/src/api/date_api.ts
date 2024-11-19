import {Axios} from "@/api/api_index";
import {Api} from "@/api/api";
import {Base} from "@/BAIYIN/HTTP";

export namespace TypeDate {
    export interface List {
        date_list: string[]
        login_data: number[]
        sign_data: number[]
    }

    export const List = (): List => ({
        date_list: [""],
        login_data: [0],
        sign_data: [0],
    })

    export interface Total {
        user: number
        user_same_day: number
        login_same_day: number
        message: number
        chat_message: number
        article: number
    }

    export const Total = (): Total => ({
        user: 0,
        user_same_day: 0,
        login_same_day: 0,
        message: 0,
        chat_message: 0,
        article: 0,
    })

}


class _ApiDate {
    List(): Promise<Base.Response<TypeDate.List>> {
        return Axios.get("/api/date_list")
    }

    Total(): Promise<Base.Response<TypeDate.Total>> {
        return Axios.get("/api/date_count")
    }
}

export const ApiDate = new _ApiDate()
