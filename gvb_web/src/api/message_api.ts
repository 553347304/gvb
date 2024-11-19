import {Base} from "@/BAIYIN/HTTP";
import {Axios} from "@/api/api_index";
import {TypeApi} from "@/api/api";
import {useStore} from "@/stores";

const store = useStore()

export namespace TypeMessage {
    export interface Info {
        user_id: number
        nike_name: string
        avatar: string
        total: number
    }

    export const Info = (): Info => ({
        user_id: 0,
        nike_name: "",
        avatar: "",
        total: 0,
    })

    export interface Talk {
        id: number
        created_at: string
        send_user_id: number
        send_user_nick_name: string
        send_user_avatar: string
        rev_user_id: number
        rev_user_nick_name: string
        rev_user_avatar: string
        is_read: number
        content: string

        _is_me?: boolean
    }

    export interface ID {
        PageInfo: TypeApi.PageInfo
        user_id: number
        receive_id: number
    }

    export interface Create {
        send_user_id: number
        rev_user_id: number
        content: string
    }

    export const Create = (): Create => ({
        send_user_id: store.userInfo.user_id,
        rev_user_id: 0,
        content: "",
    })
}

class _ApiMessage {
    _url = "/api/message"

    UserList(id: number, params: TypeApi.PageInfo): Promise<Base.ResponseList<TypeMessage.Info>> {
        return Axios.get(this._url + "_user_list/" + id, {params})
    }

    List(params: TypeMessage.ID): Promise<Base.ResponseList<TypeMessage.Talk>> {
        return Axios.get(this._url, {params})
    }

    Create(data: TypeMessage.ID): Promise<Base.ResponseList<TypeMessage.Talk>> {
        return Axios.post(this._url, data)
    }

}

export const ApiMessage = new _ApiMessage()