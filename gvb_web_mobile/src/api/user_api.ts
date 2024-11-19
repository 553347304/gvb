import {Axios} from "@/api/index";
import {Base} from "@/BAIYIN/HTTP";
import {DefineUtils} from "@/api/api_utils/type";

export namespace DefineUser {
    export const Columns = [
        {title: '昵称', dataIndex: 'nick_name',},
        {title: '头像', slotName: 'avatar',},
        {title: '邮箱', dataIndex: 'email',},
        {title: '角色', dataIndex: 'role',},
        {title: 'ip', dataIndex: 'ip',},
        {title: '注册时间', slotName: 'created_at',},
        {title: '操作', slotName: 'action',},
    ];
    export const Operation = [
        {label: "批量删除", value: 0},
        {label: "批量拉黑", value: 1},
    ];
    export const Filter = [
        {label: "角色过滤", value: 0, search: "role", option: DefineUtils.Role},
    ];

}
export namespace TypeUser {
    export interface Info {
        id: number
        created_at: string
        nick_name: string
        user_name: string
        avatar: string
        email: string
        tel: string
        addr: string
        token: string
        ip: string
        role: string
        sign_status: string
        integral: number
        sign: string
        link: string
    }

    export interface LoginEmail {
        user_name: string
        password: string
    }

    export const LoginEmail = (): LoginEmail => ({
        user_name: "",
        password: "",
    })

    export interface Create {
        nick_name: string
        user_name: string
        password: string
        pass_word: string
        email: string
        tel: string
        role: string
    }

    export interface Update {
        role: string
        nick_name: string
        user_id: number
    }

    export const Create = (): Create => ({
        nick_name: "",
        user_name: "",
        password: "",
        pass_word: "",
        email: "",
        tel: "",
        role: "管理员"
    })

    export const Update = (): Update => ({role: "", nick_name: "", user_id: 1,})

}

class _ApiUser {
    Create(data: TypeUser.Create): Promise<Base.Response<string>> {
        return Axios.post("/api/user_login", data)
    }

    LoginEmail(request: TypeUser.LoginEmail): Promise<Base.Response<string>> {
        return Axios.post("/api/email_login", request)
    }

    Update(data: TypeUser.Update): Promise<Base.Response<string>> {
        return Axios.put("/api/user_role", data)
    }

}

export const ApiUser = new _ApiUser()











