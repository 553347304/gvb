import {Base} from "@/BAIYIN/HTTP";
import {Axios} from "@/api/api_index";
import {TypeApi} from "@/api/api";
import {h, ref} from "vue";
import {Tag} from "ant-design-vue";
import {DefineUtils} from "@/api/api_utils/type";
import router from "@/router";
import {Message} from "@arco-design/web-vue";
import {cacheRequest} from "@/api/api_index";

export namespace DefineBigModel {
    const modeDict: { [key: number]: string } = {
        0: "",
        1: "精准匹配",
        2: "模糊匹配",
        3: "前缀匹配",
        4: "正则匹配",
    }
    export const modeOption = [
        {label: "精准匹配", value: 1},
        {label: "模糊匹配", value: 2},
        {label: "前缀匹配", value: 3},
        {label: "正则匹配", value: 4},
    ]
    export const Columns = [
        {title: '规则名称', dataIndex: 'name',},
        {
            title: '匹配模式', dataIndex: 'mode', render: (data: any) => {
                const record = data.record as TypeBigModel.AutoReply
                return h("span", null, {default: () => modeDict[record.mode]})
            }
        },
        {title: '规则', dataIndex: 'rule',},
        {title: '回复内容', dataIndex: 'rule_content',},
        {title: '创建时间', slotName: 'created_at',},
        {title: '操作', slotName: 'action',},
    ];
    export const ColumnsTag = [
        {title: '标签名称', dataIndex: 'title',},
        {
            title: '标签颜色', dataIndex: 'color', render: (data: any) => {
                const record = data.record as TypeBigModel.TagList
                return h(Tag, {color: record.color}, {default: () => record.title})
            }
        },
        {title: '角色数量', dataIndex: 'role_total',},
        {title: '创建时间', slotName: 'created_at',},
        {title: '操作', slotName: 'action',},
    ];
    export const ColumnsRole = [
        {title: '角色名称', dataIndex: 'name',},
        {title: '是否启用', slotName: 'enable',},
        {title: '角色图标', slotName: 'icon',},
        {title: '角色简介', slotName: 'abstract',},
        {title: '角色标签', slotName: 'tags',},
        {title: '消耗积分', dataIndex: 'scope',},
        {title: '对话开场白', slotName: 'prologue',},
        {title: '角色设定', slotName: 'prompt',},
        {title: '自动回复', slotName: 'auto_reply',},
        {title: '创建时间', slotName: 'created_at',},
        {title: '操作', slotName: 'action',},
    ];
    export const ColumnsSession = [
        {title: '用户名称', dataIndex: 'nick_name',},
        {title: '会话', dataIndex: 'session_name',},
        {title: '角色', dataIndex: 'role_name',},
        {title: '对话次数', dataIndex: 'chat_total',},
        {title: '最后一次聊天内容', dataIndex: 'last_content',},
        {title: '创建时间', slotName: 'created_at',},
        {title: '操作', slotName: 'action',},
    ];
}
export namespace TypeBigModel {
    interface setting {
        name: string
        enable: boolean
        order: number
        api_key: string
        api_secret: string
        title: string
        prompt: string
        slogan: string
        help: string
    }

    interface session {
        chat_scope: number
        session_score: number
        day_score: number
    }

    interface model_list {
        label: string
        value: string
        disable: number
    }

    export interface options {
        setting: setting
        session: session
        model_list: model_list[]
    }

    export const options = (): options => ({
        setting: {
            name: "",
            enable: true,
            order: 0,
            api_key: "",
            api_secret: "",
            title: "",
            prompt: "",
            slogan: "",
            help: "",
        },
        session: {
            chat_scope: 0,
            session_score: 0,
            day_score: 0,
        },
        model_list: [],
    })

    export interface AutoReply {
        id: number
        name: string
        mode: number
        rule: string
        rule_content: string
        created_at: string
    }

    export const AutoReply = (): AutoReply => ({
        id: 0,
        name: "",
        mode: 1,
        rule: "",
        rule_content: "",
        created_at: "",
    })

    export interface TagList {
        id: number
        created_at: string
        title: string
        color: string
        role_total: number
    }

    export interface TagCreate {
        id?: number
        title: string
        color: string
    }

    export const TagCreate = (): TagCreate => ({
        id: 0,
        title: "",
        color: "",
    })

    export interface RoleList {
        id: number
        created_at: string
        name: string
        enable: number
        icon: string
        abstract: string
        tags: TagList[]
        scope: number
        prologue: string
        prompt: string
        auto_reply: number
    }

    export interface RoleCreate {
        id?: number
        name: string
        enable: boolean
        icon: string
        abstract: string
        scope: number
        prologue: string
        prompt: string
        auto_reply: boolean
        tag_list: number[]
    }

    export const RoleCreate = (): RoleCreate => ({
        id: 0,
        name: "",
        enable: true,
        icon: "",
        abstract: "",
        scope: 3,
        prologue: "",
        prompt: "",
        auto_reply: true,
        tag_list: [],
    })

    export interface SessionList {
        id: number
        created_at: string
        user_id: number
        nick_name: string
        session_name: string
        role_name: string
        chat_total: number
        last_content: string
    }

    export interface SessionUpdate {
        session_id: number
        nick_name: string
    }

    export const SessionUpdate = (): SessionUpdate => ({
        session_id: 0,
        nick_name: "",
    })

    export interface ChatList {
        id: number
        created_at: string
        user_content: string
        user_avatar: string
        bot_content: string
        bot_avatar: string
        status: boolean
    }

    export const ChatList = (): ChatList => ({
        id: 0,
        created_at: "",
        user_content: "",
        user_avatar: "",
        bot_content: "",
        bot_avatar: "",
        status: false,
    })

    export interface SessionId {
        session_id: number
    }

    export interface ChatSse extends SessionId {
        content: string
    }

    export const ChatSse = (): ChatSse => ({
        session_id: 0,
        content: "",
    })

    export interface RoleSquare {
        id: number
        name: string
        abstract: string
        icon: string
    }

    export interface Square {
        id: number
        title: string
        role_list: RoleSquare[]
    }

    export interface RoleSessionRequest extends TypeApi.PageInfo {
        role_id: number
    }

    export interface RoleSession {
        id: number
        name: string
        created_at: string
    }

    interface RoleDetailTag {
        id: number
        title: string
        color: string
    }

    export interface RoleDetail {
        created_at: string
        id: number
        name: string
        icon: string
        abstract: string
        tags: RoleDetailTag[]
        chat_total: number
    }

    export const RoleDetail = (): RoleDetail => ({
        created_at: "",
        id: 0,
        name: "",
        icon: "",
        abstract: "",
        tags: [],
        chat_total: 0,
    })
}

class _ApiBigModel {
    _url_setting = "/api/big_model/setting"
    _url_auto_reply = "/api/big_model/auto_reply"
    _url_tag = "/api/big_model/tag"
    _url_role = "/api/big_model/role"
    _url_icon = "/api/big_model/icon"
    _url_session = "/api/big_model/session"
    _url_role_session = "/api/big_model/role_sessions"
    _url_chat = "/api/big_model/chat"
    _url_square = "/api/big_model/square"
    _url_history = "/api/big_model/role_history"
    _url_scope =   "/api/big_model/scope"


    SettingList = (params?: TypeApi.PageInfo): Promise<Base.Response<any>> => {
        return Axios.get(this._url_setting, {params})
    }
    SettingUpdate = (data: TypeBigModel.options): Promise<Base.Response<any>> => {
        data.setting.help = ""
        return Axios.put(this._url_setting, data)
    }

    AutoReplyCreate = (data: TypeBigModel.AutoReply): Promise<Base.Response<any>> => {
        return Axios.post(this._url_auto_reply, data)
    }
    AutoReplyUpdate = (data: TypeBigModel.AutoReply): Promise<Base.Response<any>> => {
        return Axios.put(this._url_auto_reply, data)
    }

    TagCreate = async (data: TypeBigModel.TagCreate): Promise<Base.Response<number>> => {
        return Axios.post(this._url_tag, data)
    }
    TagUpdate = (data: TypeBigModel.TagCreate): Promise<Base.Response<any>> => {
        return Axios.put(this._url_tag, data)
    }
    TagList = (): Promise<Base.ResponseList<DefineUtils.options>> => {
        return Axios.get(this._url_tag + "/options")
    }

    RoleList = (): Promise<Base.ResponseList<TypeBigModel.RoleList>> => {
        return Axios.get(this._url_role)
    }
    RoleCreate = (data: TypeBigModel.RoleCreate): Promise<Base.Response<any>> => {
        return Axios.post(this._url_role, data)
    }
    RoleUpdate = (data: TypeBigModel.RoleCreate): Promise<Base.Response<any>> => {
        return Axios.put(this._url_role, data)
    }
    RoleDetail = (id: number): Promise<Base.Response<TypeBigModel.RoleDetail>> => {
        return Axios.get(this._url_role + "/" + id)
    }

    RoleIcon = (): Promise<Base.ResponseList<DefineUtils.options>> => {
        return Axios.get(this._url_icon)
    }


    SessionList: (params?: TypeApi.PageInfo) => Promise<Base.Response<TypeBigModel.SessionList>> = cacheRequest(
        (params?: TypeApi.PageInfo) => {
            return Axios.get(this._url_session, {params})
        }
    )

    // SessionList = (params?: TypeApi.PageInfo): Promise<Base.Response<TypeBigModel.SessionList>> => {
    //     return Axios.get(this._url_session, {params})
    // }
    SessionCreate = (role_id: number): Promise<Base.Response<number>> => {
        return Axios.post(this._url_session, {role_id})
    }
    SessionUpdate = (data: TypeBigModel.SessionUpdate): Promise<Base.Response<any>> => {
        return Axios.put(this._url_session, data)
    }
    SessionRemove = (id: number): Promise<Base.Response<any>> => {
        return Axios.delete(this._url_session + "/" + id)
    }

    ChatList = (params: TypeBigModel.SessionId): Promise<Base.ResponseList<TypeBigModel.ChatList>> => {
        return Axios.get(this._url_chat, {params})
    }
    ChatAdminRemove = (idList: number[]): Promise<Base.ResponseList<any>> => {
        return Axios.delete(this._url_chat, {
            data: {
                id_list: idList,
            }
        })
    }

    Square = (params?: TypeApi.PageInfo): Promise<Base.ResponseList<TypeBigModel.Square>> => {
        return Axios.get(this._url_square, {params})
    }

    RoleHistory = (params?: TypeApi.PageInfo): Promise<Base.ResponseList<TypeBigModel.Square>> => {
        return Axios.get(this._url_history, {params})
    }

    RoleSession = (params?: TypeBigModel.RoleSessionRequest): Promise<Base.ResponseList<TypeBigModel.RoleSession>> => {
        return Axios.get(this._url_role_session, {params})
    }

    CheckRole = async (record: TypeBigModel.RoleSquare): Promise<void> => {
        let result = await this.RoleSession({role_id: record.id,});
        if (result.code) {
            Message.error(result.message)
            return
        }
        let sessionId: number
        if (result.data.list.length > 0) {
            sessionId = result.data.list[0].id
        } else {
            let resultCreate = await this.SessionCreate(record.id)
            if (resultCreate.code) {
                Message.error(resultCreate.message)
                return
            }
            sessionId = resultCreate.data
        }

        await router.push({
            name: "role_session",
            query: {
                roleId: record.id,
                sessionId: sessionId,
            }
        })
    }


    UserIsScope = (): Promise<Base.Response<boolean>> => {
        return Axios.get(this._url_scope)
    }
    UserScope = (status: boolean): Promise<Base.Response<boolean>> => {
        return Axios.post(this._url_scope, {status})
    }
}

export const ApiBigModel = new _ApiBigModel()

