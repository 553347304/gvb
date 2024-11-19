export namespace TypeChat {
    // type 类型
    // outChatRoom   = -1 // 离开聊天室
    // inChatRoom    = 0  // 进入聊天室
    // textMessage   = 1  // 文本消息
    // imageMessage  = 2  // 图片消息
    // voiceMessage  = 3  // 语音消息
    // videoMessage  = 4  // 视频消息
    // systemMessage = 5  // 系统消息
    export interface List {
        id: number
        created_at: string
        nick_name: string
        avatar: string
        message: string
        addr: string
        is_group: boolean
        type: number

        _is_me: boolean
    }

    export interface Message {
        nick_name: string
        avatar: string
        message: string
        type: number
        online: number
        created_at: string
    }

    export const Message = (): Message => ({
        nick_name: "",
        avatar: "",
        message: "",
        type: 0,
        online: 0,
        created_at: "",
    })

    export interface Send {
        message: string
        type: number
    }
    export const Send = (): Send => ({
        message: "",
        type: 0,
    })
}
