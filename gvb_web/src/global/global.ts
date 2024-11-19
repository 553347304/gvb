import type {Component} from "vue";
import {IconMenu, IconMessage, IconSettings, IconShareAlt, IconStorage, IconUser,} from '@arco-design/web-vue/es/icon';
import {useStore} from "@/stores";


export namespace global {
    const store = useStore()
    export const headers = {token: store.userInfo.token}

    interface MenuType {
        title: string
        icon: Component
        name?: string // 路由名
        child?: MenuType[]
    }

    export const menuList: MenuType[] = [
        {title: "首页", icon: IconMenu, name: "home", child: [],},
        {
            title: "个人中心", icon: IconUser, name: "user_center", child: [
                {title: "我的消息", icon: IconUser, name: "user_message",},
            ],
        },
        {
            title: "文章管理", icon: IconUser, name: "article", child: [
                {title: "文章列表", icon: IconUser, name: "article_list",},
                {title: "图片列表", icon: IconUser, name: "image_list",},
                {title: "评论列表", icon: IconUser, name: "comment_list",},
            ],
        },
        {
            title: "用户管理", icon: IconUser, name: "users", child: [
                {title: "用户列表", icon: IconUser, name: "user_list",},
                {title: "消息列表", icon: IconUser, name: "message_list",},
            ],
        },
        {
            title: "群聊管理", icon: IconMessage, name: "chat_group", child: [
                {title: "群聊列表", icon: IconMessage, name: "chat_list",}
            ],
        },
        {
            title: "系统管理", icon: IconSettings, name: "system", child: [
                {title: "菜单列表", icon: IconMenu, name: "system_list",},
                {title: "广告列表", icon: IconShareAlt, name: "promotion_list",},
                {title: "系统配置", icon: IconStorage, name: "config",},
            ],
        },
        {
            title: "大模型管理", icon: IconUser, name: "big_model", child: [
                {title: "参数配置", icon: IconUser, name: "options",},
                {title: "角色配置", icon: IconUser, name: "chat_role",},
                {title: "会话管理", icon: IconUser, name: "session",},
            ],
        },
    ]
}

